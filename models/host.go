package models

import (
	"fmt"
	"sys-api/dto/body"
	"time"
)

type Host struct {
	Name        string `json:"name" bson:"name"`
	DisplayName string `json:"displayName" bson:"displayName"`
	Zone        string `json:"zone" bson:"zone"`

	IP   string `json:"ip" bson:"ip"`
	Port int    `json:"port" bson:"port"`

	Enabled bool `json:"enabled" bson:"enabled"`

	DeactivatedUntil *time.Time `json:"deactivatedUntil,omitempty" bson:"deactivatedUntil,omitempty"`
	LastSeenAt       time.Time  `json:"lastSeenAt" bson:"lastSeenAt"`
	RegisteredAt     time.Time  `json:"registeredAt" bson:"registeredAt"`

	// ID is the CloudStack host ID
	// Deprecated: Use Name instead
	ID string `json:"-" bson:"-"`
	// ZoneID is the CloudStack zone ID
	// Deprecated: Use Zone instead
	ZoneID string `json:"-" bson:"-"`
	// ZoneName is the name of the zone derived from the CloudStack zone ID
	// Deprecated: Use Zone instead
	ZoneName string `json:"-" bson:"-"`
}

func (host *Host) ApiURL() string {
	return fmt.Sprintf("http://%s:%d", host.IP, host.Port)
}

func NewHost(name, displayName, zone, ip string, port int, enabled bool) *Host {
	return &Host{
		Name:        name,
		DisplayName: displayName,
		Zone:        zone,
		IP:          ip,
		Port:        port,
		Enabled:     enabled,

		// These are set to the current time to simplify database queries
		LastSeenAt:   time.Now(),
		RegisteredAt: time.Now(),

		ID:       "",
		ZoneID:   "",
		ZoneName: "",
	}
}

func NewHostByParams(params *body.HostRegisterParams) *Host {
	return &Host{
		Name:        params.Name,
		DisplayName: params.DisplayName,
		Zone:        params.Zone,
		IP:          params.IP,
		Port:        params.Port,
		Enabled:     true,

		// These are set to the current time to simplify database queries
		LastSeenAt:   time.Now(),
		RegisteredAt: time.Now(),

		ID:       "",
		ZoneID:   "",
		ZoneName: "",
	}
}
