package dto

type Status struct {
	Hosts []HostStatus `json:"hosts" bson:"hosts"`
}
