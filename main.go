package main

import (
	"context"
	"flag"
	"log"
	"os"
	"sys-api/pkg/app"
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	_ = flag.Bool("api", false, "start api")
	_ = flag.Bool("poller", false, "start poller")
	flag.Parse()

	api := isFlagPassed("api")
	poller := isFlagPassed("poller")

	var options *app.StartOptions
	if api || poller {
		options = &app.StartOptions{
			API:    api,
			Poller: poller,
		}

		log.Println("api: ", options.API)
		log.Println("poller: ", options.Poller)
	} else {
		log.Println("no workers specified, starting all")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	server := app.Start(ctx, options)
	if server != nil {
		defer func() {
			cancel()
			app.Stop(server)
		}()
	}

	quit := make(chan os.Signal)
	<-quit
	log.Println("received shutdown signal")

}
