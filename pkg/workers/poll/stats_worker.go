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

func GetK8sStats() ([]body.K8sStats, error) {
	clients := models.Config.K8s.Clients
	outputs := make([]*body.K8sStats, len(clients))
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
		outputs[worker] = &body.K8sStats{Name: name, PodCount: pods}
		mu.Unlock()

		return nil
	})

	return utils.WithoutNils(outputs), nil
}

func StatsWorker() error {
	k8sStats, err := GetK8sStats()
	if err != nil {
		return err
	}

	if k8sStats == nil {
		return fmt.Errorf("stats worker found no k8s stats. this is likely due to no k8s clusters being available")
	}

	collected := body.Stats{K8sStats: body.K8sStats{PodCount: 0}}
	for _, stat := range k8sStats {
		collected.K8sStats.PodCount += stat.PodCount
	}

	return timestamp_repository.NewClient().SaveStats(&body.TimestampedStats{
		Stats:     collected,
		Timestamp: time.Now(),
	})
}
