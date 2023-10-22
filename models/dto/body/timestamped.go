package body

import "time"

type Timestamped[T any] struct {
	Item      T         `json:"item" bson:"item"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

func CreateTimestamped[T any](item T) Timestamped[T] {
	return Timestamped[T]{
		Item:      item,
		Timestamp: time.Now(),
	}
}
