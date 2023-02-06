package dto

type HostStatus struct {
	Name string `json:"name,omitempty"`
	CPU  struct {
		Temp struct {
			Main  int   `json:"main,omitempty"`
			Cores []int `json:"cores,omitempty"`
			Max   int   `json:"max,omitempty"`
		} `json:"temp"`
		Load struct {
			Main  int   `json:"main,omitempty"`
			Cores []int `json:"cores,omitempty"`
			Max   int   `json:"max,omitempty"`
		} `json:"load"`
	} `json:"cpu"`
	RAM struct {
		Load struct {
			Main int `json:"main,omitempty"`
		} `json:"load"`
	} `json:"ram"`
	Network struct {
		Usage struct {
			ReceiveRate  int `json:"receiveRate,omitempty"`
			TransmitRate int `json:"transmitRate,omitempty"`
		} `json:"usage"`
	} `json:"network"`
}
