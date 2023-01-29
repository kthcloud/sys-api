package capacites_service

import (
	"fmt"
	capacitiesModels "landing-api/models/capacities"
	"landing-api/pkg/conf"
	"math"

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

	var cpuCore capacitiesModels.CpuCore
	var ram capacitiesModels.RAM

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
	return &capacitiesModels.GpuCapacities{}, nil
}

func convertToGB(bytes int64) int {
	return int(math.Round(float64(bytes) / float64(1073741824)))
}
