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
	CpuCore CpuCoreCapacities `json:"cpuCore" bson:"cpuCore"`
	RAM     RamCapacities     `json:"ram" bson:"ram"`
	GPU     GpuCapacities     `json:"gpu" bson:"gpu"`
	Hosts   []HostCapacities  `json:"hosts" bson:"hosts"`
}

type ClusterCapacities struct {
	Name    string `json:"cluster" bson:"cluster"`
	RAM     RamCapacities
	CpuCore CpuCoreCapacities
}

type HostGpuCapacities struct {
	Count int `json:"count" bson:"count"`
}

type HostRamCapacities struct {
	Total int `json:"total" bson:"total"`
}

type HostCapacities struct {
	HostBase            `json:",inline" bson:",inline"`
	host_api.Capacities `json:",inline" bson:",inline"`
}

type RamCapacities struct {
	Total int `json:"total" bson:"total"`
}

type CpuCoreCapacities struct {
	Total int `json:"total" bson:"total"`
}

type GpuCapacities struct {
	Total int `json:"total" bson:"total"`
}
