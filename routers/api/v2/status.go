package v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sys-api/dto/query"
	"sys-api/pkg/app"
	"sys-api/service"
)

// GetStatus godoc
// @Summary GetStatus Status
// @Description GetStatus Status
// @Tags Status
// @Accept  json
// @Produce  json
// @Param n query int false "n"
// @Success 200 {array} body.TimestampedStatus
// @Failure 400 {object} body.BindingError
// @Router /status [get]
func GetStatus(c *gin.Context) {
	context := app.NewContext(c)

	var requestQuery query.N
	if err := context.GinContext.ShouldBind(&requestQuery); err != nil {
		context.JSONResponse(http.StatusBadRequest, CreateBindingError(err))
		return
	}

	status, err := service.GetStatus(requestQuery.N)
	if err != nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, status)
}
