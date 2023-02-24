package dto

import "time"

type CapacitiesDB struct {
	Capacities Capacities `json:"capacities" bson:"capacities"`
	Timestamp  time.Time  `json:"timestamp" bson:"timestamp"`
}
