package poll

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"sys-api/models"
	"sys-api/models/capacities"
	capacitiesModels "sys-api/models/capacities"
	"sys-api/models/dto/body"
	"sys-api/models/enviroment"
	"sys-api/pkg/cloudstack"
	"sys-api/pkg/conf"
	"sys-api/utils/requestutils"
	"time"
)

func GetCsCapacities() (*capacitiesModels.CsCapacities, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to get cs capacities. details: %s", err)
	}

	cs := conf.Env.CS

	csClient := cloudstack.NewAsyncClient(
		cs.URL,
		cs.ApiKey,
		cs.Secret,
		true,
	)

	capacityParams := csClient.SystemCapacity.NewListCapacityParams()
	capacityResponse, err := csClient.SystemCapacity.ListCapacity(capacityParams)

	if err != nil {
		err = makeError(err)
		log.Println(err)
		return nil, err
	}

	var cpuCore capacitiesModels.CpuCoreCapacities
	var ram capacitiesModels.RamCapacities

	for _, capacity := range capacityResponse.Capacity {
		if capacity.Name == "CPU_CORE" {
			cpuCore.Used += int(capacity.Capacityused)
			cpuCore.Total += int(capacity.Capacitytotal)
		} else if capacity.Name == "MEMORY" {
			ram.Used += convertToGB(capacity.Capacityused)
			ram.Total += convertToGB(capacity.Capacitytotal)
		}
	}

	parsedCapacities := &capacitiesModels.CsCapacities{
		CpuCore: cpuCore,
		RAM:     ram,
	}

	return parsedCapacities, nil
}

func GetHostCapacities() ([]body.HostCapacities, error) {

	allHosts := conf.Env.GetAvailableHosts()

	outputs := make([]*body.HostCapacities, len(allHosts))

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for idx, host := range allHosts {
		wg.Add(1)
		go func(idx int, host enviroment.Host) {
			makeError := func(err error) error {
				return fmt.Errorf("failed to get capacities for host %s. details: %s", host.IP.String(), err)
			}

			url := host.ApiURL("/capacities")
			response, err := requestutils.DoRequest("GET", url, nil, nil)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			var hostCapacities body.HostCapacities
			err = requestutils.ParseBody(response.Body, &hostCapacities)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			hostCapacities.Name = allHosts[idx].Name
			hostCapacities.ZoneID = allHosts[idx].ZoneID

			mu.Lock()
			outputs[idx] = &hostCapacities
			mu.Unlock()

			wg.Done()
		}(idx, host)
	}

	wg.Wait()

	var result []body.HostCapacities

	for _, output := range outputs {
		if output != nil {
			result = append(result, *output)
		}
	}

	return result, nil
}

func convertToGB(bytes int64) int {
	return int(math.Round(float64(bytes) / float64(1073741824)))
}

func CapacitiesWorker(ctx context.Context) {
	defer log.Println("capacities worker stopped")

	makeError := func(err error) error {
		return fmt.Errorf("capacitity worker experienced an issue: %s", err)
	}

	for {
		select {
		case <-time.After(CapacitiesSleep):
			csCapacities, err := GetCsCapacities()
			if err != nil || csCapacities == nil {
				csCapacities = &capacities.CsCapacities{
					RAM: capacities.RamCapacities{
						Used:  0,
						Total: 0,
					},
					CpuCore: capacities.CpuCoreCapacities{
						Used:  0,
						Total: 0,
					},
				}
			}

			gpuTotal := 0

			hostCapacities, err := GetHostCapacities()
			if err != nil || hostCapacities == nil {
				hostCapacities = make([]body.HostCapacities, 0)
			}

			for _, host := range hostCapacities {
				gpuTotal += host.GPU.Count
			}

			collected := body.Capacities{
				RAM: body.RamCapacities{
					Used:  csCapacities.RAM.Used,
					Total: csCapacities.RAM.Total,
				},
				CpuCore: body.CpuCoreCapacities{
					Used:  csCapacities.CpuCore.Used,
					Total: csCapacities.CpuCore.Total,
				},
				GPU: body.GpuCapacities{
					Total: gpuTotal,
				},
				Hosts: hostCapacities,
			}

			timestamped := body.TimestampedCapacities{
				Capacities: collected,
				Timestamp:  time.Now(),
			}

			_, err = models.CapacitiesCollection.InsertOne(context.TODO(), timestamped)
			if err != nil {
				log.Println(makeError(err))
				log.Println("sleeping for an extra minute")
				time.Sleep(60 * time.Second)
				continue
			}

			err = DeleteUntilNItemsLeft(models.CapacitiesCollection, 1000)
			if err != nil {
				log.Println(makeError(err))
				log.Println("sleeping for an extra minute")
				time.Sleep(60 * time.Second)
				continue
			}

		case <-ctx.Done():
			return
		}
	}
}
