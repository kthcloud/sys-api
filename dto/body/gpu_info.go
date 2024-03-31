package body

import (
	"sys-api/pkg/subsystems/host_api"
	"time"
)

type GpuInfo struct {
	HostGpuInfo []HostGpuInfo `json:"hosts" bson:"hosts"`
}

type TimestampedGpuInfo struct {
	GpuInfo   GpuInfo   `json:"gpuInfo" bson:"gpuInfo"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

type HostGpuInfo struct {
	HostBase `json:",inline" bson:",inline"`
	GPUs     []host_api.GpuInfo `bson:"gpus" json:"gpus"`
}
