package main

import (
	"context"
	"flag"
	"log"
	"os"
	"sys-api/cmd"
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

	var options *cmd.StartOptions
	if api || poller {
		options = &cmd.StartOptions{
			API:    api,
			Poller: poller,
		}

		log.Println("API:", options.API)
		log.Println("Poller:", options.Poller)
	} else {
		log.Println("No workers specified, starting all")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	server := cmd.Start(ctx, options)
	if server != nil {
		defer func() {
			cancel()
			cmd.Stop(server)
		}()
	}

	quit := make(chan os.Signal)
	<-quit
	log.Println("received shutdown signal")

}
