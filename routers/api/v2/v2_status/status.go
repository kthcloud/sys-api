package v2_status

import (
	"github.com/gin-gonic/gin"
	"landing-api/models/dto"
	"landing-api/pkg/app"
	"landing-api/service"
)

func Get(c *gin.Context) {
	context := app.NewContext(c)

	hostsStatuses, err := service.GetHostStatuses()
	if err != nil {
		hostsStatuses = []dto.HostStatus{}
	}

	status := dto.Status{
		Hosts: hostsStatuses,
	}

	if hostsStatuses == nil {
		context.JSONResponse(200, status)
		return
	}

	context.JSONResponse(200, status)
}
