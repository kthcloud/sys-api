package body

import (
	"sys-api/pkg/subsystems/host_api"
	"time"
)

type TimestampedCapacities struct {
	Capacities Capacities `json:"capacities" bson:"capacities"`
	Timestamp  time.Time  `json:"timestamp" bson:"timestamp"`
}

type Capacities struct {
	RAM     RamCapacities     `json:"ram" bson:"ram"`
	CpuCore CpuCoreCapacities `json:"cpuCore" bson:"cpuCore"`
	GPU     GpuCapacities     `json:"gpu" bson:"gpu"`
	Hosts   []HostCapacities  `json:"hosts" bson:"hosts"`
}

type HostGpuCapacities struct {
	Count int `json:"count" bson:"count"`
}

type HostRamCapacities struct {
	Total int `json:"total" bson:"total"`
}

type HostCapacities struct {
	HostBase
	host_api.Capacities
}

type RamCapacities struct {
	Used  int `json:"used" bson:"used"`
	Total int `json:"total" bson:"total"`
}

type CpuCoreCapacities struct {
	Used  int `json:"used" bson:"used"`
	Total int `json:"total" bson:"total"`
}

type GpuCapacities struct {
	Total int `json:"total" bson:"total"`
}
