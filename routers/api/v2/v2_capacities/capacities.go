package v2_capacities

import (
	"github.com/gin-gonic/gin"
	"sys-api/pkg/sys"
	"sys-api/pkg/validator"
	v2 "sys-api/routers/api/v2"
	"sys-api/service"
)

// Get godoc
// @Summary Get Capacities
// @Description Get Capacities
// @Tags Capacities
// @Accept  json
// @Produce  json
// @Param n query int false "n"
// @Success 200 {array} dto.CapacitiesDB
// @Failure 400 {object} sys.ErrorResponse
// @Router /capacities [get]
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

	capacities, err := service.GetCapacities(n)
	if err != nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, capacities)
}
