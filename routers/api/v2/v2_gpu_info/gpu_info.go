package v2_gpu_info

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sys-api/models/dto/query"
	"sys-api/pkg/status_codes"
	"sys-api/pkg/sys"
	v2 "sys-api/routers/api/v2"
	"sys-api/service"
)

// Get godoc
// @Summary Get GPU info
// @Description Get GPU info
// @Tags GPU info
// @Accept  json
// @Produce  json
// @Param n query int false "n"
// @Param Authorization header string true "With the bearer started"
// @Success 200 {array} body.TimestampedGpuInfo
// @Failure 400 {object} body.BindingError
// @Failure 403 {object} sys.ErrorResponse
// @Router /internal/gpuInfo [get]
func Get(c *gin.Context) {
	context := sys.NewContext(c)

	auth, err := v2.WithAuth(&context)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to get auth info: %s", err.Error()))
		return
	}

	if auth.IsAdmin == false {
		context.ErrorResponse(http.StatusForbidden, status_codes.Error, "User is not allowed to access this resource")
		return
	}

	var requestQuery query.N
	if err := context.GinContext.Bind(&requestQuery); err != nil {
		context.JSONResponse(http.StatusBadRequest, v2.CreateBindingError(err))
		return
	}

	gpuInfo, err := service.GetGpuInfo(requestQuery.N)
	if err != nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, gpuInfo)
}
