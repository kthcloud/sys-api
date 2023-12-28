package conf

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"sys-api/models/enviroment"
	"sys-api/pkg/cloudstack"
)

var Env enviroment.Environment

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

	csClient := cloudstack.NewAsyncClient(
		Env.CS.URL,
		Env.CS.ApiKey,
		Env.CS.Secret,
		true,
	)

	log.Println("fetching available hosts")

	err = ReloadHosts()
	if err != nil {
		log.Fatalln(makeError(err))
	}

	log.Println("fetching available k8s clusters")

	// load Kubernetes clusters from cloudstack
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
		client, err := createK8sClient([]byte(configData))
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

func ReloadHosts() error {
	client := withCsClient()

	// load hosts from cloudstack
	for _, zone := range Env.CS.Zones {
		// check if zone exists
		listZonesParams := client.Zone.NewListZonesParams()
		listZonesParams.SetId(zone.ID)
		zones, err := client.Zone.ListZones(listZonesParams)
		if err != nil {
			log.Fatalln(err)
		}

		if len(zones.Zones) == 0 {
			log.Fatalln(fmt.Errorf("zone %s not found", zone.ID))
		}

		listHostsParams := client.Host.NewListHostsParams()
		listHostsParams.SetZoneid(zone.ID)
		hosts, err := client.Host.ListHosts(listHostsParams)
		if err != nil {
			log.Fatalln(err)
		}

		for _, host := range hosts.Hosts {

			name := host.Name
			displayName := host.Name
			shouldSkip := false

			if host.Hosttags != "" {
				hostTags := strings.Split(host.Hosttags, ",")
				for _, keyValue := range hostTags {
					keyValueSplit := strings.Split(keyValue, "=")

					if len(keyValueSplit) == 2 {
						key := keyValueSplit[0]
						value := keyValueSplit[1]

						if key == "displayName" {
							displayName = value
						}
					} else if len(keyValueSplit) == 1 {
						key := keyValueSplit[0]

						if key == "skipMetrics" {
							shouldSkip = true
						}
					}

				}

				if shouldSkip {
					continue
				}
			}

			hostMap := Env.GetHostMap()
			if isGoodHost(host) {
				newHost := enviroment.Host{
					ID:          host.Id,
					Name:        name,
					DisplayName: displayName,
					IP:          net.ParseIP(host.Ipaddress),
					Port:        8081, // TODO: make this configurable
					Enabled:     isHostEnabled(host),
					ZoneID:      zone.ID,
					ZoneName:    zone.Name,
				}

				hostMap[newHost.Name] = newHost
			}

			Env.SetHostMap(hostMap)
		}
	}

	zoneMap := Env.GetZoneMap()
	for name, host := range Env.GetHostMap() {
		// add host to zone
		zone, ok := Env.GetZoneMap()[host.ZoneName]
		if !ok {
			zone = enviroment.Zone{
				ID:      host.ZoneID,
				Name:    host.ZoneName,
				HostMap: make(map[string]enviroment.Host),
			}
		}

		zone.HostMap[name] = host
		zoneMap[host.ZoneName] = zone
	}
	Env.SetZoneMap(zoneMap)

	return nil
}

func withCsClient() *cloudstack.CloudStackClient {
	return cloudstack.NewAsyncClient(
		Env.CS.URL,
		Env.CS.ApiKey,
		Env.CS.Secret,
		true,
	)
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

func isGoodHost(host *cloudstack.Host) bool {
	return host.Type == "Routing" && host.State == "Up"
}

func isHostEnabled(host *cloudstack.Host) bool {
	return host.Resourcestate == "Enabled"
}
