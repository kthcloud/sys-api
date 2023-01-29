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

func (host *Host) ApiURL() string {
	return fmt.Sprintf("http://%s:%d", host.IP.String(), host.Port)
}

var Hosts []Host
