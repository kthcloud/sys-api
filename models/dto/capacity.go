package dto

type GpuCapacity struct {
	Count int `json:"count"`
}

type Capacity struct {
	GpuCapacity GpuCapacity `json:"gpu"`
}
