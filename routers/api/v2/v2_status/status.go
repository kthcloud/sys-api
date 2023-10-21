package v2_status

import (
	"github.com/gin-gonic/gin"
	"sys-api/pkg/sys"
	"sys-api/pkg/validator"
	v2 "sys-api/routers/api/v2"
	"sys-api/service"
)

// Get godoc
// @Summary Get Status
// @Description Get Status
// @Tags Status
// @Accept  json
// @Produce  json
// @Param n query int false "n"
// @Success 200 {array} dto.StatusDB
// @Failure 400 {object} sys.ErrorResponse
// @Router /status [get]
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
