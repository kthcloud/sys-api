package dto

type GpuInfo struct {
	HostGPUInfo []HostGPUInfo `json:"hosts" bson:"hosts"`
}
