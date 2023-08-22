package v2_gpu_info

import (
	"github.com/gin-gonic/gin"
	"landing-api/pkg/sys"
	"landing-api/pkg/validator"
	v2 "landing-api/routers/api/v2"
	"landing-api/service"
)

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
