package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sys-api/models"
	"sys-api/models/dto/body"
)

func GetStats(n int) ([]body.TimestampedStats, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch stats from db. details: %s", err)
	}

	if n == 0 {
		n = 1
	}

	result, err := models.StatsCollection.Find(context.TODO(), bson.M{}, &options.FindOptions{
		Limit: intPtr(int64(n)),
		Sort:  bson.M{"timestamp": -1},
	})

	if err != nil {
		return nil, makeError(err)
	}

	var collected []body.TimestampedStats
	for result.Next(context.TODO()) {
		var stats body.TimestampedStats
		err := result.Decode(&stats)
		if err != nil {
			return nil, makeError(err)
		}
		collected = append(collected, stats)
	}

	return collected, nil
}
