package main

import (
	"context"
	"iter"
	"log"

	"ernie.org/goe/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const strava_activities_collection_name = "activities"

const default_polyline_limit = 2000

//const default_timeout_seconds = 60

func getPolylinesCollection() (*mongo.Client, *mongo.Collection, error) {
	return getCollectionByName(strava_activities_collection_name)
}

//func (yield func(Item) bool)

type StravaActivity struct {
	MapPolyline string `bson:"map_polyline" json:"map_polyline"`
}

func getPolylines(ctx context.Context, req *proto.GetPolylinesRequest) iter.Seq[*proto.GetPolylinesResponse] {

	return func(yield func(*proto.GetPolylinesResponse) bool) {
		_, coll, err := getPolylinesCollection()
		if err != nil {
			panic(err)
		}
		filter := bson.D{}
		opts := options.Find().SetLimit(default_polyline_limit)
		cursor, err := coll.Find(ctx, filter, opts)
		if err != nil {
			panic(err)
		}
		for cursor.Next(ctx) {
			var result *StravaActivity
			if err := cursor.Decode(&result); err != nil {
				log.Fatal(err)
			}
			one_activity_polyline := &proto.ActivityPolyline{Polyline: &result.MapPolyline}
			var activity_polyline_array []*proto.ActivityPolyline
			activity_polyline_array = append(activity_polyline_array, one_activity_polyline)
			//activity_polyline_array[0] = one_activity_polyline
			if !yield(&proto.GetPolylinesResponse{Polylines: activity_polyline_array}) {
				return
			}
			if err := cursor.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
