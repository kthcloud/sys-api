package conf

import "k8s.io/client-go/kubernetes"

type Environment struct {
	Port int `yaml:"port" default:"8080"`

	HostsPath     string `yaml:"hostsPath"`
	SessionSecret string `yaml:"sessionSecret"`

	Keycloak struct {
		URL   string `yaml:"url"`
		Realm string `yaml:"realm"`
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
	} `yaml:"cs"`

	DB struct {
		URL  string `yaml:"url"`
		Name string `yaml:"name"`
	} `yaml:"db"`
}

var Env Environment
