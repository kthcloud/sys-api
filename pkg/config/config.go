package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"regexp"
	"strings"
	"sys-api/models"
	"sys-api/pkg/imp/cloudstack"
	"sys-api/pkg/repository"
	"sys-api/pkg/subsystems/cs"
)

func Setup() error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup environment. details: %s", err)
	}

	filepath, found := os.LookupEnv("LANDING_CONFIG_FILE")
	if !found {
		log.Fatalln(makeError(fmt.Errorf("config file not found. please set LANDING_CONFIG_FILE environment variable")))
	}

	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return makeError(err)
	}

	err = yaml.Unmarshal(yamlFile, &models.Config)
	if err != nil {
		return makeError(err)
	}

	csClient := cloudstack.NewAsyncClient(
		models.Config.CS.URL,
		models.Config.CS.ApiKey,
		models.Config.CS.Secret,
		true,
	)

	// Load Kubernetes clusters from cloudstack
	listClusterParams := csClient.Kubernetes.NewListKubernetesClustersParams()
	listClusterParams.SetListall(true)
	clusters, err := csClient.Kubernetes.ListKubernetesClusters(listClusterParams)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	fetchConfig := func(name string, publicUrl string) string {
		clusterIdx := -1
		for idx, cluster := range clusters.KubernetesClusters {
			if cluster.Name == name {
				clusterIdx = idx
				break
			}
		}

		if clusterIdx == -1 {
			fmt.Println("cluster", name, "not found")
			return ""
		}

		params := csClient.Kubernetes.NewGetKubernetesClusterConfigParams()
		params.SetId(clusters.KubernetesClusters[clusterIdx].Id)

		k8sConfig, err := csClient.Kubernetes.GetKubernetesClusterConfig(params)
		if err != nil {
			log.Fatalln(makeError(err))
		}

		// use regex to replace the private ip in config.ConffigData 172.31.1.* with the public ip
		regex := regexp.MustCompile(`https://172.31.1.[0-9]+:6443`)

		k8sConfig.ClusterConfig.Configdata = regex.ReplaceAllString(k8sConfig.ClusterConfig.Configdata, publicUrl)

		return k8sConfig.ClusterConfig.Configdata
	}

	models.Config.K8s.Clients = make(map[string]kubernetes.Clientset)

	for _, cluster := range models.Config.K8s.Clusters {
		// get the public ip of the cluster
		publicURL := cluster.URL

		// get the config data from cloudstack
		configData := fetchConfig(cluster.Name, publicURL)
		if configData == "" {
			continue
		}

		// create the k8s client
		client, err := createK8sClient([]byte(configData))
		if err != nil {
			fmt.Println(makeError(errors.New("failed to connect to k8s cluster " + cluster.Name + ". details: " + err.Error())))
			continue
		}

		if client == nil {
			fmt.Println(makeError(errors.New("failed to connect to k8s cluster " + cluster.Name + ", client is nil")))
			continue
		}

		models.Config.K8s.Clients[cluster.Name] = *client
	}

	clusterNames := make([]string, len(models.Config.K8s.Clients))
	i := 0
	for name := range models.Config.K8s.Clients {
		clusterNames[i] = name
		i++
	}

	if len(clusterNames) > 0 {
		fmt.Println("successfully connected to k8s clusters:", strings.Join(clusterNames, ", "))
	} else {
		fmt.Println("failed to connect to any k8s clusters")
	}

	return nil
}

func SyncCloudStackHosts() error {
	zones := make(map[string]models.Zone)

	// Register hosts
	hosts, err := cs.NewClient(cs.ClientConfig{
		URL:    models.Config.CS.URL,
		ApiKey: models.Config.CS.ApiKey,
		Secret: models.Config.CS.Secret,
	}).ListHosts()

	for _, host := range hosts {
		newHost := models.NewHost(host.Name, host.DisplayName, host.Zone, host.IP, host.Port, host.Enabled)
		if err = repository.NewClient().RegisterHost(newHost); err != nil {
			return err
		}

		convertedName := convertCloudStackZone(host.Zone)
		if convertedName == nil {
			log.Printf("zone %s not found in cloudstack zone name conversion map. this is likely a unoffical zone\n", host.Zone)
			continue
		}

		if _, exists := zones[host.Zone]; !exists {
			zones[host.Zone] = models.Zone{Name: *convertedName}
		}
	}

	// Register zones
	for _, zone := range zones {
		if err := repository.NewClient().RegisterZone(&zone); err != nil {
			return err
		}
	}

	return nil
}

// convertCloudStackZone converts a cloudstack zone
// It is a temporary solution since CloudStack does not use the same zone names as other systems
func convertCloudStackZone(zoneName string) *string {
	switch zoneName {
	case "Flemingsberg":
		s := "se-flem"
		return &s
	case "Kista":
		s := "se-kista"
		return &s
	}
	return nil
}

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
