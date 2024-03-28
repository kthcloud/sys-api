package timestamp_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sys-api/dto/body"
	"sys-api/pkg/db"
)

func (c *Client) SaveStats(stats *body.TimestampedStats) error {
	_, err := c.getCollection(db.ColStats).InsertOne(context.TODO(), *stats)
	if err != nil {
		return err
	}

	err = c.deleteUntilNItemsLeft(c.getCollection(db.ColStats), c.MaxDocuments)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) FetchStats() ([]body.TimestampedStats, error) {
	res, err := c.getCollection(db.ColStats).Find(context.TODO(), nil, &options.FindOptions{
		Limit: intPtr(int64(c.MaxDocuments)),
	})

	if err != nil {
		return nil, err
	}

	return parseList[body.TimestampedStats](res)
}
