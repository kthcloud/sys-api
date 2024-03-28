package timestamp_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sys-api/pkg/db"
	"time"
)

type Client struct {
	MaxDocuments int
}

type ClientConfig struct {
	MaxDocuments int
}

func NewClient(maxDocuments ...int) *Client {
	var maxDocs int
	if len(maxDocuments) > 0 {
		maxDocs = maxDocuments[0]
	} else {
		maxDocs = 1000
	}

	return &Client{
		MaxDocuments: maxDocs,
	}
}

func (c *Client) deleteUntilNItemsLeft(collection *mongo.Collection, n int) error {
	// Fetch n'th item
	skip := int64(n - 1)

	var withTimestamp struct {
		Timestamp time.Time `bson:"timestamp"`
	}

	err := collection.FindOne(context.TODO(), bson.D{}, &options.FindOneOptions{
		Skip: &skip,
		Sort: bson.M{
			"timestamp": -1,
		},
	}).Decode(&withTimestamp)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}

		return err
	}

	// Delete all items before n'th item
	_, err = collection.DeleteMany(context.TODO(), bson.M{
		"timestamp": bson.M{
			"$lt": withTimestamp.Timestamp,
		},
	})

	return err
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

func intPtr(i int64) *int64 {
	return &i
}
