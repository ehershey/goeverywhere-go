package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"ernie.org/goe/proto"
	servertiming "github.com/mitchellh/go-server-timing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	req := proto.StatsRequest{}
	stats, err := getStats(ctx, &req)

	if err != nil {
		log.Printf("Got an error calling getStats(ctx,&req): %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)

}

func getStats(ctx context.Context, req *proto.StatsRequest) (*proto.StatsResponse, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), default_timeout_seconds*time.Second)
	// defer cancel()

	client, collection, err := getStatsCollection()
	if err != nil {
		wrappedErr := fmt.Errorf("got an error calling getStatsCollection(): %w", err)
		return nil, wrappedErr
	}

	defer client.Disconnect(ctx)

	log.Println("got db client and collection ref")

	oldest_find_opts := options.FindOne()
	// find_opts.SetLimit(1)
	empty_query := bson.D{}
	oldest_sort := empty_query // bson.D{{"entry_date", 1}}
	oldest_find_opts.SetSort(oldest_sort)
	//find_opts.SetSort(bson.D{{Key: "entry_date", Value: -1}})
	//query := bson.M{"entry_date": bson.M{"$exists": true}}
	log.Printf("empty_query: %v\n", empty_query)
	log.Printf("oldest_sort: %v\n", oldest_sort)

	var oldestPoint gps_log_point
	err = collection.FindOne(ctx, empty_query, oldest_find_opts).Decode(&oldestPoint)

	if err != nil {
		wrappedErr := fmt.Errorf("got an error calling collection.FindOne(...) for oldest point: %w", err)
		return nil, wrappedErr
	}

	log.Printf("oldestPoint: %v\n", oldestPoint)

	oldestPointTimestamp := oldestPoint.GetEntryDate()

	log.Printf("oldestPointTimestamp: %v\n", oldestPointTimestamp)

	newest_find_opts := options.FindOne()
	newest_sort := bson.D{{Key: "entry_date", Value: -1}}
	newest_find_opts.SetSort(newest_sort)

	log.Printf("newest_sort: %v\n", newest_sort)

	var newestPoint gps_log_point
	err = collection.FindOne(ctx, empty_query, newest_find_opts).Decode(&newestPoint)

	if err != nil {
		wrappedErr := fmt.Errorf("got an error calling collection.FindOne(...) for newest point: %w", err)
		return nil, wrappedErr
	}
	log.Printf("newestPoint: %v\n", newestPoint)

	newestPointTimestamp := newestPoint.GetEntryDate()

	log.Printf("newestPointTimestamp: %v\n", newestPointTimestamp)

	PointCount, err := collection.EstimatedDocumentCount(ctx)

	log.Printf("PointCount: %v\n", PointCount)

	if err != nil {
		wrappedErr := fmt.Errorf("error getting PointCount: %w", err)
		return nil, wrappedErr
	}

	log.Printf("Making distinct query with empty_query: %v\n", empty_query)
	distinct, err := collection.Distinct(ctx, "entry_source", empty_query)

	log.Printf("distinct: %v\n", distinct)

	EntrySources := []string{}

	for _, result := range distinct {
		if result != nil {
			EntrySources = append(EntrySources, result.(string))
		}
	}
	log.Printf("EntrySources: %v\n", EntrySources)

	if err != nil {
		wrappedErr := fmt.Errorf("got an error getting entry sources: %w", err)
		return nil, wrappedErr
	}

	response := &proto.StatsResponse{
		OldestPointTimestamp: timestamppb.New(oldestPointTimestamp),
		NewestPointTimestamp: timestamppb.New(newestPointTimestamp),
		PointCount:           uint32(PointCount),
		EntrySources:         EntrySources,
	}
	log.Printf("response: %v\n", response)

	return response, nil
}

type gps_log_point struct {
	Entry_source string    `json:"entry_source"`
	Altitude     float32   `json:"altitude" bson:"altitude,truncate"`
	Speed        float32   `json:"speed"`
	Entry_date   time.Time `json:"entry_date"`
	Loc          geometry  `json:"loc"`
	ActivityType string    `json:"activityType"`
	Elevation    float32   `json:"elevation,omitempty" bson:"elevation,truncate"`
}
type geometry struct{}

func (point *gps_log_point) GetEntryDate() time.Time {
	return point.Entry_date
}
