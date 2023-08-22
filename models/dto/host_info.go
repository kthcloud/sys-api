package dto

type HostInfo struct {
	Name   string `bson:"name" json:"name"`
	ZoneID string `bson:"zoneId" json:"zoneId"`
}
