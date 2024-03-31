package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sys-api/models"
	"sys-api/pkg/config"
	"sys-api/pkg/db"
	"sys-api/pkg/workers/poll"
	"sys-api/routers"
	"time"
)

type StartOptions struct {
	API    bool
	Poller bool
}

type InitTask struct {
	Name      string
	Task      func() error
	Composite bool
}

func (it *InitTask) LogBegin(prefix string) {
	orange := "\033[38;5;208m"
	grey := "\033[90m"
	reset := "\033[0m"
	now := time.Now().Format("2006/01/02 15:04:05")
	taskName := it.Name
	fmt.Printf("[%s] %s %s%s%s %s...%s ", now, prefix, orange, taskName, reset, grey, reset)
	if it.Composite {
		lightBlue := "\033[38;5;39m"
		fmt.Println(lightBlue)
	}
}

func (it *InitTask) LogCompleted() {
	green := "\033[32m"
	grey := "\033[90m"
	reset := "\033[0m"
	if it.Composite {
		fmt.Printf("%s... done %s✓%s\n", grey, green, reset)
	} else {
		fmt.Printf("%s✓%s\n", green, reset)
	}
}

func (it *InitTask) LogFailed() {
	red := "\033[31m"
	grey := "\033[90m"
	reset := "\033[0m"
	if it.Composite {
		fmt.Printf("%s... failed %s✗%s\n", grey, red, reset)
	} else {
		fmt.Printf("%s✗%s\n", red, reset)
	}
}

func Start(ctx context.Context, options *StartOptions) *http.Server {
	initTasks := []InitTask{
		{Name: "Setup config", Task: config.Setup, Composite: true},
		{Name: "Setup repositories", Task: db.Setup},
	}

	for idx, task := range initTasks {
		task.LogBegin(fmt.Sprintf("(%d/%d)", idx+1, len(initTasks)))
		err := task.Task()
		if err != nil {
			task.LogFailed()
			log.Fatalf("Task %s failed. See error below:\n%s", task.Name, err.Error())
		}
		task.LogCompleted()
	}

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
			Addr:    fmt.Sprintf("0.0.0.0:%d", models.Config.Port),
			Handler: routers.NewRouter(),
		}

		go func() {
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
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

	log.Println("server exited successfully")
}
