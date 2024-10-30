package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const STRAVA_COLLECTION_NAME = "activities"

func getCollectionByName(collection_name string) (*mongo.Client, *mongo.Collection, error) {

	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	uri := config.MongoDB_Uri
	db_name := config.DB_Name
	if collection_name == STRAVA_COLLECTION_NAME {
		uri = config.Strava_MongoDB_Uri
		db_name = config.Strava_DB_Name
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		wrappedErr := fmt.Errorf("Error creating mongodb client: %w", err)
		return nil, nil, wrappedErr
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		wrappedErr := fmt.Errorf("Error connecting to mongodb: %w", err)
		return nil, nil, wrappedErr
	}
	collection := client.Database(db_name).Collection(collection_name)
	return client, collection, nil
}
