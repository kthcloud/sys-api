package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"landing-api/pkg/cloudstack"
	"log"
	"os"
	"regexp"
	"strings"
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

	log.Println("reading hosts from", Env.HostsPath)
	hostsJson, err := os.ReadFile(Env.HostsPath)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	err = json.Unmarshal(hostsJson, &Hosts)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	log.Println("successfully loaded", len(Hosts), "hosts")

	log.Println("fetching available k8s clusters")

	csClient := cloudstack.NewAsyncClient(
		Env.CS.URL,
		Env.CS.ApiKey,
		Env.CS.Secret,
		true,
	)
	listClusterParams := csClient.Kubernetes.NewListKubernetesClustersParams()
	listClusterParams.SetListall(true)
	clusters, err := csClient.Kubernetes.ListKubernetesClusters(listClusterParams)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	fetchConfig := func(name string, publicUrl string) string {

		log.Println("fetching k8s cluster config for", name)

		clusterIdx := -1
		for idx, cluster := range clusters.KubernetesClusters {
			if cluster.Name == name {
				clusterIdx = idx
				break
			}
		}

		if clusterIdx == -1 {
			log.Println("cluster", name, "not found")
			return ""
		}

		params := csClient.Kubernetes.NewGetKubernetesClusterConfigParams()
		params.SetId(clusters.KubernetesClusters[clusterIdx].Id)

		config, err := csClient.Kubernetes.GetKubernetesClusterConfig(params)
		if err != nil {
			log.Fatalln(makeError(err))
		}

		// use regex to replace the private ip in config.ConffigData 172.31.1.* with the public ip
		regex := regexp.MustCompile(`https://172.31.1.[0-9]+:6443`)

		config.ClusterConfig.Configdata = regex.ReplaceAllString(config.ClusterConfig.Configdata, publicUrl)

		return config.ClusterConfig.Configdata
	}

	Env.K8s.Clients = make(map[string]*kubernetes.Clientset)

	for _, cluster := range Env.K8s.Clusters {
		// get the public ip of the cluster
		publicURL := cluster.URL

		// get the config data from cloudstack
		configData := fetchConfig(cluster.Name, publicURL)
		if configData == "" {
			continue
		}

		// create the k8s client
		client, err := createClient([]byte(configData))
		if err != nil {
			log.Println(makeError(errors.New("failed to connect to k8s cluster " + cluster.Name + ". details: " + err.Error())))
			continue
		}

		Env.K8s.Clients[cluster.Name] = client
	}

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

func createClient(configData []byte) (*kubernetes.Clientset, error) {
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
