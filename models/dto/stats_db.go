package dto

import "time"

type StatsDB struct {
	Stats     Stats     `json:"stats" bson:"stats"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
