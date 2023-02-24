package dto

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

type Capacities struct {
	RAM     RamCapacities     `json:"ram" bson:"ram"`
	CpuCore CpuCoreCapacities `json:"cpuCore" bson:"cpuCore"`
	GPU     GpuCapacities     `json:"gpu" bson:"gpu"`
	Hosts   []HostCapacities  `json:"hosts" bson:"hosts"`
}
