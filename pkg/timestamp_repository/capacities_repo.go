package timestamp_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sys-api/dto/body"
	"sys-api/pkg/db"
)

func (c *Client) SaveCapacities(capacities *body.TimestampedCapacities) error {
	_, err := c.getCollection(db.ColCapacities).InsertOne(context.TODO(), *capacities)
	if err != nil {
		return err
	}

	err = c.deleteUntilNItemsLeft(c.getCollection(db.ColCapacities), c.MaxDocuments)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) FetchCapacities() ([]body.TimestampedCapacities, error) {
	res, err := c.getCollection(db.ColCapacities).Find(context.TODO(), bson.D{}, &options.FindOptions{
		Limit: intPtr(int64(c.MaxDocuments)),
		Sort:  bson.M{"timestamp": -1},
	})

	if err != nil {
		return nil, err
	}

	return parseList[body.TimestampedCapacities](res)
}
