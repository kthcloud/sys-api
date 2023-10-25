package poll

import (
	"context"
	"fmt"
	"log"
	"sys-api/pkg/conf"
	"time"
)

func HostFetchWorker(ctx context.Context) {
	defer log.Println("host fetch worker stopped")

	makeError := func(err error) error {
		return fmt.Errorf("host fetch worker experienced an issue: %s", err)
	}

	for {
		select {
		case <-time.After(HostFetchSleep):
			err := conf.ReloadHosts()
			if err != nil {
				log.Println(makeError(err))
				continue
			}
		case <-ctx.Done():
			log.Println("shutting down host fetch worker")
			return
		}
	}
}
