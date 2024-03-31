package poll

import (
	"context"
	"log"
	"sys-api/models"
)

func Setup(ctx context.Context) {
	log.Println("Starting pollers")

	go Poller(ctx, "hostFetchWorker", models.Config.Timer.HostFetch, HostFetchWorker)
	go Poller(ctx, "statsWorker", models.Config.Timer.Stats, StatsWorker)
	go Poller(ctx, "capacitiesWorker", models.Config.Timer.Capacities, CapacitiesWorker)
	go Poller(ctx, "statusWorker", models.Config.Timer.Status, StatusWorker)
	go Poller(ctx, "gpuInfoWorker", models.Config.Timer.GpuInfo, GpuInfoWorker)
}
