package dto

import "time"

type StatusDB struct {
	Status    Status    `json:"status" bson:"status"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
