package timestamp_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sys-api/dto/body"
	"sys-api/pkg/db"
)

func (c *Client) SaveStatus(status *body.TimestampedStatus) error {
	_, err := c.getCollection(db.ColStatus).InsertOne(context.TODO(), *status)
	if err != nil {
		return err
	}

	err = c.deleteUntilNItemsLeft(c.getCollection(db.ColStatus), c.MaxDocuments)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) FetchStatus() ([]body.TimestampedStatus, error) {
	res, err := c.getCollection(db.ColStatus).Find(context.TODO(), nil, &options.FindOptions{
		Limit: intPtr(int64(c.MaxDocuments)),
	})

	if err != nil {
		return nil, err
	}

	return parseList[body.TimestampedStatus](res)
}
