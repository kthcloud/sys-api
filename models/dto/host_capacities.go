package dto

type HostGpuCapacities struct {
	Count int `json:"count"`
}

type HostRamCapacities struct {
	Total int `json:"total"`
}

type HostCapacities struct {
	Name string            `json:"name"`
	GPU  HostGpuCapacities `json:"gpu"`
	RAM  HostRamCapacities `json:"ram"`
}
