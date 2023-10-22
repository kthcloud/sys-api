package body

type GpuInfo struct {
	HostGpuInfo []HostGpuInfo `json:"hosts" bson:"hosts"`
}

type TimestampedGpuInfo Timestamped[GpuInfo]
