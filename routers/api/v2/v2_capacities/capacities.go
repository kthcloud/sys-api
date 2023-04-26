package v2_capacities

import (
	"github.com/gin-gonic/gin"
	"landing-api/pkg/app"
	"landing-api/pkg/validator"
	v2 "landing-api/routers/api/v2"
	"landing-api/service"
)

func Get(c *gin.Context) {
	context := app.NewContext(c)

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

	capacities, err := service.GetCapacities(n)
	if err != nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, capacities)
}
