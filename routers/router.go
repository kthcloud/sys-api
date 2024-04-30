package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "sys-api/docs/api"
	"sys-api/pkg/app"
	"sys-api/pkg/auth"
	v2 "sys-api/routers/api/v2"
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

	apiv2.GET("/hostInfo", v2.GetHostInfo)
	apiv2.GET("/capacities", v2.GetCapacities)
	apiv2.GET("/stats", v2.GetStats)
	apiv2.GET("/status", v2.GetStatus)
	apiv2.POST("/register", v2.Register)

	internal := apiv2.Group("/internal")
	internal.Use(auth.New(auth.Check(), app.GetKeycloakConfig()))

	internal.GET("/gpuInfo", v2.GetGpuInfo)

	return router
}
