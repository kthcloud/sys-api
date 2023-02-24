package dto

import "time"

type CapacitiesDB struct {
	Capacities Capacities `json:"capacities"`
	Timestamp  time.Time  `json:"timestamp"`
}
