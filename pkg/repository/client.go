package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"sys-api/pkg/db"
)

type Client struct {
	IncludeDeactivated bool
}

func NewClient() *Client {
	return &Client{
		IncludeDeactivated: false,
	}
}

func (c *Client) WithDeactivated() *Client {
	c.IncludeDeactivated = true
	return c
}

func parseList[T any](cursor *mongo.Cursor) ([]T, error) {
	var res []T
	err := cursor.All(context.Background(), &res)
	if err != nil {
		return nil, nil
	}
	return res, nil
}

func (c *Client) getCollection(name string) *mongo.Collection {
	collection, ok := db.DB.CollectionMap[name]
	if !ok {
		panic("collection not found")
	}

	return collection
}
