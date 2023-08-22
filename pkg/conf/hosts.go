package conf

import (
	"fmt"
	"net"
)

type Host struct {
	Name string `json:"name"`
	IP   net.IP `json:"ip"`
	Port int    `json:"port"`
}

type ZoneHost struct {
	Host
	ZoneID string `json:"zoneId"`
}

type Zone struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Hosts []Host `json:"hosts"`
}

func (host *Host) ApiURL(route string) string {
	return fmt.Sprintf("http://%s:%d%s", host.IP.String(), host.Port, route)
}

var Zones []Zone

func GetAllHosts() []ZoneHost {
	noHosts := 0
	for _, zone := range Zones {
		noHosts += len(zone.Hosts)
	}

	hosts := make([]ZoneHost, noHosts)
	hostIdx := 0
	for _, zone := range Zones {
		for _, host := range zone.Hosts {
			hosts[hostIdx] = ZoneHost{
				Host:   host,
				ZoneID: zone.ID,
			}
			hostIdx++
		}
	}

	return hosts
}
