package dto

type RAM struct {
	Used  int `json:"used"`
	Total int `json:"total"`
}

type CpuCore struct {
	Used  int `json:"used"`
	Total int `json:"total"`
}

type GPU struct {
	Total int `json:"total"`
}

type Capacities struct {
	RAM     RAM     `json:"ram"`
	CpuCore CpuCore `json:"cpuCore"`
	GPU     GPU     `json:"gpu"`
}
