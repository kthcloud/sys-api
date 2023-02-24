package dto

import "time"

type StatusDB struct {
	Status    Status    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
