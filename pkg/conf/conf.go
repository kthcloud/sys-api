package conf

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"strings"
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

	setupClient := func(name string, b64 string) {
		client, err := createClient(b64)
		if err != nil {
			log.Println(makeError(fmt.Errorf("failed to connect to k8s cluster %s. Details: ", err)))
		}

		// if we can fetch namespaces, we assume cluster is ok
		_, err = client.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
		if err == nil {
			Env.K8s.Clients[name] = client
		} else {
			log.Println(makeError(err))
		}
	}

	Env.K8s.Clients = make(map[string]*kubernetes.Clientset)

	setupClient("sys", Env.K8s.Configs.Sys)
	setupClient("prod", Env.K8s.Configs.Prod)
	setupClient("dev", Env.K8s.Configs.Dev)

	clusterNames := make([]string, len(Env.K8s.Clients))
	i := 0
	for name := range Env.K8s.Clients {
		clusterNames[i] = name
		i++
	}

	if len(clusterNames) > 0 {
		log.Println("successfully connected to k8s clusters:", strings.Join(clusterNames, ", "))
	} else {
		log.Println("failed to connect to any k8s clusters")
	}

}

func createConfigFromB64(b64Config string) []byte {
	configB64 := b64Config
	config, _ := base64.StdEncoding.DecodeString(configB64)
	return config
}

func createClient(b64 string) (*kubernetes.Clientset, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create k8s client. details: %s", err)
	}

	configData := createConfigFromB64(b64)
	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig(configData)
	if err != nil {
		return nil, makeError(err)
	}

	k8sClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, makeError(err)
	}

	return k8sClient, nil
}
