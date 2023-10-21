package v2_gpu_info

import (
	"github.com/gin-gonic/gin"
	"sys-api/pkg/sys"
	"sys-api/pkg/validator"
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
// @Success 200 {array} dto.GpuInfoDB
// @Failure 400 {object} sys.ErrorResponse
// @Router /internal/gpuInfo [get]
func Get(c *gin.Context) {
	context := sys.NewContext(c)

	isAdmin := v2.IsAdmin(&context)
	if !isAdmin {
		context.Unauthorized()
		return
	}

	rules := validator.MapData{
		"n": []string{
			"integer",
		},
	}

	validationErrors := context.ValidateQueryParams(&rules)
	if len(validationErrors) > 0 {
		context.ResponseValidationError(validationErrors)
		return
	}

	n, err := v2.GetN(context)

	gpuInfo, err := service.GetGpuInfo(n)
	if err != nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, gpuInfo)
}
