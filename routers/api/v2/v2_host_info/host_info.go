package v2_host_info

import (
	"github.com/gin-gonic/gin"
	"landing-api/pkg/sys"
	"landing-api/service"
)

func Get(c *gin.Context) {
	context := sys.NewContext(c)

	hostInfo := service.GetHostInfo()

	context.JSONResponse(200, hostInfo)
}
