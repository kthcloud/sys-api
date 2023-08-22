package poll

import (
	"context"
	"log"
	"time"
)

var CapacitiesSleep = 60 * time.Second
var StatusSleep = 1 * time.Second
var StatsSleep = 60 * time.Second
var GpuInfoSleep = 300 * time.Second

func Setup(ctx context.Context) {
	log.Println("starting pollers")

	go StatsWorker(ctx)
	go CapacitiesWorker(ctx)
	go StatusWorker(ctx)
	go GpuInfoWorker(ctx)
}
