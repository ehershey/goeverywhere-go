package main

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"ernie.org/goe/proto"
)

const test_entry_source = "Test Entry Source"

func TestSavePosition(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	randLon := r.Float64()
	randLat := r.Float64()
	now := time.Now()

	nowts := timestamppb.New(now)
	// "GET /save_position.cgi?timestamp=1732614502060
	// [ ] coords%5Blongitude%5D=-73.26432781959026
	// [ ] coords%5Blatitude%5D=41.183462092797
	// [ ] coords%5Bspeed%5D=3.779438376414303
	// [ ] coords%5Bheading%5D=292.1596296058412
	// [ ] coords%5Baltitude%5D=3.925023391842842
	// [ ] coords%5Baccuracy%5D=52.98641791661295
	// [ ] coords%5BaltitudeAccuracy%5D=3

	randLatLng := latlng.LatLng{Longitude: randLon, Latitude: randLat}
	geom := proto.Geometry{Coordinates: &randLatLng}

	point := proto.Point{Loc: &geom, EntryDate: nowts}
	test_request := &proto.SavePositionRequest{Coords: &point}

	// test_response, err := savePosition(context.Background(), test_entry_source, test_request)
	test_response, err := savePosition(ctx, test_entry_source, test_request)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}

	log.Printf("test_response: %v\n", test_response)

	// {"_id": {"$oid": "672f7c3731563ed4908fb1ec"}, "entry_date": {"$date": "2024-11-09T15:23:29.811Z"}, "loc": {"type": "Point", "coordinates": [-74.05715643152593, 40.711422093747366]}, "entry_source": "goeverywhere-local.ernie.org", "accuracy": 35.0}
	// {"_id": {"$oid": "672f7c3731563ed4908fb1ec"}
	// [ ]  "entry_date": {"$date": "2024-11-09T15:23:29.811Z"}
	// [ ]  "loc": {"type": "Point"
	// [x]  "coordinates": [-74.05715643152593
	// [x]  40.711422093747366]}
	// [ ]  "entry_source": "goeverywhere-local.ernie.org"
	// [ ]  "accuracy": 35.0}

	responseLat := test_response.GetLat()
	if responseLat != randLat {
		t.Errorf("Got different latitude in save position response (%f) than sent in request (%f)", responseLat, randLat)
	}

	responseLon := test_response.GetLon()
	if responseLon != randLon {
		t.Errorf("Got different latitude in save position response (%f) than sent in request (%f)", responseLon, randLon)
	}

	responseEntrySource := test_response.raw.SavedPoint.EntrySource
	if responseEntrySource != test_entry_source {
		t.Errorf("Got different entry_source in save position response (%s) than sent in request (%s)", responseEntrySource, test_entry_source)
	}

}

func TestSavePositionNoTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	randLon := r.Float64()
	randLat := r.Float64()

	randLatLng := latlng.LatLng{Longitude: randLon, Latitude: randLat}
	geom := proto.Geometry{Coordinates: &randLatLng}

	point := proto.Point{Loc: &geom}
	test_request := &proto.SavePositionRequest{Coords: &point}

	test_response, err := savePosition(ctx, test_entry_source, test_request)
	if err == nil || !errors.Is(err, validationError) {
		t.Fatalf("Didn't get a validation error but expected one for no timestamp passed in (test_response: %v, %v)", test_response, err)
	}
}

func TestSavePositionZeroFields(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	randLon := r.Float64()
	randLat := r.Float64()
	now := time.Now()

	nowts := timestamppb.New(now)

	randLatLng := latlng.LatLng{Longitude: randLon, Latitude: randLat}
	geom := proto.Geometry{Coordinates: &randLatLng}

	point := proto.Point{Loc: &geom, EntryDate: nowts}
	test_request := &proto.SavePositionRequest{Coords: &point}
	log.Printf("test_request: %v\n", test_request)

	// test_response, err := savePosition(context.Background(), test_entry_source, test_request)
	test_response, err := savePosition(context.Background(), test_entry_source, test_request)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}

	// query DB For new point looking for existing fields that shouldn't exist

	log.Printf("test_response: %v\n", test_response)

	responseLat := test_response.GetLat()
	if responseLat != randLat {
		t.Errorf("Got different latitude in save position response (%f) than sent in request (%f)", responseLat, randLat)
	}

	responseLon := test_response.GetLon()
	if responseLon != randLon {
		t.Errorf("Got different latitude in save position response (%f) than sent in request (%f)", responseLon, randLon)
	}

	responseEntrySource := test_response.raw.SavedPoint.EntrySource
	if responseEntrySource != test_entry_source {
		t.Errorf("Got different entry_source in save position response (%s) than sent in request (%s)", responseEntrySource, test_entry_source)
	}

	client, collection, err := getPointsCollection()
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	defer client.Disconnect(ctx)

	oid, err := primitive.ObjectIDFromHex(test_response.raw.SavedPoint.Id)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	query := bson.M{"_id": oid}
	log.Println("query:", query)

	db_point, err := collection.FindOne(ctx, query).Raw()
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	log.Println("db_point:", db_point)

	zeroFields := []string{"speed", "activity_type", "elevation", "heading", "truncate"}

	for _, field := range zeroFields {
		query[field] = bson.M{"$exists": false}
	}
	log.Println("query:", query)

	db_point, err = collection.FindOne(ctx, query).Raw()
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	log.Println("db_point:", db_point)
}
