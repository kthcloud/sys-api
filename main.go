package main

import (
	"fmt"
	"landing-api/models"
	"landing-api/pkg/app"
	"landing-api/pkg/conf"
	"landing-api/routers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setup(context *app.Context) {
	conf.Setup()
	models.Setup()
}

func shutdown() {
	models.Shutdown()
}

func main() {
	context := app.Context{}

	setup(&context)
	defer shutdown()

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

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start http server. details: %s\n", err)
	}

}
