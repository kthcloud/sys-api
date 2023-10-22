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
			Name:   host.Name,
			ZoneID: host.ZoneID,
		}
		result = append(result, hostInfo)
	}

	return result
}
