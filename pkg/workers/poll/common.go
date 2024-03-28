package poll

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
	"sys-api/models"
	"time"
)

func Poller(ctx context.Context, name string, sleep time.Duration, job func() error) {
	failSleep := sleep

	for {
		select {
		case <-time.After(sleep):
			err := job()
			if err != nil {
				log.Printf("%s failed (sleeping for extra %s). details: %s\n", name, failSleep.String(), err)
				failSleep = failSleep * 2
				time.Sleep(failSleep)
			} else {
				failSleep = sleep
			}
		case <-ctx.Done():
			log.Println(name + " stopped")
			return
		}
	}
}

func ForEachHost(taskName string, hosts []models.Host, job func(worker int, host *models.Host) error) {
	wg := sync.WaitGroup{}

	for idx, host := range hosts {
		wg.Add(1)

		i := idx
		h := host

		go func(i int) {
			err := job(i, &h)
			if err != nil {
				log.Printf("failed to execute task %s for host %s. details: %s\n", taskName, h.IP, err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
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
