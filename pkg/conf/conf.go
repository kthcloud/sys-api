package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Setup() {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup environment. details: %s", err)
	}

	filepath, found := os.LookupEnv("LANDING_CONFIG_FILE")
	if !found {
		log.Fatalln(makeError(fmt.Errorf("config file not found. please set LANDING_CONFIG_FILE environment variable")))
	}

	log.Println("reading config from", filepath)
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf(makeError(err).Error())
	}

	err = yaml.Unmarshal(yamlFile, &Env)
	if err != nil {
		log.Fatalf(makeError(err).Error())
	}
}
