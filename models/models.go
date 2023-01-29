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

var NewsCollection *mongo.Collection
var client *mongo.Client

func Setup() {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup database. details: %s", err)
	}

	// Connect to db
	uri := conf.Env.DB.Url
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

	NewsCollection = client.Database(conf.Env.DB.Name).Collection("vms")
	if err != nil {
		log.Fatalln(makeError(err))
	}

	log.Println("found collection news")
}

func Shutdown() {
	makeError := func(err error) error {
		return fmt.Errorf("failed to shutdown database. details: %s", err)
	}

	NewsCollection = nil

	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatalln(makeError(err))
	}
}
