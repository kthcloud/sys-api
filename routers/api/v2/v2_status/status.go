package v2_status

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sys-api/models/dto/query"
	"sys-api/pkg/sys"
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
// @Success 200 {array} body.TimestampedStatus
// @Failure 400 {object} body.BindingError
// @Router /status [get]
func Get(c *gin.Context) {
	context := sys.NewContext(c)

	var requestQuery query.N
	if err := context.GinContext.ShouldBind(&requestQuery); err != nil {
		context.JSONResponse(http.StatusBadRequest, v2.CreateBindingError(err))
		return
	}

	status, err := service.GetStatus(requestQuery.N)
	if err != nil {
		context.JSONResponse(200, make([]interface{}, 0))
		return
	}

	context.JSONResponse(200, status)
}
