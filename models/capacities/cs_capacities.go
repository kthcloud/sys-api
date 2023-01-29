package capacities

type RamCapacities struct {
	Used  int `json:"used"`
	Total int `json:"total"`
}

type CpuCoreCapacities struct {
	Used  int `json:"used"`
	Total int `json:"total"`
}

type CsCapacities struct {
	RAM     RamCapacities     `json:"ram"`
	CpuCore CpuCoreCapacities `json:"cpuCore"`
}
