package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sys-api/dto/query"
	"sys-api/pkg/app"
	"sys-api/pkg/app/status_codes"
	"sys-api/service"
)

// GetGpuInfo godoc
// @Summary GetGpuInfo GPU info
// @Description GetGpuInfo GPU info
// @Tags GPU info
// @Accept  json
// @Produce  json
// @Param n query int false "n"
// @Param Authorization header string true "With the bearer started"
// @Success 200 {array} body.TimestampedGpuInfo
// @Failure 400 {object} body.BindingError
// @Failure 403 {object} sys.ErrorResponse
// @Router /internal/gpuInfo [get]
func GetGpuInfo(c *gin.Context) {
	context := app.NewContext(c)

	auth, err := WithAuth(&context)
	if err != nil {
		context.ErrorResponse(http.StatusInternalServerError, status_codes.Error, fmt.Sprintf("Failed to get auth info: %s", err.Error()))
		return
	}

	if auth.IsAdmin == false {
		context.ErrorResponse(http.StatusForbidden, status_codes.Error, "User is not allowed to access this resource")
		return
	}

	var requestQuery query.N
	if err := context.GinContext.ShouldBind(&requestQuery); err != nil {
		context.JSONResponse(http.StatusBadRequest, CreateBindingError(err))
		return
	}

	gpuInfo, err := service.GetGpuInfo(requestQuery.N)
	if err != nil {
		context.ServerError(err, fmt.Errorf("failed to get gpu info"))
		return
	}

	if gpuInfo == nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, gpuInfo)
}
