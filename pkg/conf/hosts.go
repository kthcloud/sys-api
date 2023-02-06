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

func (host *Host) ApiURL(route string) string {
	return fmt.Sprintf("http://%s:%d%s", host.IP.String(), host.Port, route)
}

var Hosts []Host
