package dto

type HostStatus struct {
	Name string `json:"name"`
	CPU  struct {
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
