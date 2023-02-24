package service

import (
	"fmt"
	capacitiesModels "landing-api/models/capacities"
	"landing-api/models/dto"
	"landing-api/pkg/conf"
	"landing-api/utils/requestutils"
	"log"
	"math"
	"sync"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func GetCsCapacites() (*capacitiesModels.CsCapacities, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to get cs capacities. details: %s", err)
	}

	cs := conf.Env.CS

	csClient := cloudstack.NewAsyncClient(
		cs.Url,
		cs.ApiKey,
		cs.Secret,
		true,
	)

	csParams := csClient.SystemCapacity.NewListCapacityParams()
	csResponse, err := csClient.SystemCapacity.ListCapacity(csParams)

	if err != nil {
		err = makeError(err)
		log.Println(err)
		return nil, err
	}

	var cpuCore capacitiesModels.CpuCoreCapacities
	var ram capacitiesModels.RamCapacities

	for _, capacity := range csResponse.Capacity {
		if capacity.Name == "CPU_CORE" {
			cpuCore.Used = int(capacity.Capacityused)
			cpuCore.Total = int(capacity.Capacitytotal)
		} else if capacity.Name == "MEMORY" {
			ram.Used = convertToGB(capacity.Capacityused)
			ram.Total = convertToGB(capacity.Capacitytotal)
		}
	}

	parsedCapacities := &capacitiesModels.CsCapacities{
		CpuCore: cpuCore,
		RAM:     ram,
	}

	return parsedCapacities, nil
}

func GetHostCapacities() ([]dto.HostCapacities, error) {

	outputs := make([]*dto.HostCapacities, len(conf.Hosts))

	wg := sync.WaitGroup{}

	for idx, host := range conf.Hosts {
		wg.Add(1)
		go func(idx int, host conf.Host) {
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

			var hostCapacities dto.HostCapacities
			err = requestutils.ParseBody(response.Body, &hostCapacities)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			hostCapacities.Name = conf.Hosts[idx].Name

			outputs[idx] = &hostCapacities

			wg.Done()
		}(idx, host)
	}

	wg.Wait()

	var result []dto.HostCapacities

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