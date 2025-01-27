package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"ernie.org/goe/proto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/genproto/googleapis/type/latlng"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func saveBookmark(ctx context.Context, label string, req *proto.SaveBookmarkRequest) (*proto.SaveBookmarkResponse, error) {

	_, collection, err := getBookmarksCollection()
	if err != nil {
		wrappedErr := fmt.Errorf("Error getting bookmarks collection: %w", err)
		log.Println("got an error:", wrappedErr)
		return nil, wrappedErr
	}

	pointtype := "Point"
	bookmark := &proto.OldBookmark{Label: &label, Loc: &proto.OldGeometry{Coordinates: []float32{float32(req.Loc.GetLon()), float32(req.Loc.GetLat())}, Type: &pointtype}, CreationDate: timestamppb.Now()}

	if err := validateBookmark(bookmark); err != nil {
		wrappedErr := fmt.Errorf("Error validating bookmark: %w", err)
		log.Println("got an error:", wrappedErr)
		return nil, wrappedErr
	}

	result, err := collection.InsertOne(ctx, bookmark)
	if err != nil {
		wrappedErr := fmt.Errorf("Error inserting : %w", err)
		log.Println("got an error:", wrappedErr)
		return nil, wrappedErr
	}
	id := result.InsertedID.(bson.ObjectID).Hex()
	bookmark.Id = &id

	newType := bookmark.GetLoc().GetType()
	newCoordinates := latlng.LatLng{Latitude: bookmark.GetLat(), Longitude: bookmark.GetLon()}
	newLoc := proto.Geometry{Type: &newType, Coordinates: &newCoordinates}
	new_bookmark := &proto.Bookmark{
		Id:           bookmark.Id,
		Label:        bookmark.Label,
		Loc:          &newLoc,
		CreationDate: bookmark.CreationDate,
	}

	r := proto.SaveBookmarkResponse{Bookmark: new_bookmark}
	return &r, nil

}

func validateBookmark(bookmark *proto.OldBookmark) error {
	if bookmark.Loc == nil {
		return fmt.Errorf("Missing Loc in Bookmark: %w", validationError)
	}
	if bookmark.CreationDate == nil {
		return fmt.Errorf("Missing CreationDate in Bookmark: %w", validationError)
	}
	MAX_LOCATION_AGE := time.Second * 86400
	now := time.Now()
	fmt.Printf("bookmark.CreationDate: %v\n", bookmark.CreationDate)
	bookmarkTimestamp := bookmark.CreationDate.AsTime()
	fmt.Printf("bookmarkTimestamp: %v\n", bookmarkTimestamp)
	bookmarkTimestampAge := now.Sub(bookmarkTimestamp)
	//oldest_possible_timestamp := now.Sub(time.Duration(int(time.Second) * MAX_LOCATION_AGE_SECS))
	//if bookmark.CreationDate.AsTime().Before(oldest_possible_timestamp) {
	if bookmarkTimestampAge > MAX_LOCATION_AGE {
		return fmt.Errorf("CreationDate too old in Bookmark (%v - %v > %v): %w", now, bookmarkTimestamp, MAX_LOCATION_AGE, validationError)
	}

	return nil
}
