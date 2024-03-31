package poll

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
	"sys-api/dto/body"
	"sys-api/models"
	"sys-api/pkg/subsystems/k8s"
	"sys-api/pkg/timestamp_repository"
	"sys-api/utils"
	"time"
)

func GetClusterStats() ([]body.ClusterStats, error) {
	clients := models.Config.K8s.Clients
	outputs := make([]*body.ClusterStats, len(clients))
	mu := sync.Mutex{}

	ForEachCluster("fetch-k8s-stats", clients, func(worker int, name string, cluster *kubernetes.Clientset) error {
		makeError := func(err error) error {
			return fmt.Errorf("failed to list pods from cluster %s. details: %s", name, err)
		}

		pods, err := k8s.NewClient(cluster).GetTotalPods()
		if err != nil {
			log.Println(makeError(err))
			return nil
		}

		mu.Lock()
		outputs[worker] = &body.ClusterStats{Name: name, PodCount: pods}
		mu.Unlock()

		return nil
	})

	return utils.WithoutNils(outputs), nil
}

func StatsWorker() error {
	clusterStats, err := GetClusterStats()
	if err != nil {
		return err
	}

	if clusterStats == nil {
		return fmt.Errorf("stats worker found no k8s stats. this is likely due to no k8s clusters being available")
	}

	collected := body.Stats{K8sStats: body.K8sStats{PodCount: 0, Clusters: clusterStats}}
	for _, cluster := range clusterStats {
		collected.K8sStats.PodCount += cluster.PodCount
	}

	return timestamp_repository.NewClient().SaveStats(&body.TimestampedStats{
		Stats:     collected,
		Timestamp: time.Now(),
	})
}
