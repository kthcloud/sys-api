package v2_stats

import (
	"github.com/gin-gonic/gin"
	"landing-api/models/dto"
	"landing-api/models/stats"
	"landing-api/pkg/app"
	"landing-api/service"
)

func Get(c *gin.Context) {
	context := app.NewContext(c)

	k8sStats, err := service.GetK8sStats()
	if err != nil {
		k8sStats = &stats.K8sStats{
			PodCount: 0,
		}
	}

	collected := dto.Stats{
		K8sStats: dto.K8sStats{
			PodCount: k8sStats.PodCount,
		},
	}

	context.JSONResponse(200, collected)
}
