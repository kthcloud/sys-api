package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sys-api/models"
	"sys-api/pkg/db"
	"time"
)

func (c *Client) RegisterHost(host *models.Host) error {
	// Register host
	// Check if it already exists, if so, update it
	// Otherwise, insert it
	var current models.Host
	err := c.getCollection(db.ColHosts).FindOne(context.TODO(), bson.M{"name": host.Name}).Decode(&current)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}

		// Host does not exist, insert it
		_, err = c.getCollection(db.ColHosts).InsertOne(context.TODO(), host)
		return err
	}

	// Host exists, update it
	update := bson.M{
		"$set": bson.D{
			{"displayName", host.DisplayName},
			{"zone", host.Zone},
			{"ip", host.IP},
			{"port", host.Port},
			{"lastSeenAt", time.Now()},
		},
	}
	_, err = c.getCollection(db.ColHosts).UpdateOne(context.TODO(), bson.M{"name": host.Name}, update)
	return err
}

func (c *Client) FetchHosts() ([]models.Host, error) {
	cursor, err := c.getCollection(db.ColHosts).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	return parseList[models.Host](cursor)
}
