package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"strings"
	"sys-api/models"
)

func Setup() error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup environment. details: %s", err)
	}

	filepath, found := os.LookupEnv("SYS_API_CONFIG_FILE")
	if !found {
		return makeError(fmt.Errorf("config file not found. please set SYS_API_CONFIG_FILE environment variable"))
	}

	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return makeError(err)
	}

	err = yaml.Unmarshal(yamlFile, &models.Config)
	if err != nil {
		return makeError(err)
	}

	models.Config.K8s.Clients = make(map[string]kubernetes.Clientset)

	// Load config from models.Config.K8s.ConfigDir
	// Filename without an extension is used as the cluster name
	files, err := os.ReadDir(models.Config.K8s.ConfigDir)
	if err != nil {
		return makeError(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		clusterName := strings.TrimSuffix(filename, filename[strings.LastIndex(filename, "."):])
		configData, err := os.ReadFile(models.Config.K8s.ConfigDir + "/" + file.Name())
		if err != nil {
			fmt.Println(makeError(err))
			continue
		}

		client, err := createK8sClient(configData)
		if err != nil {
			fmt.Println(makeError(err))
			continue
		}

		models.Config.K8s.Clients[clusterName] = *client

		fmt.Println("Successfully connected to k8s cluster:", clusterName)
	}

	if len(models.Config.K8s.Clients) == 0 {
		fmt.Println("No k8s clusters found. Please check your config file and ensure that the k8s config directory is correct")
	}

	return nil
}

// createK8sClient creates a k8s client from the text content of a kubeconfig file
func createK8sClient(configData []byte) (*kubernetes.Clientset, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create k8s client. details: %s", err)
	}

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
