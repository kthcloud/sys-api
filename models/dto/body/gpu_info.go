package body

import "time"

type GpuInfo struct {
	HostGpuInfo []HostGpuInfo `json:"hosts" bson:"hosts"`
}

type TimestampedGpuInfo struct {
	GpuInfo   GpuInfo   `json:"gpuInfo" bson:"gpuInfo"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
