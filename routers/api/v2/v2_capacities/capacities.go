package v2_capacities

import (
	"github.com/gin-gonic/gin"
	"landing-api/models/capacities"
	"landing-api/models/dto"
	"landing-api/pkg/app"
	"landing-api/service"
)

func Get(c *gin.Context) {
	context := app.NewContext(c)

	csCapacites, err := service.GetCsCapacites()
	if err != nil {
		csCapacites = &capacities.CsCapacities{
			RAM: capacities.RamCapacities{
				Used:  0,
				Total: 0,
			},
			CpuCore: capacities.CpuCoreCapacities{
				Used:  0,
				Total: 0,
			},
		}
	}

	gpuTotal := 0

	hostCapacities, err := service.GetHostCapacities()
	if err != nil {
		hostCapacities = make([]dto.HostCapacities, 0)
	}

	for _, host := range hostCapacities {
		gpuTotal += host.GPU.Count
	}

	collected := dto.Capacities{
		RAM: dto.RamCapacities{
			Used:  csCapacites.RAM.Used,
			Total: csCapacites.RAM.Total,
		},
		CpuCore: dto.CpuCoreCapacities{
			Used:  csCapacites.CpuCore.Used,
			Total: csCapacites.CpuCore.Total,
		},
		GPU: dto.GpuCapacities{
			Total: gpuTotal,
		},
		Hosts: hostCapacities,
	}

	context.JSONResponse(200, collected)
}
