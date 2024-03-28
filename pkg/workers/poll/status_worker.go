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

func GetHostStatuses() ([]body.HostStatus, error) {
	allHosts, err := repository.NewClient().FetchHosts()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hosts. details: %s", err)
	}

	outputs := make([]*body.HostStatus, len(allHosts))
	mu := sync.Mutex{}

	ForEachHost("fetch-status", allHosts, func(idx int, host *models.Host) error {
		makeError := func(err error) error {
			return fmt.Errorf("failed to get  for host %s. details: %s", host.IP, err)
		}

		client := host_api.NewClient(host.ApiURL())

		status, err := client.GetStatus()
		if err != nil {
			return makeError(err)
		}

		hostStatus := body.HostStatus{
			Status: *status,
			HostBase: body.HostBase{
				Name:        host.Name,
				DisplayName: host.DisplayName,
				Zone:        host.Zone,
			},
		}

		mu.Lock()
		outputs[idx] = &hostStatus
		mu.Unlock()

		return nil
	})

	return utils.WithoutNils(outputs), nil
}

func StatusWorker() error {
	hostStatuses, err := GetHostStatuses()
	if err != nil {
		return err
	}

	if len(hostStatuses) == 0 {
		return fmt.Errorf("status worker found no host statuses. this is likely due to no hosts being available")
	}

	status := body.Status{
		Hosts: hostStatuses,
	}

	return timestamp_repository.NewClient().SaveStatus(&body.TimestampedStatus{
		Status:    status,
		Timestamp: time.Now(),
	})
}
