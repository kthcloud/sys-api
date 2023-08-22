package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"landing-api/models"
	"landing-api/pkg/conf"
	"landing-api/pkg/workers/poll"
	"landing-api/routers"
	"log"
	"net/http"
	"os"
	"time"
)

type StartOptions struct {
	API    bool
	Poller bool
}

func shutdown() {
	models.Shutdown()
}

func Start(ctx context.Context, options *StartOptions) *http.Server {
	conf.Setup()
	models.Setup()

	if options == nil {
		options = &StartOptions{
			API:    true,
			Poller: true,
		}
	}

	if options.Poller {
		poll.Setup(ctx)
	}
	if options.API {
		ginMode, exists := os.LookupEnv("GIN_MODE")
		if exists {
			gin.SetMode(ginMode)
		} else {
			gin.SetMode("debug")
		}

		server := &http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%d", conf.Env.Port),
			Handler: routers.NewRouter(),
		}

		go func() {
			err := server.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Fatalf("failed to start http server. details: %s\n", err)
			}
		}()

		return server
	}

	return nil
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server. details: %s\n", err)
	}

	select {
	case <-ctx.Done():
		log.Println("waiting for server to shutdown...")
	}

	shutdown()

	log.Println("server exited successfully")
}
