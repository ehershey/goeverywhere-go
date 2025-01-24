package main

import (
	"context"
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

func getBookmarks(ctx context.Context, req *proto.GetBookmarksRequest) iter.Seq[*proto.GetBookmarksResponse] {

	return func(yield func(*proto.GetBookmarksResponse) bool) {
		_, coll, err := getBookmarksCollection()
		if err != nil {
			panic(err)
		}
		filter := bson.D{}
		opts := options.Find().SetLimit(default_bookmark_limit)
		cursor, err := coll.Find(ctx, filter, opts)
		if err != nil {
			panic(err)
		}
		for cursor.Next(ctx) {
			var oldbookmark *proto.OldBookmark
			if err := cursor.Decode(&oldbookmark); err != nil {
				log.Fatal(err)
			}
			log.Printf("oldbookmark.CreationDate: %v", oldbookmark.CreationDate)

			newType := oldbookmark.GetLoc().GetType()
			newCoordinates := latlng.LatLng{Latitude: oldbookmark.GetLat(), Longitude: oldbookmark.GetLon()}
			newLoc := proto.Geometry{Type: &newType, Coordinates: &newCoordinates}
			bookmark := proto.Bookmark{
				Id:           oldbookmark.Id,
				Label:        oldbookmark.Label,
				Loc:          &newLoc,
				CreationDate: oldbookmark.CreationDate,
			}
			log.Printf("bookmark.CreationDate: %v", bookmark.CreationDate)

			if !yield(&proto.GetBookmarksResponse{Bookmark: &bookmark}) {
				return
			}
			if err := cursor.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
