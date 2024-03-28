package poll

import (
	"fmt"
	"sync"
	"sys-api/dto/body"
	"sys-api/models"
	"sys-api/pkg/repository"
	"sys-api/pkg/subsystems/host_api"
	"sys-api/pkg/timestamp_repository"
	"sys-api/utils"
	"time"
)

func GetHostGpuInfo() ([]body.HostGpuInfo, error) {
	allHosts, err := repository.NewClient().FetchHosts()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hosts. details: %s", err)
	}

	outputs := make([]*body.HostGpuInfo, len(allHosts))
	mu := sync.RWMutex{}

	ForEachHost("fetch-capacities", allHosts, func(idx int, host *models.Host) error {
		makeError := func(err error) error {
			return fmt.Errorf("failed to get capacities for host %s. details: %s", host.IP, err)
		}

		client := host_api.NewClient(host.ApiURL())

		gpuInfo, err := client.GetGpuInfo()
		if err != nil {
			return makeError(err)
		}

		hostCapacities := body.HostGpuInfo{
			GPUs: gpuInfo,
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

func GpuInfoWorker() error {
	hostGpuInfo, err := GetHostGpuInfo()
	if err != nil {
		return err
	}

	if len(hostGpuInfo) == 0 {
		return fmt.Errorf("gpu info worker found no gpu info. this is likealy due to no available hosts")
	}

	return timestamp_repository.NewClient().SaveGpuInfo(&body.TimestampedGpuInfo{
		GpuInfo: body.GpuInfo{
			HostGpuInfo: hostGpuInfo,
		},
		Timestamp: time.Now(),
	})
}
