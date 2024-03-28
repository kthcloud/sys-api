package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sys-api/models"
	"sys-api/pkg/db"
)

func (c *Client) RegisterZone(zone *models.Zone) error {
	// Register zone
	// Check if it already exists, if so, update it
	// Otherwise, insert it
	var current models.Zone
	err := c.getCollection(db.ColZones).FindOne(context.TODO(), bson.M{"name": zone.Name}).Decode(&current)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}

		// Zone does not exist, insert it
		_, err = c.getCollection(db.ColZones).InsertOne(context.TODO(), zone)
		return err
	}

	// Zone exists, update it
	// Right now zone does not have any fields to update, but this is a placeholder for future updates
	update := bson.M{
		"$set": bson.D{},
	}
	_, err = c.getCollection(db.ColHosts).UpdateOne(context.TODO(), bson.M{"name": zone.Name}, update)
	return err
}

func (c *Client) ListZones() ([]models.Zone, error) {
	cursor, err := c.getCollection(db.ColZones).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	return parseList[models.Zone](cursor)
}
