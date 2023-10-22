package poll

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sys-api/models"
)

func DeleteUntilNItemsLeft(collection *mongo.Collection, n int) error {

	// fetch n'th item
	skip := int64(n - 1)

	var withTimestamp models.WithTimestamp
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

	// delete all items before n'th item
	_, err = collection.DeleteMany(context.TODO(), bson.M{
		"timestamp": bson.M{
			"$lt": withTimestamp.Timestamp,
		},
	})

	return err
}
