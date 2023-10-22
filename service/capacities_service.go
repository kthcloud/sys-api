package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sys-api/models"
	"sys-api/models/dto/body"
)

func GetCapacities(n int) ([]body.TimestampedCapacities, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch capacities from db. details: %s", err)
	}

	if n == 0 {
		n = 1
	}

	result, err := models.CapacitiesCollection.Find(context.TODO(), bson.M{}, &options.FindOptions{
		Limit: intPtr(int64(n)),
		Sort:  bson.M{"timestamp": -1},
	})

	if err != nil {
		return nil, makeError(err)
	}

	var capacities []body.TimestampedCapacities
	for result.Next(context.TODO()) {
		var capacity body.TimestampedCapacities
		err := result.Decode(&capacity)
		if err != nil {
			return nil, makeError(err)
		}
		capacities = append(capacities, capacity)
	}

	return capacities, nil
}

func intPtr(i int64) *int64 {
	return &i
}
