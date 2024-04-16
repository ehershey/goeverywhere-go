package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	servertiming "github.com/mitchellh/go-server-timing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	timing := servertiming.FromContext(r.Context())

	metric := timing.NewMetric("get status").Start()

	metric.Stop()

	// for excludes:
	// https://stackoverflow.com/a/37533144/408885
	// func arrayToString(a []int, delim string) string {
	// return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	// //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	// //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
	// }
	stats, err := getStats(&StatsRequest{})
	if err != nil {
		log.Println("got an error getting stats:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)

}

func getStats(req StatsRequest) (StatsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), default_timeout_seconds*time.Second)
	defer cancel()

	client, collection, err := getNodesCollection()
	if err != nil {
		log.Println("got an error:", err)
		return nil, err
	}
	defer client.Disconnect(ctx)

	find_opts := options.FindOne()
	// find_opts.SetLimit(1)
	find_opts.SetSort(bson.D{{"entry_date", 1}})
	query := bson.M{"entry_date": bson.M{"$exists": true}}

	oldest_point_timestamp := 0

	err = collection.FindOne(ctx, query, find_opts).Decode(&oldestPoint)
	if err != nil {
		log.Println("got an error:", err)
		return nil, err
	}

	oldestPointTimestamp := oldestPoint.GetEntryDate()

	find_opts.SetSort(bson.D{{"entry_date", -1}})

	err = collection.FindOne(ctx, query, find_opts).Decode(&newestPoint)
	if err != nil {
		log.Println("got an error:", err)
		return nil, err
	}

	newestPointTimestamp := newestPoint.GetEntryDate()

	PointCount := collection.EstimatedDocumentCount(ctx)

	response = &stats_pb2.StatsResponse{
		OldestPointTimestamp: oldestPointTimestamp,
		NewestPointTimestamp: newestPointTimestamp,
		PointCount:           PointCount,
		EntrySources:         EntrySources,
	}

	var ands []bson.M

	if roptions.NodeId != 0 {
		ands = append(ands, bson.M{"external_id": roptions.NodeId})
	}

	if roptions.FromLat != 0 && roptions.FromLon != 0 && roptions.FromLat != -1 && roptions.FromLon != -1 {
		coords := make([]float64, 2)
		coords[0] = roptions.FromLon
		coords[1] = roptions.FromLat
		current_location := point{Type: "Point", Coordinates: coords}
		// var loc_doc []bson.M
		var loc_doc bson.D
		loc_doc = append(loc_doc, bson.E{Key: "$near", Value: current_location})

		if roptions.MaxDistance > 0 {
			// near_query["loc"].(map[string]interface{})["$maxDistance"] = roptions.MaxDistance
			//near_query["loc"].(bson.D)["$maxDistance"] = roptions.MaxDistance
			loc_doc = append(loc_doc, bson.E{Key: "$maxDistance", Value: roptions.MaxDistance})
		}
		near_query := bson.M{"loc": loc_doc}

		ands = append(ands, near_query)
	}

	if roptions.MinLon != 0 || roptions.MinLat != 0 || roptions.MaxLon != 0 || roptions.MaxLat != 0 {
		box_query := bson.M{"loc": bson.M{"$geoIntersects": bson.M{"$geometry": bson.M{"type": "Polygon",
			"coordinates": bson.A{bson.A{bson.A{roptions.MinLon,
				roptions.MinLat},
				bson.A{roptions.MinLon,
					roptions.MaxLat},
				bson.A{roptions.MaxLon,
					roptions.MaxLat},
				bson.A{roptions.MaxLon,
					roptions.MinLat},
				bson.A{roptions.MinLon,
					roptions.MinLat}}},
		},
		},
		},
		}

		ands = append(ands, box_query)
	}

	// Fields generally only present in db when == true
	// defaults are to return only priority nodes that aren't ignored
	if roptions.AllowIgnored == false {
		ands = append(ands, bson.M{"ignored": bson.M{"$ne": true}})
	}

	// Can make this more interesting later.. for now only even acknowledge
	// deactivated nodes exist if searching for a specific one by ID
	//
	if roptions.NodeId == 0 {
		ands = append(ands, bson.M{"deactivated": bson.M{"$ne": true}})
	}

	if roptions.RequirePriority == true {
		ands = append(ands, bson.M{"priority": true})
	}

	if len(roptions.Exclude) > 0 {
		exclude_array := make([]int64, 0)
		for _, exclude_id := range strings.Split(roptions.Exclude, "|") {
			var exclude_id_int int64
			if exclude_id_int, err = strconv.ParseInt(exclude_id, 10, 64); err != nil {
				return nil, fmt.Errorf("Error parsing exclude id into int64: %w", err)
			}
			exclude_array = append(exclude_array, exclude_id_int)
		}

		exclude_query := bson.M{"external_id": bson.M{"$nin": exclude_array}}
		ands = append(ands, exclude_query)
	}

	// query := bson.E{Key: "$and", Value: ands}
	query := bson.M{}
	if len(ands) > 0 {
		query["$and"] = ands
	}

	// log.Println("ands:", ands)
	// log.Println("query:", query)

	// query = bson.M{"$and": []bson.M{bson.M{"storyID": 123}, bson.M{"parentID": 432}}}
	find_opts := options.Find()
	find_opts.SetLimit(int64(roptions.Limit))

	var nodes []node

	cursor, err := collection.Find(ctx, query, find_opts)
	if err != nil {
		log.Println("got an error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}
