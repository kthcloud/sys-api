package dto

type RamCapacities struct {
	Used  int `json:"used"`
	Total int `json:"total"`
}

type CpuCoreCapacities struct {
	Used  int `json:"used"`
	Total int `json:"total"`
}

type GpuCapacities struct {
	Total int `json:"total"`
}

type Capacities struct {
	RAM     RamCapacities     `json:"ram"`
	CpuCore CpuCoreCapacities `json:"cpuCore"`
	GPU     GpuCapacities     `json:"gpu"`
}
