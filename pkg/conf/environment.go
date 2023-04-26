package conf

type Environment struct {
	Port int `yaml:"port" default:"8080"`

	Keycloak struct {
		URL        string `yaml:"url"`
		Realm      string `yaml:"realm"`
		AdminGroup string `yaml:"adminGroup"`
	} `yaml:"keycloak"`

	DB struct {
		URL  string `yaml:"url"`
		Name string `yaml:"name"`
	} `yaml:"db"`
}

var Env Environment
