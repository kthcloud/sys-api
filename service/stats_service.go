package service

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"landing-api/models/dto"
	"landing-api/models/stats"
	"landing-api/pkg/conf"
	"log"
	"sync"
)

func GetK8sStats() (*stats.K8sStats, error) {

	outputs := make(map[string]*dto.K8sStats)

	wg := sync.WaitGroup{}

	for name, cluster := range conf.Env.K8s.Clients {
		wg.Add(1)
		go func(name string, cluster *kubernetes.Clientset) {
			makeError := func(err error) error {
				return fmt.Errorf("failed to list pods from cluster %s. details: %s", name, err)
			}

			list, err := cluster.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			outputs[name] = &dto.K8sStats{
				PodCount: len(list.Items),
			}

			wg.Done()

		}(name, cluster)
	}

	wg.Wait()

	var result stats.K8sStats

	for _, output := range outputs {
		if output != nil {
			result.PodCount += output.PodCount
		}
	}

	return &result, nil
}
