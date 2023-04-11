package models

import (
	"context"
	"fmt"
	"landing-api/pkg/conf"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var StatsCollection *mongo.Collection
var CapacitiesCollection *mongo.Collection
var StatusCollection *mongo.Collection
var client *mongo.Client

func Setup() {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup database. details: %s", err)
	}

	// Connect to db
	uri := conf.Env.DB.URL
	clientResult, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(makeError(err))
	}
	client = clientResult

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	log.Println("successfully connected to database")

	StatsCollection = client.Database(conf.Env.DB.Name).Collection("stats")
	if err != nil {
		log.Fatalln(makeError(err))
	}
	log.Println("found collection stats")

	CapacitiesCollection = client.Database(conf.Env.DB.Name).Collection("capacities")
	if err != nil {
		log.Fatalln(makeError(err))
	}
	log.Println("found collection capacities")

	StatusCollection = client.Database(conf.Env.DB.Name).Collection("status")
	if err != nil {
		log.Fatalln(makeError(err))
	}
	log.Println("found collection status")
}

func Shutdown() {
	makeError := func(err error) error {
		return fmt.Errorf("failed to shutdown database. details: %s", err)
	}

	StatsCollection = nil
	CapacitiesCollection = nil
	StatusCollection = nil

	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatalln(makeError(err))
	}
}
