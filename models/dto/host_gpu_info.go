package dto

type HostGPU struct {
	Name     string `bson:"name" json:"name"`
	Slot     string `bson:"slot" json:"slot"`
	Vendor   string `bson:"vendor" json:"vendor"`
	VendorID string `bson:"vendorId" json:"vendorId"`
	Bus      string `bson:"bus" json:"bus"`
	DeviceID string `bson:"deviceId" json:"deviceId"`
}

type HostGPUInfo struct {
	Name string    `bson:"name" json:"name,omitempty"`
	GPUs []HostGPU `bson:"gpus" json:"gpus"`
}
