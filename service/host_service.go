package service

import (
	"sys-api/models/dto/body"
	"sys-api/pkg/conf"
)

func GetHostInfo() []body.HostInfo {
	allHosts := conf.Env.GetEnabledHosts()

	var result []body.HostInfo
	for _, host := range allHosts {
		hostInfo := body.HostInfo{
			HostBase: body.HostBase{
				ID:          host.ID,
				Name:        host.Name,
				DisplayName: host.DisplayName,
				ZoneID:      host.ZoneID,
			},
		}
		result = append(result, hostInfo)
	}

	return result
}
