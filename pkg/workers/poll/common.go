package poll

import (
	"context"
	"errors"
	"k8s.io/client-go/kubernetes"
	"log"
	"strings"
	"sync"
	"sys-api/models"
	"sys-api/pkg/repository"
	"time"
)

func Poller(ctx context.Context, name string, sleep time.Duration, job func() error) {
	failSleep := sleep

	for {
		select {
		case <-time.After(sleep):
			err := job()
			if err != nil {
				var failedTaskErr *FailedTaskErr
				if errors.As(err, &failedTaskErr) {
					log.Printf("%s failed for some hosts (%s). deactivating them temporarily\n", name, strings.Join(failedTaskErr.Hosts, ","))
					for _, host := range failedTaskErr.Hosts {
						_ = repository.NewClient().DeactiveHost(host, time.Now().Add(5*time.Minute))
					}

					failSleep = sleep
					continue
				}

				log.Printf("%s failed (sleeping for extra %s). details: %s\n", name, failSleep.String(), err)
				failSleep = failSleep * 2
				time.Sleep(failSleep)
				continue
			}

			failSleep = sleep
		case <-ctx.Done():
			log.Println(name + " stopped")
			return
		}
	}
}

func ForEachHost(taskName string, hosts []models.Host, job func(worker int, host *models.Host) error) error {
	wg := sync.WaitGroup{}

	mutex := sync.RWMutex{}
	var failedHosts []string

	for idx, host := range hosts {
		wg.Add(1)

		i := idx
		h := host

		go func(i int) {
			err := job(i, &h)
			if err != nil {
				log.Printf("failed to execute task %s for host %s. details: %s\n", taskName, h.IP, err)
				mutex.Lock()
				failedHosts = append(failedHosts, h.Name)
				mutex.Unlock()
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	if len(failedHosts) > 0 {
		return NewFailedTaskErr(failedHosts)
	}

	return nil
}

func ForEachCluster(taskName string, clusters map[string]kubernetes.Clientset, job func(worker int, name string, cluster *kubernetes.Clientset) error) {
	wg := sync.WaitGroup{}

	idx := 0
	for name, cluster := range clusters {
		wg.Add(1)

		i := idx
		n := name
		c := cluster

		go func() {
			err := job(i, n, &c)
			if err != nil {
				log.Printf("failed to execute task %s for cluster %s. details: %s\n", taskName, n, err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
