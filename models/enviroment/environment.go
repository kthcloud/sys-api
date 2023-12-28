package enviroment

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"net"
	"sort"
	"sync"
)

type Host struct {
	ID          string
	Name        string
	DisplayName string
	IP          net.IP
	Port        int
	Enabled     bool
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

var HostMutex sync.Mutex

type Environment struct {
	Port int `yaml:"port" default:"8080"`

	hostMap map[string]Host
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

func (e *Environment) GetHostMap() map[string]Host {
	HostMutex.Lock()
	defer HostMutex.Unlock()

	if e.hostMap == nil {
		e.hostMap = make(map[string]Host)
	}

	return e.hostMap
}

func (e *Environment) GetZoneMap() map[string]Zone {
	HostMutex.Lock()
	defer HostMutex.Unlock()

	if e.ZoneMap == nil {
		e.ZoneMap = make(map[string]Zone)
	}

	return e.ZoneMap
}

func (e *Environment) SetHostMap(hostMap map[string]Host) {
	HostMutex.Lock()
	defer HostMutex.Unlock()

	e.hostMap = hostMap
}

func (e *Environment) SetZoneMap(zoneMap map[string]Zone) {
	HostMutex.Lock()
	defer HostMutex.Unlock()

	e.ZoneMap = zoneMap
}

func (e *Environment) GetEnabledHosts() []Host {
	HostMutex.Lock()
	defer HostMutex.Unlock()

	hosts := make([]Host, 0, len(e.hostMap))
	for _, host := range e.hostMap {
		if host.Enabled {
			hosts = append(hosts, host)
		}
	}

	sort.Slice(hosts, func(i, j int) bool {
		return hosts[i].Name < hosts[j].Name
	})

	return hosts
}

func (e *Environment) GetAvailableHosts() []Host {
	HostMutex.Lock()
	defer HostMutex.Unlock()

	hosts := make([]Host, 0, len(e.hostMap))
	for _, host := range e.hostMap {
		hosts = append(hosts, host)
	}

	sort.Slice(hosts, func(i, j int) bool {
		return hosts[i].Name < hosts[j].Name
	})

	return hosts
}
