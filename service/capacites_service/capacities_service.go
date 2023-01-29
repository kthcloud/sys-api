package capacites_service

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
		return nil, makeError(err)
	}

	var cpuCore capacitiesModels.CpuCoreCapacities
	var ram capacitiesModels.RamCapacities

	for _, capacity := range csResponse.Capacity {
		if capacity.Name == "CPU_CORE" {
			cpuCore.Used = convertToGB(capacity.Capacityused)
			cpuCore.Total = convertToGB(capacity.Capacitytotal)
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

func GetGpuCapacities() (*capacitiesModels.GpuCapacities, error) {

	outputs := make([]*dto.HostCapacites, len(conf.Hosts))

	wg := sync.WaitGroup{}

	for idx, host := range conf.Hosts {
		wg.Add(1)
		go func(idx int, host conf.Host) {
			makeError := func(err error) error {
				return fmt.Errorf("failed to get gpu capacitity for host %s. details: %s", host.IP.String(), err)
			}

			url := fmt.Sprintf("%s/capacities", host.ApiURL())
			response, err := requestutils.DoRequest("GET", url, nil, nil)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			var gpuCapacity dto.HostCapacites
			err = requestutils.ParseBody(response.Body, &gpuCapacity)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			outputs[idx] = &gpuCapacity

			wg.Done()
		}(idx, host)
	}

	wg.Wait()

	var result capacitiesModels.GpuCapacities

	for _, output := range outputs {
		if output != nil {
			result.Total += output.GPU.Count
		}
	}

	return &result, nil
}

func convertToGB(bytes int64) int {
	return int(math.Round(float64(bytes) / float64(1073741824)))
}
