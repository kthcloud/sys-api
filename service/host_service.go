package service

import (
	"sys-api/models/dto"
	"sys-api/pkg/conf"
)

func GetHostInfo() []dto.HostInfo {
	allHosts := conf.Env.GetEnabledHosts()

	var result []dto.HostInfo
	for _, host := range allHosts {
		hostInfo := dto.HostInfo{
			Name:   host.Name,
			ZoneID: host.ZoneID,
		}
		result = append(result, hostInfo)
	}

	return result
}
