package poll

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sys-api/models"
	"sys-api/models/dto"
	"sys-api/models/enviroment"
	"sys-api/pkg/conf"
	"sys-api/utils/requestutils"
	"time"
)

func GetHostGpuInfo() ([]dto.HostGPUInfo, error) {

	allHosts := conf.Env.GetEnabledHosts()

	outputs := make([]*dto.HostGPUInfo, len(allHosts))

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for idx, host := range allHosts {

		wg.Add(1)
		go func(idx int, host enviroment.Host) {
			makeError := func(err error) error {
				return fmt.Errorf("failed to get  for host %s. details: %s", host.IP.String(), err)
			}

			result, err := requestutils.DoRequest("GET", host.ApiURL("/gpuInfo"), nil, nil)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			var hostGpus []dto.HostGPU
			err = requestutils.ParseBody(result.Body, &hostGpus)
			if err != nil {
				log.Println(makeError(err))
				wg.Done()
				return
			}

			hostGpuInfo := dto.HostGPUInfo{
				Name:   allHosts[idx].Name,
				ZoneID: allHosts[idx].ZoneID,
				GPUs:   hostGpus,
			}

			mu.Lock()
			outputs[idx] = &hostGpuInfo
			mu.Unlock()

			wg.Done()
		}(idx, host)
	}

	wg.Wait()

	var result []dto.HostGPUInfo

	for _, output := range outputs {
		if output != nil {
			result = append(result, *output)
		}
	}

	return result, nil
}

func GpuInfoWorker(ctx context.Context) {
	defer log.Println("gpu info worker stopped")

	makeError := func(err error) error {
		return fmt.Errorf("gpu info worker experienced an issue: %s", err)
	}

	for {
		select {
		case <-time.After(GpuInfoSleep):

			hostGpuInfo, err := GetHostGpuInfo()
			if err != nil {
				log.Println(makeError(err))
				time.Sleep(GpuInfoSleep)
				continue
			}

			if len(hostGpuInfo) == 0 {
				log.Println(makeError(fmt.Errorf("no host gpu info was found")))
			} else {
				gpuInfoDB := dto.GpuInfoDB{
					GpuInfo: dto.GpuInfo{
						HostGPUInfo: hostGpuInfo,
					},
					Timestamp: time.Now().UTC(),
				}

				_, err = models.GpuInfoCollection.InsertOne(context.TODO(), gpuInfoDB)
				if err != nil {
					log.Println(makeError(err))
					log.Println("sleeping for an extra minute")
					time.Sleep(60 * time.Second)
					continue
				}

				err = DeleteUntilNItemsLeft(models.GpuInfoCollection, 100)
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
