package conf

import (
	"encoding/json"
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
)

func Setup() {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup environment. details: %s", err)
	}

	deployEnv, found := os.LookupEnv("LANDING_ENV_FILE")
	if found {
		log.Println("using env-file:", deployEnv)
		err := godotenv.Load(deployEnv)
		if err != nil {
			log.Fatalln(makeError(err))
		}
	}

	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	hostsJson, err := os.Open(Env.HostsPath)
	if err != nil {
		log.Fatalln(makeError(err))
	}
	defer func(hostsJson *os.File) {
		err := hostsJson.Close()
		if err != nil {
			log.Fatalln(makeError(err))
		}
	}(hostsJson)

	byteValue, _ := io.ReadAll(hostsJson)

	err = json.Unmarshal(byteValue, &Hosts)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	log.Println("successfully loaded", len(Hosts), "hosts")
}
