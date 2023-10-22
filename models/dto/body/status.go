package body

type Status struct {
	Hosts []HostStatus `json:"hosts" bson:"hosts"`
}

type TimestampedStatus Timestamped[Status]
