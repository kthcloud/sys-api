package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sys-api/models"
	"sys-api/models/dto"
)

func GetCapacities(n int) ([]dto.CapacitiesDB, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch capacities from db. details: %s", err)
	}

	result, err := models.CapacitiesCollection.Find(context.TODO(), bson.M{}, &options.FindOptions{
		Limit: intPtr(int64(n)),
		Sort:  bson.M{"timestamp": -1},
	})

	if err != nil {
		return nil, makeError(err)
	}

	var capacities []dto.CapacitiesDB
	for result.Next(context.TODO()) {
		var capacity dto.CapacitiesDB
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
