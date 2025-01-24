package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"iter"
	"log"
	"net/http"
	"time"

	"ernie.org/goe/proto"
	servertiming "github.com/mitchellh/go-server-timing"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/genproto/googleapis/type/latlng"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const points_collection_name = "gps_log"

func getPointsCollection() (*mongo.Client, *mongo.Collection, error) {
	return getCollectionByName(points_collection_name)
}

// GetPointsHandlerWithTiming wraps our handler with
// the server timing middleware
var GetPointsHandlerWithTiming = servertiming.Middleware(http.HandlerFunc(GetPointsHandler), nil)

// GetPointsHandler returns Points for gps_log data
// without server timing headers
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
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
	req := proto.GetPointsRequest{}
	stats, err := getPoints(ctx, &req)

	if err != nil {
		log.Printf("Got an error calling getPoints(ctx,&req): %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		fmt.Printf("Error encoding points: %v\n", err)
	}

}

func getPoints(ctx context.Context, req *proto.GetPointsRequest) (iter.Seq2[*proto.GetPointsResponse, error], error) {

	client, collection, err := getPointsCollection()
	if err != nil {
		wrappedErr := fmt.Errorf("got an error calling getPointsCollection(): %w", err)
		return nil, wrappedErr
	}

	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Printf("Error disconnecting from db: %v\n", err)
		}
	}()

	log.Println("got db client and collection ref")
	//return func(yield func(V, error) bool) {

	return func(yield func(*proto.GetPointsResponse, error) bool) {
		// ctx, cancel := context.WithTimeout(context.Background(), default_timeout_seconds*time.Second)
		// defer cancel()

		if req.GetMinLon() > req.GetMaxLon() {
			yield(nil, errors.New("min_lon must be <= max_lon"))
			return
		}
		if req.GetMinLat() > req.GetMaxLat() {
			yield(nil, errors.New("min_lat must be <= max_lat"))
			return
		}

		var ands []bson.M

		if req.GetMinLon() != 0 || req.GetMinLat() != 0 || req.GetMaxLon() != 0 || req.GetMaxLat() != 0 {
			box_query := bson.M{"loc": bson.M{"$geoIntersects": bson.M{"$geometry": bson.M{"type": "Polygon",
				"coordinates": bson.A{bson.A{bson.A{req.GetMinLon(),
					req.GetMinLat()},
					bson.A{req.GetMinLon(),
						req.GetMaxLat()},
					bson.A{req.GetMaxLon(),
						req.GetMaxLat()},
					bson.A{req.GetMaxLon(),
						req.GetMinLat()},
					bson.A{req.GetMinLon(),
						req.GetMinLat()}}},
			},
			},
			},
			}

			ands = append(ands, box_query)
		}

		query := bson.D{{Key: "MaxLat", Value: req.MaxLat}}
		cursor, err := collection.Find(ctx, query)

		if err != nil {
			wrappedErr := fmt.Errorf("got an error calling collection.FindMany(...) for points: %w", err)
			yield(nil, wrappedErr)
			return
		}
		for cursor.Next(ctx) {
			var result *gps_log_point
			if err := cursor.Decode(&result); err != nil {
				log.Fatal(err)
			}

			latLng := latlng.LatLng{Longitude: result.GetLon(), Latitude: result.GetLat()}
			geom := proto.Geometry{Coordinates: &latLng}

			point := proto.Point{Loc: &geom, EntryDate: timestamppb.New(result.GetEntryDate())}

			if !yield(&proto.GetPointsResponse{Point: &point}, nil) {
				return
			}
			if err := cursor.Err(); err != nil {
				log.Fatal(err)
			}
		}

	}, nil
}

// copied from livetrack_db.go
type gps_log_point struct {
	Id               bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EntrySource      string        `json:"entry_source" bson:"entry_source"`
	Altitude         float32       `json:"altitude,omitempty" bson:"altitude,truncate,omitempty"`
	Speed            float32       `json:"speed,omitempty" bson:"speed,omitempty"`
	EntryDate        time.Time     `json:"entry_date" bson:"entry_date"`
	Loc              geopoint      `json:"loc"`
	ActivityType     string        `json:"activityType,omitempty" bson:"activityType,omitempty"`
	Heading          float32       `json:"heading,omitempty" bson:"heading,omitempty"`
	Accuracy         float32       `json:"accuracy,omitempty" bson:"accuracy,truncate,omitempty"`
	AltitudeAccuracy float32       `json:"altitude_accuracy,omitempty" bson:"altitude_accuracy,truncate,omitempty"`
}

type geopoint struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func (point *gps_log_point) GetEntryDate() time.Time {
	return point.EntryDate
}

func (point *gps_log_point) GetLat() float64 {
	return point.Loc.GetLat()
}
func (point *gps_log_point) GetLon() float64 {
	return point.Loc.GetLon()
}
func (loc *geopoint) GetLon() float64 {
	return loc.Coordinates[0]
}

func (loc *geopoint) GetLat() float64 {
	return loc.Coordinates[1]
}
