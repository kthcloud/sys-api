package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"landing-api/pkg/app"
	"landing-api/pkg/auth"
	"landing-api/routers/api/v2/v2_capacities"
)

func NewRouter() *gin.Engine {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AddAllowHeaders("authorization")

	router := gin.New()
	router.Use(cors.New(corsConfig))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiv2 := router.Group("/v2")
	apiv2.Use(auth.New(auth.Check(), app.GetKeyCloakConfig()))

	setupCapacitiesRoutes(apiv2)
	setupNewsRoutes(apiv2)
	setupStatsRoutes(apiv2)
	setupStatusRoutes(apiv2)

	return router
}

func setupCapacitiesRoutes(base *gin.RouterGroup) {
	base.GET("/capacities", v2_capacities.Get)
}

func setupNewsRoutes(base *gin.RouterGroup) {
	//base.GET("/news", v2_news.GetMany)
	//
	//base.GET("/news/:newsId", v2_news.Get)
	//base.POST("/news", v2_news.Create)
	//base.DELETE("/news/:newsId", v2_news.Delete)
}

func setupStatsRoutes(base *gin.RouterGroup) {
	//base.GET("/stats", v2_stats.Get)
}

func setupStatusRoutes(base *gin.RouterGroup) {
	//base.GET("/stats", v2_status.Get)
}
