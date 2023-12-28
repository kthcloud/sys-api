package poll

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sys-api/models"
	"sys-api/models/dto/body"
	"sys-api/models/enviroment"
	"sys-api/pkg/conf"
	"sys-api/utils/requestutils"
	"time"
)

func GetHostStatuses() ([]body.HostStatus, error) {
	allHosts := conf.Env.GetAvailableHosts()

	outputs := make([]*body.HostStatus, len(allHosts))

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for idx, host := range allHosts {

		wg.Add(1)
		go func(idx int, host enviroment.Host) {
			makeError := func(err error) error {
				return fmt.Errorf("failed to get  for host %s. details: %s", host.IP.String(), err)
			}

			result, err := requestutils.DoRequest("GET", host.ApiURL("/status"), nil, nil)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			var hostStatus body.HostStatus
			err = requestutils.ParseBody(result.Body, &hostStatus)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			hostStatus.ID = allHosts[idx].ID
			hostStatus.Name = allHosts[idx].Name
			hostStatus.DisplayName = allHosts[idx].DisplayName
			hostStatus.ZoneID = allHosts[idx].ZoneID

			mu.Lock()
			outputs[idx] = &hostStatus
			mu.Unlock()

			wg.Done()
		}(idx, host)
	}

	wg.Wait()

	var result []body.HostStatus

	for _, output := range outputs {
		if output != nil {
			result = append(result, *output)
		}
	}

	return result, nil
}

func StatusWorker(ctx context.Context) {
	defer log.Println("status worker stopped")

	makeError := func(err error) error {
		return fmt.Errorf("status worker experienced an issue: %s", err)
	}

	for {
		select {
		case <-time.After(StatusSleep):
			hostsStatuses, err := GetHostStatuses()
			if err != nil {
				log.Println(makeError(err))
				time.Sleep(StatusSleep)
				continue
			}

			if len(hostsStatuses) == 0 {
				log.Println(makeError(fmt.Errorf("no hosts statuses received")))
			} else {
				status := body.Status{
					Hosts: hostsStatuses,
				}

				timestamped := body.TimestampedStatus{
					Status:    status,
					Timestamp: time.Now(),
				}

				_, err = models.StatusCollection.InsertOne(context.TODO(), timestamped)
				if err != nil {
					log.Println(makeError(err))
					log.Println("sleeping for an extra minute")
					time.Sleep(60 * time.Second)
					continue
				}

				err = DeleteUntilNItemsLeft(models.StatusCollection, 1000)
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
