package body

import (
	"sys-api/pkg/subsystems/host_api"
	"time"
)

type Status struct {
	Hosts []HostStatus `json:"hosts" bson:"hosts"`
}

type TimestampedStatus struct {
	Status    Status    `json:"status" bson:"status"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

type HostStatus struct {
	HostBase        `json:",inline" bson:",inline"`
	host_api.Status `json:",inline" bson:",inline"`
}
