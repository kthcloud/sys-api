package v2_status

import (
	"github.com/gin-gonic/gin"
	"landing-api/pkg/app"
	"landing-api/pkg/validator"
	"landing-api/service"
	"strconv"
)

func Get(c *gin.Context) {
	context := app.NewContext(c)

	rules := validator.MapData{
		"n": []string{
			"required",
			"integer",
		},
	}

	validationErrors := context.ValidateQueryParams(&rules)
	if len(validationErrors) > 0 {
		context.ResponseValidationError(validationErrors)
		return
	}

	n, err := strconv.Atoi(context.GinContext.Query("n"))
	if err != nil {
		context.JSONResponse(500, err)
	}

	status, err := service.GetStatus(n)
	if err != nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, status)
}
