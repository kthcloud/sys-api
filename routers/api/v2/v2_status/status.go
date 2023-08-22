package v2_status

import (
	"github.com/gin-gonic/gin"
	"landing-api/pkg/sys"
	"landing-api/pkg/validator"
	v2 "landing-api/routers/api/v2"
	"landing-api/service"
)

func Get(c *gin.Context) {
	context := sys.NewContext(c)

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

	status, err := service.GetStatus(n)
	if err != nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, status)
}
