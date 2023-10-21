package enviroment

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"net"
)

type Host struct {
	Name    string
	IP      net.IP
	Port    int
	Enabled bool
	// CloudStack zone ID
	ZoneID   string
	ZoneName string
}

type Zone struct {
	// CloudStack zone ID
	ID      string
	Name    string
	HostMap map[string]Host
}

func (host *Host) ApiURL(route string) string {
	return fmt.Sprintf("http://%s:%d%s", host.IP.String(), host.Port, route)
}

type Environment struct {
	Port int `yaml:"port" default:"8080"`

	HostMap map[string]Host
	ZoneMap map[string]Zone

	Keycloak struct {
		URL        string `yaml:"url"`
		Realm      string `yaml:"realm"`
		AdminGroup string `yaml:"adminGroup"`
	} `yaml:"keycloak"`

	K8s struct {
		Clusters []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		} `yaml:"clusters"`
		Clients map[string]*kubernetes.Clientset
	} `yaml:"k8s"`

	CS struct {
		URL    string `yaml:"url"`
		ApiKey string `yaml:"apiKey"`
		Secret string `yaml:"secret"`
		Zones  []struct {
			ID   string `yaml:"id"`
			Name string `yaml:"name"`
		}
	} `yaml:"cs"`

	DB struct {
		URL  string `yaml:"url"`
		Name string `yaml:"name"`
	} `yaml:"db"`
}

func (e *Environment) GetEnabledHosts() []Host {
	hosts := make([]Host, 0, len(e.HostMap))
	for _, host := range e.HostMap {
		if host.Enabled {
			hosts = append(hosts, host)
		}
	}
	return hosts
}
