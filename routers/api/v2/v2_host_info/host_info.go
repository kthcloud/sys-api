package v2_host_info

import (
	"github.com/gin-gonic/gin"
	"sys-api/pkg/sys"
	"sys-api/service"
)

// Get godoc
// @Summary Get Host Info
// @Description Get Host Info
// @Tags Host Info
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.HostInfo
// @Failure 400 {object} sys.ErrorResponse
// @Router /hostInfo [get]
func Get(c *gin.Context) {
	context := sys.NewContext(c)

	hostInfo := service.GetHostInfo()

	context.JSONResponse(200, hostInfo)
}
