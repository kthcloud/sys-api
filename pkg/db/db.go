package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sys-api/models"
)

var DB Context

// Context is the database context for the application.
// It contains the mongo client, as well as
// a map of all collections and their definitions.
// It is used as a singleton and should be initialized with
// the Setup() function.
type Context struct {
	MongoClient *mongo.Client

	CollectionMap           map[string]*mongo.Collection
	CollectionDefinitionMap map[string]CollectionDefinition
}

func Setup() error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to setup database. details: %s", err)
	}

	DB = Context{
		CollectionMap: make(map[string]*mongo.Collection),
	}

	var err error
	DB.MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(models.Config.MongoDB.URL))
	if err != nil {
		return makeError(err)
	}

	err = DB.MongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalln(makeError(err))
	}

	// Find collections
	DB.CollectionMap = make(map[string]*mongo.Collection)
	DB.CollectionDefinitionMap = getCollectionDefinitions()

	for _, def := range DB.CollectionDefinitionMap {
		DB.CollectionMap[def.Name] = DB.MongoClient.Database(models.Config.MongoDB.Name).Collection(def.Name)
	}

	// Create indexes
	for _, def := range DB.CollectionDefinitionMap {
		for _, indexName := range def.Indexes {
			var keys bson.D
			for _, key := range indexName {
				keys = append(keys, bson.E{Key: key, Value: 1})
			}

			_, err = DB.CollectionMap[def.Name].Indexes().CreateOne(context.Background(), mongo.IndexModel{
				Keys:    keys,
				Options: options.Index(),
			})
			if err != nil {
				return makeError(err)
			}
		}
		for _, indexName := range def.UniqueIndexes {
			var keys bson.D
			for _, key := range indexName {
				keys = append(keys, bson.E{Key: key, Value: 1})
			}

			_, err = DB.CollectionMap[def.Name].Indexes().CreateOne(context.Background(), mongo.IndexModel{
				Keys:    keys,
				Options: options.Index().SetUnique(true),
			})
			if err != nil {
				return makeError(err)
			}
		}
	}

	return nil
}
