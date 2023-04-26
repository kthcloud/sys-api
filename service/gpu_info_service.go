package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"landing-api/models"
	"landing-api/models/dto"
)

func GetGpuInfo(n int) ([]dto.GpuInfoDB, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch status from db. details: %s", err)
	}

	result, err := models.GpuInfoCollection.Find(context.TODO(), bson.M{}, &options.FindOptions{
		Limit: intPtr(int64(n)),
		Sort:  bson.M{"timestamp": -1},
	})

	if err != nil {
		return nil, makeError(err)
	}

	var collected []dto.GpuInfoDB
	for result.Next(context.TODO()) {
		var status dto.GpuInfoDB
		err := result.Decode(&status)
		if err != nil {
			return nil, makeError(err)
		}
		collected = append(collected, status)
	}

	return collected, nil
}
