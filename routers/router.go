package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "sys-api/docs"
	"sys-api/pkg/auth"
	"sys-api/pkg/sys"
	"sys-api/routers/api/v2/v2_capacities"
	"sys-api/routers/api/v2/v2_gpu_info"
	"sys-api/routers/api/v2/v2_host_info"
	"sys-api/routers/api/v2/v2_stats"
	"sys-api/routers/api/v2/v2_status"
)

func NewRouter() *gin.Engine {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AddAllowHeaders("authorization")

	router := gin.New()
	router.Use(cors.New(corsConfig)).Use(gin.Logger()).Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/v2"
	apiv2 := router.Group("/v2")
	apiv2.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	setupHostInfoRoutes(apiv2)
	setupCapacitiesRoutes(apiv2)
	setupStatsRoutes(apiv2)
	setupStatusRoutes(apiv2)

	internal := apiv2.Group("/internal")
	internal.Use(auth.New(auth.Check(), sys.GetKeyCloakConfig()))

	setupGpuInfoRoutes(internal)

	return router
}

func setupHostInfoRoutes(base *gin.RouterGroup) {
	base.GET("/hostInfo", v2_host_info.Get)
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

func setupGpuInfoRoutes(base *gin.RouterGroup) {
	base.GET("/gpuInfo", v2_gpu_info.Get)
}
