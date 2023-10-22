package v2_capacities

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sys-api/models/dto/query"
	"sys-api/pkg/sys"
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
// @Success 200 {array} body.TimestampedCapacities
// @Failure 400 {object} body.BindingError
// @Router /capacities [get]
func Get(c *gin.Context) {
	context := sys.NewContext(c)

	var requestQuery query.N
	if err := context.GinContext.Bind(&requestQuery); err != nil {
		context.JSONResponse(http.StatusBadRequest, v2.CreateBindingError(err))
		return
	}

	capacities, err := service.GetCapacities(requestQuery.N)
	if err != nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, capacities)
}
