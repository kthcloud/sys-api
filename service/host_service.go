package service

import (
	"sys-api/dto/body"
	"sys-api/pkg/repository"
)

func GetHostInfo() ([]body.HostInfo, error) {
	allHosts, err := repository.NewClient().FetchHosts()
	if err != nil {
		return nil, err
	}

	var result []body.HostInfo
	for _, host := range allHosts {
		hostInfo := body.HostInfo{
			HostBase: body.HostBase{
				Name:        host.Name,
				DisplayName: host.DisplayName,
				Zone:        host.Zone,
			},
		}
		result = append(result, hostInfo)
	}

	return result, nil
}
