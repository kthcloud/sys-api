package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sys-api/pkg/app"
	"sys-api/service"
)

// GetHostInfo godoc
// @Summary GetHostInfo Host Info
// @Description GetHostInfo Host Info
// @Tags Host Info
// @Accept  json
// @Produce  json
// @Success 200 {array} body.HostInfo
// @Router /hostInfo [get]
func GetHostInfo(c *gin.Context) {
	context := app.NewContext(c)

	hostInfo, err := service.GetHostInfo()
	if err != nil {
		context.ServerError(err, fmt.Errorf("failed to get host info"))
	}

	if hostInfo == nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, hostInfo)
}
