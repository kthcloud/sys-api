package poll

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
	"sys-api/dto/body"
	"sys-api/models"
	"sys-api/pkg/repository"
	"sys-api/pkg/subsystems/host_api"
	"sys-api/pkg/subsystems/k8s"
	"sys-api/pkg/timestamp_repository"
	"sys-api/utils"
	"time"
)

func GetClusterCapacities() ([]body.ClusterCapacities, error) {
	clients := models.Config.K8s.Clients
	outputs := make([]*body.ClusterCapacities, len(clients))
	mu := sync.Mutex{}

	ForEachCluster("fetch-k8s-stats", clients, func(worker int, name string, cluster *kubernetes.Clientset) error {
		makeError := func(err error) error {
			return fmt.Errorf("failed to list pods from cluster %s. details: %s", name, err)
		}

		nodes, err := k8s.NewClient(cluster).GetNodes()
		if err != nil {
			log.Println(makeError(err))
			return nil
		}

		clusterCapacities := body.ClusterCapacities{
			Name:    name,
			RAM:     body.RamCapacities{},
			CpuCore: body.CpuCoreCapacities{},
		}

		for _, node := range nodes {
			clusterCapacities.RAM.Total += node.RAM.Total
			clusterCapacities.CpuCore.Total += node.CPU.Total
		}

		mu.Lock()
		outputs[worker] = &clusterCapacities
		mu.Unlock()

		return nil
	})

	return utils.WithoutNils(outputs), nil
}

func GetHostCapacities() ([]body.HostCapacities, error) {
	allHosts, err := repository.NewClient().FetchHosts()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hosts. details: %s", err)
	}

	outputs := make([]*body.HostCapacities, len(allHosts))
	mu := sync.RWMutex{}

	err = ForEachHost("fetch-capacities", allHosts, func(idx int, host *models.Host) error {
		makeError := func(err error) error {
			return fmt.Errorf("failed to get capacities for host %s. details: %s", host.IP, err)
		}

		client := host_api.NewClient(host.ApiURL())

		capacities, err := client.GetCapacities()
		if err != nil {
			return makeError(err)
		}

		hostCapacities := body.HostCapacities{
			Capacities: *capacities,
			HostBase: body.HostBase{
				Name:        host.Name,
				DisplayName: host.DisplayName,
				Zone:        host.Zone,
			},
		}

		mu.Lock()
		outputs[idx] = &hostCapacities
		mu.Unlock()

		return nil
	})

	return utils.WithoutNils(outputs), err
}

func CapacitiesWorker() error {
	clusterCapacities, err := GetClusterCapacities()
	if err != nil || clusterCapacities == nil {
		clusterCapacities = make([]body.ClusterCapacities, 0)
	}

	if clusterCapacities == nil {
		log.Println("capacities worker could not get cloudstack capacities. this is likely due to cloudstack not being available")
	}

	hostCapacities, err := GetHostCapacities()
	if err != nil {
		return err
	}

	if len(hostCapacities) == 0 {
		log.Println("capacities worker found no host capacities. this is likely due to no hosts being available")
		hostCapacities = make([]body.HostCapacities, 0)
	}

	if len(hostCapacities) == 0 && clusterCapacities == nil {
		return fmt.Errorf("capacities worker found no capacities. this is likely due to no hosts or cloudstack being available")
	}

	gpuTotal := 0
	for _, host := range hostCapacities {
		gpuTotal += host.GPU.Count
	}

	collected := body.Capacities{
		RAM:     body.RamCapacities{},
		CpuCore: body.CpuCoreCapacities{},
		GPU: body.GpuCapacities{
			Total: gpuTotal,
		},
		Hosts: hostCapacities,
	}

	for _, cluster := range clusterCapacities {
		collected.RAM.Total += cluster.RAM.Total
		collected.CpuCore.Total += cluster.CpuCore.Total
	}

	return timestamp_repository.NewClient().SaveCapacities(&body.TimestampedCapacities{
		Capacities: collected,
		Timestamp:  time.Now(),
	})
}
