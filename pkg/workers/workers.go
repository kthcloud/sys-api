package workers

import (
	"landing-api/pkg/app"
	"log"
)

func Setup(ctx *app.Context) {
	log.Println("starting workers")

	//go StatWorker()
	//go CapacitiesWorker(ctx)
	go StatusWorker()
}
