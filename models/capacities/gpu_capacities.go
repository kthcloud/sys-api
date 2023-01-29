package capacities

type GPU struct {
	Total int `json:"total"`
}

type GpuCapacities struct {
	GPU GPU `json:"gpu"`
}
