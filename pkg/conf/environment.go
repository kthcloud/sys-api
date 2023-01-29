package conf

type Environment struct {
	Port int `env:"DEPLOY_PORT,default=8080"`

	HostsPath     string `env:"LANDING_HOSTS_PATH"`
	SessionSecret string `env:"LANDING_SESSION_SECRET,required=true"`

	Keycloak struct {
		Url   string `env:"LANDING_KEYCLOAK_URL,required=true"`
		Realm string `env:"LANDING_KEYCLOAK_REALM,required=true"`
	}

	K8s struct {
		Sys  string `env:"LANDING_K8S_SYS_CONFIG"`
		Prod string `env:"LANDING_K8S_PROD_CONFIG"`
		Dev  string `env:"LANDING_K8S_DEV_CONFIG"`
	}

	CS struct {
		Url    string `env:"LANDING_CS_URL,required=true"`
		ApiKey string `env:"LANDING_CS_API_KEY,required=true"`
		Secret string `env:"LANDING_CS_SECRET,required=true"`
	}

	DB struct {
		Url  string `env:"LANDING_DB_URL,required=true"`
		Name string `env:"LANDING_DB_NAME,required=true"`
	}
}

var Env Environment
