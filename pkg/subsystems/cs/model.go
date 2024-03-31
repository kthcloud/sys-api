package cs

type Capacities struct {
	CpuCore struct {
		Used  int `json:"used"`
		Total int `json:"total"`
	} `json:"cpuCore"`
	RAM struct {
		Used  int `json:"used"`
		Total int `json:"total"`
	} `json:"ram"`
}

func ZeroCapacities() *Capacities {
	return &Capacities{}
}

type Host struct {
	Name        string
	DisplayName string
	IP          string
	Port        int
	Enabled     bool
	Zone        string
}
