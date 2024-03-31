package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sys-api/dto/query"
	"sys-api/pkg/app"
	"sys-api/service"
)

// GetCapacities godoc
// @Summary GetCapacities Capacities
// @Description GetCapacities Capacities
// @Tags Capacities
// @Accept  json
// @Produce  json
// @Param n query int false "n"
// @Success 200 {array} body.TimestampedCapacities
// @Failure 400 {object} body.BindingError
// @Router /capacities [get]
func GetCapacities(c *gin.Context) {
	context := app.NewContext(c)

	var requestQuery query.N
	if err := context.GinContext.ShouldBind(&requestQuery); err != nil {
		context.JSONResponse(http.StatusBadRequest, CreateBindingError(err))
		return
	}

	capacities, err := service.GetCapacities(requestQuery.N)
	if err != nil {
		context.ServerError(err, fmt.Errorf("failed to fetch capacities"))
		return
	}

	if capacities == nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, capacities)
}
