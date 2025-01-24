package main

import (
	// "context"

	"context"
	"encoding/json"
	"time"

	"fmt"

	"log"
	"net/http"

	servertiming "github.com/mitchellh/go-server-timing"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// IgnoreNodesOptions is options for ignoring nodes
type IgnoreNodesOptions struct {
	NodeId int    `schema:"node_id"`
	Action string `schema:"action"`
}

// IgnoreNodesResponse is a response from ignoring nodes
// It's just the new full node
type IgnoreNodesResponse Node

// IgnoreNodesHandlerWithTiming wraps our handler with
// the server timing middleware
var IgnoreNodesHandlerWithTiming = servertiming.Middleware(http.HandlerFunc(IgnoreNodesHandler), nil)

// IgnoreNodesHandler returns nodes based on an HTTP request
// without server timing headers
func IgnoreNodesHandler(w http.ResponseWriter, r *http.Request) {

	timing := servertiming.FromContext(r.Context())

	metric := timing.NewMetric("translate input for query").Start()

	err := r.ParseForm()

	if err != nil {
		log.Println("got an error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var querymap map[string][]string

	if r.Method == "GET" {
		querymap = r.URL.Query()
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

	var roptions IgnoreNodesOptions

	err = decoder.Decode(&roptions, querymap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	metric.Stop()

	if roptions.NodeId == 0 || roptions.Action == "" {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
	}

	node, err := ignoreNodes(roptions)
	if err != nil {
		log.Println("got an error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := &node
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}

}

//func decodeIgnoreNodesResponse(jsondata []byte) (*IgnoreNodesResponse, error) {
//var response IgnoreNodesResponse
//json.Unmarshal(jsondata, &response)
//return &response, nil
//}

// IgnoreNodes either ignores or unignores a node
func ignoreNodes(roptions IgnoreNodesOptions) (*IgnoreNodesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), default_timeout_seconds*time.Second)
	defer cancel()

	client, collection, err := getNodesCollection()
	if err != nil {
		log.Println("got an error:", err)
		return nil, fmt.Errorf("error finding collection: %v", err)
	}

	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Printf("Error disconnecting from db: %v\n", err)
		}
	}()

	action := roptions.Action
	var new_value bool
	if action == "ignore" {
		new_value = true
	} else if action == "unignore" {
		new_value = false
	} else {
		return nil, fmt.Errorf("Invalid action: %s", action)
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "ignored", Value: new_value}}}}

	filter := bson.D{{Key: "external_id", Value: roptions.NodeId}}

	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, fmt.Errorf("Error updating db: %w", err)
	}

	if result.MatchedCount != 1 {
		return nil, fmt.Errorf("Zero or too many matching nodes: %d", result.MatchedCount)
	}
	if result.ModifiedCount != 1 {
		return nil, fmt.Errorf("Zero or too many updated nodes: %d", result.ModifiedCount)
	}

	nodes, err := getNodes(GetNodesOptions{NodeId: roptions.NodeId, AllowIgnored: true})
	if err != nil {
		log.Println("got an error:", err)
		return nil, fmt.Errorf("error finding collection: %v", err)
	}

	node := nodes[0]
	resp := IgnoreNodesResponse(node)

	return &resp, nil
}
