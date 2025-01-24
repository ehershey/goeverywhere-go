package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ernie.org/goe/proto"
	servertiming "github.com/mitchellh/go-server-timing"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const stats_collection_name = "gps_log"

func getStatsCollection() (*mongo.Client, *mongo.Collection, error) {
	return getCollectionByName(stats_collection_name)
}

// GetStatsHandlerWithTiming wraps our handler with
// the server timing middleware
var GetStatsHandlerWithTiming = servertiming.Middleware(http.HandlerFunc(GetStatsHandler), nil)

// GetStatsHandler returns Statsu for gps_log data
// without server timing headers
func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	timing := servertiming.FromContext(ctx)

	metric := timing.NewMetric("get status").Start()

	metric.Stop()

	// for excludes:
	// https://stackoverflow.com/a/37533144/408885
	// func arrayToString(a []int, delim string) string {
	// return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	// //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	// //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
	// }
	req := proto.GetStatsRequest{}
	stats, err := getStats(ctx, &req)

	if err != nil {
		log.Printf("Got an error calling getStats(ctx,&req): %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		fmt.Printf("Error encoding stats: %v\n", err)
	}

}

func getStats(ctx context.Context, req *proto.GetStatsRequest) (*proto.GetStatsResponse, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), default_timeout_seconds*time.Second)
	// defer cancel()

	client, collection, err := getStatsCollection()
	if err != nil {
		wrappedErr := fmt.Errorf("got an error calling getStatsCollection(): %w", err)
		return nil, wrappedErr
	}

	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Printf("Error disconnecting from db: %v\n", err)
		}
	}()

	log.Println("got db client and collection ref")

	oldest_find_opts := options.FindOne()
	empty_query := bson.D{}
	oldest_sort := bson.D{{Key: "entry_date", Value: 1}}
	oldest_find_opts.SetSort(oldest_sort)

	var oldestPoint gps_log_point
	err = collection.FindOne(ctx, empty_query, oldest_find_opts).Decode(&oldestPoint)

	if err != nil {
		wrappedErr := fmt.Errorf("got an error calling collection.FindOne(...) for oldest point: %w", err)
		return nil, wrappedErr
	}

	oldestPointTimestamp := oldestPoint.GetEntryDate()

	newest_find_opts := options.FindOne()
	newest_sort := bson.D{{Key: "entry_date", Value: -1}}
	newest_find_opts.SetSort(newest_sort)

	var newestPoint gps_log_point
	err = collection.FindOne(ctx, empty_query, newest_find_opts).Decode(&newestPoint)

	if err != nil {
		wrappedErr := fmt.Errorf("got an error calling collection.FindOne(...) for newest point: %w", err)
		return nil, wrappedErr
	}

	newestPointTimestamp := newestPoint.GetEntryDate()

	PointCount, err := collection.EstimatedDocumentCount(ctx)

	if err != nil {
		wrappedErr := fmt.Errorf("error getting PointCount: %w", err)
		return nil, wrappedErr
	}

	distinct := collection.Distinct(ctx, "entry_source", empty_query)

	err = distinct.Err()
	if err != nil {
		wrappedErr := fmt.Errorf("error making distinct call: %w", err)
		return nil, wrappedErr
	}

	EntrySources := []string{}

	err = distinct.Decode(&EntrySources)
	//for _, result := range  {
	//if result != nil {
	//EntrySources = append(EntrySources, result.(string))
	//}
	//}

	if err != nil {
		wrappedErr := fmt.Errorf("got an error getting entry sources: %w", err)
		return nil, wrappedErr
	}

	pointCount := uint32(PointCount)
	response := &proto.GetStatsResponse{
		OldestPointTimestamp: timestamppb.New(oldestPointTimestamp),
		NewestPointTimestamp: timestamppb.New(newestPointTimestamp),
		PointCount:           &pointCount,
		EntrySources:         EntrySources,
	}

	return response, nil
}
