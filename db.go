package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func getCollectionByName(collection_name string) (*mongo.Client, *mongo.Collection, error) {

	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoDB_Uri))
	if err != nil {
		log.Println("got an error:", err)
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	collection := client.Database(config.DB_Name).Collection(collection_name)
	return client, collection, nil
}
