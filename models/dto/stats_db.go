package dto

import "time"

type StatsDB struct {
	Stats     Stats     `json:"stats"`
	Timestamp time.Time `json:"timestamp"`
}
