package v2_capacities

import (
	"github.com/gin-gonic/gin"
	"landing-api/models/capacities"
	"landing-api/models/dto"
	"landing-api/pkg/app"
	"landing-api/service/capacites_service"
)

func Get(c *gin.Context) {
	context := app.NewContext(c)

	csCapacites, err := capacites_service.GetCsCapacites()
	if err != nil {
		csCapacites = &capacities.CsCapacities{
			RAM: capacities.RAM{
				Used:  0,
				Total: 0,
			},
			CpuCore: capacities.CpuCore{
				Used:  0,
				Total: 0,
			},
		}
	}

	gpuCapacites, err := capacites_service.GetGpuCapacities()
	if err != nil {
		gpuCapacites = &capacities.GpuCapacities{
			GPU: capacities.GPU{
				Total: 0,
			},
		}
	}

	collected := dto.Capacities{
		RAM: dto.RAM{
			Used:  csCapacites.RAM.Used,
			Total: csCapacites.RAM.Total,
		},
		CpuCore: dto.CpuCore{
			Used:  csCapacites.CpuCore.Used,
			Total: csCapacites.CpuCore.Total,
		},
		GPU: dto.GPU{
			Total: gpuCapacites.GPU.Total,
		},
	}

	context.JSONResponse(200, collected)
}
