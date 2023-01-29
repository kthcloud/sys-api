package capacities

type RAM struct {
	Used  int `json:"used"`
	Total int `json:"total"`
}

type CpuCore struct {
	Used  int `json:"used"`
	Total int `json:"total"`
}

type CsCapacities struct {
	RAM     RAM     `json:"ram"`
	CpuCore CpuCore `json:"cpuCore"`
}
