package dto

type HostGpuCapacities struct {
	Count int `json:"count" bson:"count"`
}

type HostRamCapacities struct {
	Total int `json:"total" bson:"total"`
}

type HostCapacities struct {
	Name string            `json:"name" bson:"name"`
	GPU  HostGpuCapacities `json:"gpu" bson:"gpu"`
	RAM  HostRamCapacities `json:"ram" bson:"ram"`
}
