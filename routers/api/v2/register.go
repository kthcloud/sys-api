package v2

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"sys-api/dto/body"
	"sys-api/pkg/app"
	"sys-api/service"
)

// Register godoc
// @Summary Register resource
// @Description Register resource
// @Tags Register
// @Accept  json
// @Produce  json
// @Success 204
// @Router /hostInfo [get]
func Register(c *gin.Context) {
	context := app.NewContext(c)

	// Try parse body as body.HostRegisterParams
	var requestQueryJoin body.HostRegisterParams
	if err := context.GinContext.ShouldBindBodyWith(&requestQueryJoin, binding.JSON); err == nil {
		err = service.RegisterNode(&requestQueryJoin)
		if err != nil {
			if errors.Is(err, service.BadDiscoveryTokenErr) {
				context.UserError("Invalid token")
				return
			}

			context.ServerError(err, fmt.Errorf("failed to register node"))
			return
		}

		context.OkNoContent()
		return
	}

	context.UserError("Invalid request body")
}
