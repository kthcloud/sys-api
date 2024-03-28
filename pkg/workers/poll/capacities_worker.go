package poll

import (
	"fmt"
	"log"
	"sync"
	"sys-api/dto/body"
	"sys-api/models"
	"sys-api/pkg/repository"
	"sys-api/pkg/subsystems/cs"
	"sys-api/pkg/subsystems/host_api"
	"sys-api/pkg/timestamp_repository"
	"sys-api/utils"
	"time"
)

func GetCsCapacities() (*cs.Capacities, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to get cs capacities. details: %s", err)
	}

	client := cs.NewClient(cs.ClientConfig{
		URL:    models.Config.CS.URL,
		ApiKey: models.Config.CS.ApiKey,
		Secret: models.Config.CS.Secret,
	})

	capacities, err := client.GetCapacities()
	if err != nil {
		return nil, makeError(err)
	}

	return capacities, nil
}

func GetHostCapacities() ([]body.HostCapacities, error) {
	allHosts, err := repository.NewClient().FetchHosts()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hosts. details: %s", err)
	}

	outputs := make([]*body.HostCapacities, len(allHosts))
	mu := sync.RWMutex{}

	ForEachHost("fetch-capacities", allHosts, func(idx int, host *models.Host) error {
		makeError := func(err error) error {
			return fmt.Errorf("failed to get capacities for host %s. details: %s", host.IP, err)
		}

		client := host_api.NewClient(host.ApiURL())

		capacities, err := client.GetCapacities()
		if err != nil {
			return makeError(err)
		}

		hostCapacities := body.HostCapacities{
			Capacities: *capacities,
			HostBase: body.HostBase{
				Name:        host.Name,
				DisplayName: host.DisplayName,
				Zone:        host.Zone,
			},
		}

		mu.Lock()
		outputs[idx] = &hostCapacities
		mu.Unlock()

		return nil
	})

	return utils.WithoutNils(outputs), nil
}

func CapacitiesWorker() error {
	csCapacities, err := GetCsCapacities()
	if err != nil || csCapacities == nil {
		csCapacities = cs.ZeroCapacities()
	}

	if csCapacities == nil {
		log.Println("capacities worker could not get cloudstack capacities. this is likely due to cloudstack not being available")
	}

	hostCapacities, err := GetHostCapacities()
	if err != nil {
		return err
	}

	if len(hostCapacities) == 0 {
		log.Println("capacities worker found no host capacities. this is likely due to no hosts being available")
	}

	if len(hostCapacities) == 0 && csCapacities == nil {
		return fmt.Errorf("capacities worker found no capacities. this is likely due to no hosts or cloudstack being available")
	}

	gpuTotal := 0
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

	return timestamp_repository.NewClient().SaveCapacities(&body.TimestampedCapacities{
		Capacities: collected,
		Timestamp:  time.Now(),
	})
}
