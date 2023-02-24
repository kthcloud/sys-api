package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"landing-api/routers/api/v2/v2_capacities"
	"landing-api/routers/api/v2/v2_stats"
	"landing-api/routers/api/v2/v2_status"
)

func NewRouter() *gin.Engine {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AddAllowHeaders("authorization")

	router := gin.New()
	router.Use(cors.New(corsConfig)).Use(gin.Logger()).Use(gin.Recovery())

	apiv2 := router.Group("/v2")

	setupCapacitiesRoutes(apiv2)
	setupStatsRoutes(apiv2)
	setupStatusRoutes(apiv2)

	return router
}

func setupCapacitiesRoutes(base *gin.RouterGroup) {
	base.GET("/capacities", v2_capacities.Get)
}

func setupStatsRoutes(base *gin.RouterGroup) {
	base.GET("/stats", v2_stats.Get)
}

func setupStatusRoutes(base *gin.RouterGroup) {
	base.GET("/status", v2_status.Get)
}
