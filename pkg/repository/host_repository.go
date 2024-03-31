package repository

import (
	"context"
	"errors"
	"fmt"
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
			return fmt.Errorf("failed to find host %s. details: %w", host.Name, err)
		}

		// Host does not exist, insert it
		_, err = c.getCollection(db.ColHosts).InsertOne(context.TODO(), host)
		if err != nil {
			return fmt.Errorf("failed to insert host %s. details: %w", host.Name, err)
		}

		return nil
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
	if err != nil {
		return fmt.Errorf("failed to update host %s. details: %w", host.Name, err)
	}

	return nil
}

func (c *Client) FetchHosts() ([]models.Host, error) {
	filter := bson.D{}
	if !c.IncludeDeactivated {
		// deactivedUntil < now or deactivedUntil is not set
		filter = bson.D{{"$or", bson.A{
			bson.D{{"deactivatedUntil", bson.D{{"$lt", time.Now()}}}},
			bson.D{{"deactivatedUntil", bson.D{{"$exists", false}}}},
		}}}
	}

	cursor, err := c.getCollection(db.ColHosts).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return parseList[models.Host](cursor)
}

func (c *Client) DeactiveHost(name string, until ...time.Time) error {
	var deactivedUntil time.Time
	if len(until) > 0 {
		deactivedUntil = until[0]
	} else {
		deactivedUntil = time.Now().AddDate(1000, 0, 0) // 1000 years is sort of forever ;)
	}

	update := bson.M{
		"$set": bson.M{
			"deactivatedUntil": deactivedUntil,
		},
	}
	_, err := c.getCollection(db.ColHosts).UpdateOne(context.TODO(), bson.M{"name": name}, update)
	return err
}
