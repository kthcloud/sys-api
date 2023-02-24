package dto

type Status struct {
	Hosts []HostStatus `json:"hosts,omitempty"`
}
