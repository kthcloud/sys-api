package timestamp_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sys-api/dto/body"
	"sys-api/pkg/db"
)

func (c *Client) SaveGpuInfo(gpuInfo *body.TimestampedGpuInfo) error {
	_, err := c.getCollection(db.ColGpuInfo).InsertOne(context.TODO(), *gpuInfo)
	if err != nil {
		return err
	}

	err = c.deleteUntilNItemsLeft(c.getCollection(db.ColGpuInfo), c.MaxDocuments)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) FetchGpuInfo() ([]body.TimestampedGpuInfo, error) {
	res, err := c.getCollection(db.ColGpuInfo).Find(context.TODO(), nil, &options.FindOptions{
		Limit: intPtr(int64(c.MaxDocuments)),
	})

	if err != nil {
		return nil, err
	}

	return parseList[body.TimestampedGpuInfo](res)
}
