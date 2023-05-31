package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const kv_collection_name = "goekv"

type kv struct {
	K string `json:"k" schema:"k"`
	V string `json:"v" schema:"v"`
}

type KeyValueOptions kv

type keyValueResponse kv

type keyValueEntry kv

func KeyValueHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	var querymap map[string][]string

	if r.Method == "GET" {
		querymap = r.URL.Query()
	} else {
		querymap = r.PostForm
	}

	if err != nil {
		log.Println("got an error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var roptions KeyValueOptions

	err = decoder.Decode(&roptions, querymap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	v, err := getKV(roptions)
	if err != nil {
		log.Println("got an error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := &keyValueResponse{
		K: roptions.K,
		V: v,
	}

	json.NewEncoder(w).Encode(response)

}

func getKV(roptions KeyValueOptions) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, collection, err := getCollectionByName(kv_collection_name)
	if err != nil {
		log.Println("got an error:", err)
		return "", err
	}
	defer client.Disconnect(ctx)

	query := bson.M{"k": roptions.K}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "v", Value: roptions.V}}}}

	var value string

	if roptions.V != "" {
		update_opts := options.Update().SetUpsert(true)
		value = roptions.V
		_, err := collection.UpdateOne(ctx, query, update, update_opts)
		if err != nil {
			log.Println("got an error:", err)
			return "", err
		}

	} else {
		find_opts := options.FindOne()
		var result keyValueEntry
		err = collection.FindOne(ctx, query, find_opts).Decode(&result)
		value = result.V
		if err != nil {
			// ErrNoDocuments means that the filter did not match any documents in the collection
			if err == mongo.ErrNoDocuments {
				return "", nil
			}
			log.Println("got an error:", err)
			return "", err
		}
	}

	return value, nil
}
