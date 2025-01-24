package main

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	latlng "google.golang.org/genproto/googleapis/type/latlng"

	"ernie.org/goe/proto"
)

const test_label_template = "Test Label %v"

func TestSaveBookmark(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	randLon := get_rand_lon()
	randLat := get_rand_lat()

	now := time.Now()

	randLatLng := latlng.LatLng{Longitude: randLon, Latitude: randLat}
	geom := proto.Geometry{Coordinates: &randLatLng}

	test_label := get_test_label()
	test_request := &proto.SaveBookmarkRequest{Loc: &geom, Label: &test_label}

	test_response, err := saveBookmark(ctx, test_label, test_request)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}

	t.Log("test_response", test_response)

	responseLat := test_response.Bookmark.GetLat()
	if math.Abs(responseLat-randLat) >= .0001 {
		t.Errorf("Got different latitude in save bookmark response (%f) than sent in request (%f)", responseLat, randLat)
	}

	responseLon := test_response.Bookmark.GetLon()
	if math.Abs(responseLon-randLon) >= .0001 {
		t.Errorf("Got different latitude in save bookmark response (%f) than sent in request (%f)", responseLon, randLon)
	}

	responseLabel := test_response.Bookmark.Label
	if *responseLabel != test_label {
		t.Errorf("Got different label in save bookmark response (%s) than sent in request (%s)", *responseLabel, test_label)
	}

	responseCreationDate := test_response.Bookmark.CreationDate
	if now.Unix()-responseCreationDate.Seconds >= 10 {
		t.Errorf("Got too large diff between now (%v) and saved bookmark creation date (%v)", now, responseCreationDate)
	}

	client, collection, err := getBookmarksCollection()
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	defer client.Disconnect(ctx)

	oid, err := bson.ObjectIDFromHex(*test_response.Bookmark.Id)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	query := bson.M{"_id": oid}
	t.Log("query", query)

	resp := collection.FindOne(ctx, query)
	db_bookmark := proto.OldBookmark{}

	if err := resp.Decode(&db_bookmark); err != nil {
		t.Fatalf("got an error: %v", err)
	}

	dbLat := db_bookmark.GetLat()
	if math.Abs(dbLat-randLat) >= .0001 {
		t.Errorf("Got different latitude in save bookmark db (%f) than sent in request (%f)", dbLat, randLat)
	}

	dbLon := db_bookmark.GetLon()
	if math.Abs(dbLon-randLon) >= .0001 {
		t.Errorf("Got different latitude in save bookmark db (%f) than sent in request (%f)", dbLon, randLon)
	}

	dbLabel := db_bookmark.Label
	if *dbLabel != test_label {
		t.Errorf("Got different label in save bookmark db (%s) than sent in request (%s)", *dbLabel, test_label)
	}

	dbCreationDate := db_bookmark.CreationDate
	if now.Unix()-dbCreationDate.Seconds >= 10 {
		t.Errorf("Got too large diff between now (%v) and db bookmark creation date (%v)", now, dbCreationDate)
	}

}

func get_test_label() string {
	return fmt.Sprintf(test_label_template, r.Float64())
}
