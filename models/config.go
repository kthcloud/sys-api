package models

import (
	"k8s.io/client-go/kubernetes"
	"time"
)

var Config ConfigType

type ConfigType struct {
	Port int `yaml:"port" default:"8080"`

	Discovery struct {
		Token string `yaml:"token"`
	}

	Keycloak struct {
		URL        string `yaml:"url"`
		Realm      string `yaml:"realm"`
		AdminGroup string `yaml:"adminGroup"`
	} `yaml:"keycloak"`

	K8s struct {
		ConfigDir string `yaml:"configDir"`
		Clients   map[string]kubernetes.Clientset
	} `yaml:"k8s"`

	CS struct {
		URL    string `yaml:"url"`
		ApiKey string `yaml:"apiKey"`
		Secret string `yaml:"secret"`
	} `yaml:"cs"`

	MongoDB struct {
		URL  string `yaml:"url"`
		Name string `yaml:"name"`
	} `yaml:"mongodb"`

	Timer struct {
		HostFetch  time.Duration `yaml:"hostFetch"`
		Capacities time.Duration `yaml:"capacities"`
		Status     time.Duration `yaml:"status"`
		Stats      time.Duration `yaml:"stats"`
		GpuInfo    time.Duration `yaml:"gpuInfo"`
	}
}
