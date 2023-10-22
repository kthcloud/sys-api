package poll

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
	"sys-api/models"
	"sys-api/models/dto/body"
	"sys-api/models/stats"
	"sys-api/pkg/conf"
	"time"
)

func GetK8sStats() (*stats.K8sStats, error) {

	outputs := make(map[string]*body.K8sStats)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

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

			mu.Lock()
			outputs[name] = &body.K8sStats{
				PodCount: len(list.Items),
			}
			mu.Unlock()

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

func StatsWorker(ctx context.Context) {
	defer log.Println("stats worker stopped")

	makeError := func(err error) error {
		return fmt.Errorf("stats worker experienced an issue: %s", err)
	}

	for {
		select {
		case <-time.After(StatsSleep):
			k8sStats, err := GetK8sStats()
			if err != nil {
				log.Println(makeError(err))
				time.Sleep(StatsSleep)
				continue
			}

			if k8sStats == nil {
				log.Println(makeError(fmt.Errorf("no k8s stats were found")))
			} else {
				collected := body.Stats{
					K8sStats: body.K8sStats{
						PodCount: k8sStats.PodCount,
					},
				}

				timestamped := body.TimestampedStats{
					Stats:     collected,
					Timestamp: time.Now(),
				}

				_, err = models.StatsCollection.InsertOne(context.TODO(), timestamped)
				if err != nil {
					log.Println(makeError(err))
					log.Println("sleeping for an extra minute")
					time.Sleep(60 * time.Second)
					continue
				}

				err = DeleteUntilNItemsLeft(models.StatsCollection, 1000)
				if err != nil {
					log.Println(makeError(err))
					log.Println("sleeping for an extra minute")
					time.Sleep(60 * time.Second)
					continue
				}
			}

		case <-ctx.Done():
			return
		}

	}
}
