package cs

import (
	"math"
)

func (c *Client) GetCapacities() (*Capacities, error) {
	capacityParams := c.CsClient.SystemCapacity.NewListCapacityParams()
	capacityResponse, err := c.CsClient.SystemCapacity.ListCapacity(capacityParams)

	if err != nil {
		return nil, err
	}

	capacities := ZeroCapacities()

	for _, capacity := range capacityResponse.Capacity {
		if capacity.Name == "CPU_CORE" {
			capacities.CpuCore.Used += int(capacity.Capacityused)
			capacities.CpuCore.Total += int(capacity.Capacitytotal)
		} else if capacity.Name == "MEMORY" {
			capacities.RAM.Used += convertToGB(capacity.Capacityused)
			capacities.RAM.Total += convertToGB(capacity.Capacitytotal)
		}
	}

	return capacities, nil
}

func convertToGB(bytes int64) int {
	return int(math.Round(float64(bytes) / float64(1073741824)))
}
