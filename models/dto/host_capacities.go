package dto

type HostGpuCapacities struct {
	Count int `json:"count"`
}

type HostCapacites struct {
	GPU HostGpuCapacities `json:"gpu"`
}
