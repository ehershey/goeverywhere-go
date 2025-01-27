package main

import (
	"context"
	"fmt"
	"iter"
	"log"

	"ernie.org/goe/proto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/genproto/googleapis/type/latlng"
)

const bookmarks_collection_name = "gps_bookmarks"

const default_bookmark_limit = 1000

func getBookmarksCollection() (*mongo.Client, *mongo.Collection, error) {
	return getCollectionByName(bookmarks_collection_name)
}

func getBookmarks(ctx context.Context, req *proto.GetBookmarksRequest) iter.Seq2[*proto.GetBookmarksResponse, error] {

	return func(yield func(*proto.GetBookmarksResponse, error) bool) {
		_, coll, err := getBookmarksCollection()
		if err != nil {
			wrappedErr := fmt.Errorf("error getting bookmarks collection: %w", err)
			log.Printf("Got an error: %v\n", wrappedErr)
			yield(nil, wrappedErr)
		}
		filter := bson.D{}
		opts := options.Find().SetLimit(default_bookmark_limit)
		cursor, err := coll.Find(ctx, filter, opts)
		if err != nil {
			wrappedErr := fmt.Errorf("error running Find() on bookmarks collection: %w", err)
			log.Printf("Got an error: %v\n", wrappedErr)
			yield(nil, wrappedErr)
		}
		for cursor.Next(ctx) {
			log.Printf("In cursor.Next() loop\n")
			var oldbookmark *proto.OldBookmark
			if err := cursor.Decode(&oldbookmark); err != nil {
				log.Printf("In cursor.Next() loop got err\n")
				wrappedErr := fmt.Errorf("error decoding bookmark: %w", err)
				log.Printf("Got an error: %v\n", wrappedErr)
				yield(nil, wrappedErr)
			}

			newType := oldbookmark.GetLoc().GetType()
			newCoordinates := latlng.LatLng{Latitude: oldbookmark.GetLat(), Longitude: oldbookmark.GetLon()}
			newLoc := proto.Geometry{Type: &newType, Coordinates: &newCoordinates}
			bookmark := proto.Bookmark{
				Id:           oldbookmark.Id,
				Label:        oldbookmark.Label,
				Loc:          &newLoc,
				CreationDate: oldbookmark.CreationDate,
			}

			if !yield(&proto.GetBookmarksResponse{Bookmark: &bookmark}, nil) {
				log.Printf("Returning in iterator function\n")
				return
			}
			if err := cursor.Err(); err != nil {
				wrappedErr := fmt.Errorf("error from bookmarks cursor: %w", err)
				log.Printf("Got an error: %v\n", wrappedErr)
				yield(nil, wrappedErr)
			}
		}
	}
}
