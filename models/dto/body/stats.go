package body

import "time"

type K8sStats struct {
	PodCount int `json:"podCount" bson:"podCount"`
}

type Stats struct {
	K8sStats K8sStats `json:"k8s" bson:"k8s"`
}

type TimestampedStats struct {
	Stats     Stats     `json:"stats" bson:"stats"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
