package main

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"log"
	"time"

	"ernie.org/goe/proto"
	"github.com/getsentry/sentry-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/genproto/googleapis/type/latlng"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const points_collection_name = "gps_log"

// # minimum time span in milliseconds between points for new point to be included
const MINIMUM_POINT_DELTA_MILLIS = 5000

// # minimum distance in meters between points new point to be included
const MINIMUM_POINT_DELTA_METERS = 40

const QUERY_DEFAULT_LIMIT = 20000

func getPointsCollection() (*mongo.Client, *mongo.Collection, error) {
	return getCollectionByName(points_collection_name)
}

func getPoints(ctx context.Context, req *proto.GetPointsRequest) iter.Seq2[*proto.GetPointsResponse, error] {
	//return func(yield func(V, error) bool) {
	defer func() {
		err := recover()

		if err != nil {
			log.Printf("got error in defered getPoints() func\n")
			log.Printf("got error in defered getPoints() func: %v\n", err)
			sentry.CurrentHub().Recover(err)
			sentry.Flush(time.Second * 5)
		}
	}()

	return func(yield func(*proto.GetPointsResponse, error) bool) {
		log.Printf("getPoints() in yield func\n")
		defer func() {
			err := recover()

			if err != nil {
				log.Printf("got error in defered yield() func\n")
				log.Printf("got error in defered yield() func: %v\n", err)
				sentry.CurrentHub().Recover(err)
				sentry.Flush(time.Second * 5)
			}
		}()

		client, collection, err := getPointsCollection()
		if err != nil {
			wrappedErr := fmt.Errorf("got an error calling getPointsCollection(): %w", err)
			yield(nil, wrappedErr)
		}
		log.Printf("getPoints() got db client and collection ref\n")

		defer func() {
			err := client.Disconnect(ctx)
			if err != nil {
				fmt.Printf("Error disconnecting from db: %v\n", err)
			}
		}()

		if req.GetMinLon() > req.GetMaxLon() {
			yield(nil, errors.New("min_lon must be <= max_lon"))
			return
		}
		if req.GetMinLat() > req.GetMaxLat() {
			yield(nil, errors.New("min_lat must be <= max_lat"))
			return
		}

		log.Printf("getPoints() got past minmax checks\n")
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
		log.Printf("getPoints() got past minmax config\n")

		log.Printf("req.GetFrom(): %v\n", req.GetFrom())
		log.Printf("req.GetTo(): %v\n", req.GetTo())
		if req.GetFrom() != nil && req.GetTo() != nil {
			time_query := bson.M{"entry_date": bson.M{"$gte": req.GetFrom(), "$lte": req.GetTo()}}
			ands = append(ands, time_query)
		}

		query := bson.M{}
		if len(ands) > 0 {
			query["$and"] = ands
		}

		// E QUERY_DEFAULT_LIMIT = 20000 # E: Expected a { to open the function definition.
		//default_point_limit := 20000
		default_point_limit := 2000
		limit := int64(default_point_limit)
		if req.Limit != nil {
			limit = int64(req.GetLimit())
		}
		sort := bson.M{"entry_date": 1}
		opts := options.Find().SetLimit(limit).SetSort(sort)
		log.Printf("getPoints() query: %v\n", query)
		log.Printf("getPoints() sort: %v\n", sort)
		log.Printf("getPoints() opts: %v\n", opts)
		cursor, err := collection.Find(ctx, query, opts)
		log.Printf("getPoints() ran Find()\n")

		if err != nil {
			wrappedErr := fmt.Errorf("got an error calling collection.FindMany(...) for points: %w", err)
			yield(nil, wrappedErr)
			return
		}
		log.Printf("getPoints() checked Find() error\n")
		for cursor.Next(ctx) {
			var result *gps_log_point
			if err := cursor.Decode(&result); err != nil {
				wrappedErr := fmt.Errorf("got an error decoding from cursor: %w", err)
				log.Printf("Got an error: %v", wrappedErr)
				yield(nil, wrappedErr)
			}

			latLng := latlng.LatLng{Longitude: result.GetLon(), Latitude: result.GetLat()}
			geom := proto.Geometry{Coordinates: &latLng}

			point := proto.Point{Loc: &geom, EntryDate: timestamppb.New(result.GetEntryDate()),
				EntrySource:  &result.EntrySource,
				ActivityType: &result.ActivityType,
			}

			if !yield(&proto.GetPointsResponse{Point: &point}, nil) {
				log.Printf("yield returned false - returning\n")
				return
			}
			if err := cursor.Err(); err != nil {
				wrappedErr := fmt.Errorf("got an error from cursor: %w", err)
				log.Printf("Got an error: %v", wrappedErr)
				yield(nil, wrappedErr)
			}
		}

	}
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
