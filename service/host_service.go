package service

import (
	"landing-api/models/dto"
	"landing-api/pkg/conf"
)

func GetHostInfo() []dto.HostInfo {
	allHosts := conf.GetAllHosts()

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