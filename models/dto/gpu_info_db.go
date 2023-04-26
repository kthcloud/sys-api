package dto

import "time"

type GpuInfoDB struct {
	GpuInfo   GpuInfo   `json:"gpuInfo" bson:"gpuInfo"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
