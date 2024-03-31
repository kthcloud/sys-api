package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sys-api/dto/query"
	"sys-api/pkg/app"
	"sys-api/service"
)

// GetStats godoc
// @Summary GetStats Stats
// @Description GetStats Stats
// @Tags Stats
// @Accept  json
// @Produce  json
// @Param n query int false "n"
// @Success 200 {array} body.TimestampedStats
// @Failure 400 {object} body.BindingError
// @Router /stats [get]
func GetStats(c *gin.Context) {
	context := app.NewContext(c)

	var requestQuery query.N
	if err := context.GinContext.ShouldBind(&requestQuery); err != nil {
		context.JSONResponse(http.StatusBadRequest, CreateBindingError(err))
		return
	}

	stats, err := service.GetStats(requestQuery.N)
	if err != nil {
		context.ServerError(err, fmt.Errorf("failed to fetch stats"))
	}

	if len(stats) == 0 {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, stats)
}
