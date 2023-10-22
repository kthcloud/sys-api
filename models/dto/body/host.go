package body

type HostGpuCapacities struct {
	Count int `json:"count" bson:"count"`
}

type HostRamCapacities struct {
	Total int `json:"total" bson:"total"`
}

type HostCapacities struct {
	Name   string            `json:"name" bson:"name"`
	GPU    HostGpuCapacities `json:"gpu" bson:"gpu"`
	RAM    HostRamCapacities `json:"ram" bson:"ram"`
	ZoneID string            `json:"zoneId" bson:"zoneId"`
}

type HostGPU struct {
	Name     string `bson:"name" json:"name"`
	Slot     string `bson:"slot" json:"slot"`
	Vendor   string `bson:"vendor" json:"vendor"`
	VendorID string `bson:"vendorId" json:"vendorId"`
	Bus      string `bson:"bus" json:"bus"`
	DeviceID string `bson:"deviceId" json:"deviceId"`
}

type HostGpuInfo struct {
	Name   string    `bson:"name" json:"name,omitempty"`
	ZoneID string    `bson:"zoneId" json:"zoneId,omitempty"`
	GPUs   []HostGPU `bson:"gpus" json:"gpus"`
}

type HostInfo struct {
	Name   string `bson:"name" json:"name"`
	ZoneID string `bson:"zoneId" json:"zoneId"`
}

type HostStatus struct {
	Name   string `json:"name"`
	ZoneID string `json:"zoneId"`
	CPU    struct {
		Temp struct {
			Main  int   `json:"main"`
			Cores []int `json:"cores"`
			Max   int   `json:"max"`
		} `json:"temp"`
		Load struct {
			Main  int   `json:"main"`
			Cores []int `json:"cores"`
			Max   int   `json:"max"`
		} `json:"load"`
	} `json:"cpu"`
	RAM struct {
		Load struct {
			Main int `json:"main"`
		} `json:"load"`
	} `json:"ram"`
	Network struct {
		Usage struct {
			ReceiveRate  int `json:"receiveRate"`
			TransmitRate int `json:"transmitRate"`
		} `json:"usage"`
	} `json:"network"`

	GPU *struct {
		Temp []struct {
			Main int `json:"main"`
		} `json:"temp"`
	} `json:"gpu,omitempty"`
}
