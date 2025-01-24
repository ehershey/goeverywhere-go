package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"ernie.org/goe/proto"
)

const test_entry_source_template = "Test Entry Source %v"

func TestSavePosition(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	randLon := get_rand_lon()
	randLat := get_rand_lat()

	randAccuracy := r.Float32() + float32(r.Int())
	randSpeed := r.Float32() + float32(r.Int())
	randHeading := r.Float32() + float32(r.Int())
	randAltitude := r.Float32() + float32(r.Int())
	randAltitudeAccuracy := r.Float32() + float32(r.Int())
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

	point := proto.Point{Loc: &geom, EntryDate: nowts,

		Accuracy:         &randAccuracy,
		Speed:            &randSpeed,
		Heading:          &randHeading,
		Altitude:         &randAltitude,
		AltitudeAccuracy: &randAltitudeAccuracy,
	}
	test_request := &proto.SavePositionRequest{Coords: &point}

	test_entry_source := get_test_entry_source()
	test_response, err := savePosition(ctx, test_entry_source, test_request)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}

	t.Log("test_response", test_response)

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
	if *responseEntrySource != test_entry_source {
		t.Errorf("Got different entry_source in save position response (%s) than sent in request (%s)", *responseEntrySource, test_entry_source)
	}

	//Accuracy:         randAccuracy,
	//Speed:            randSpeed,
	//Heading:          randHeading,
	//Altitude:         randAltitude,
	//AltitudeAccuracy: randAltitudeAccuracy,

	responseAccuracy := test_response.raw.SavedPoint.GetAccuracy()
	if responseAccuracy != randAccuracy {
		t.Errorf("Got different Accuracy in save position response (%f) than sent in request (%f)", responseAccuracy, randAccuracy)
	} else {
		t.Logf("Got same Accuracy in save position response (%f) as sent in request (%f)", responseAccuracy, randAccuracy)
	}
	responseSpeed := test_response.raw.SavedPoint.GetSpeed()
	if responseSpeed != randSpeed {
		t.Errorf("Got different Speed in save position response (%f) than sent in request (%f)", responseSpeed, randSpeed)
	}
	responseHeading := test_response.raw.SavedPoint.GetHeading()
	if responseHeading != randHeading {
		t.Errorf("Got different Heading in save position response (%f) than sent in request (%f)", responseHeading, randHeading)
	}
	responseAltitude := test_response.raw.SavedPoint.GetAltitude()
	if responseAltitude != randAltitude {
		t.Errorf("Got different Altitude in save position response (%f) than sent in request (%f)", responseAltitude, randAltitude)
	}
	responseAltitudeAccuracy := test_response.raw.SavedPoint.GetAltitudeAccuracy()
	if responseAltitudeAccuracy != randAltitudeAccuracy {
		t.Errorf("Got different AltitudeAccuracy in save position response (%f) than sent in request (%f)", responseAltitudeAccuracy, randAltitudeAccuracy)
	}

	client, collection, err := getPointsCollection()
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	defer client.Disconnect(ctx)

	oid, err := bson.ObjectIDFromHex(*test_response.raw.SavedPoint.Id)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	query := bson.M{"_id": oid}
	t.Log("query", query)

	resp := collection.FindOne(ctx, query)
	db_point, err := resp.Raw()
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	t.Log("db_point", db_point)

	db_gps_log_point := gps_log_point{}
	if err := resp.Decode(&db_gps_log_point); err != nil {
		t.Fatalf("got an error: %v", err)
	}
	db_point_obj := &db_gps_log_point
	t.Log("db_point_obj", &db_point_obj)

	dbLat := db_point_obj.GetLat()
	if dbLat != randLat {
		t.Errorf("Got different latitude in save position db (%f) than sent in request (%f)", dbLat, randLat)
	}

	dbLon := db_point_obj.GetLon()
	if dbLon != randLon {
		t.Errorf("Got different latitude in save position db (%f) than sent in request (%f)", dbLon, randLon)
	}

	dbEntrySource := db_point_obj.EntrySource
	if dbEntrySource != test_entry_source {
		t.Errorf("Got different entry_source in save position db (%s) than sent in request (%s)", dbEntrySource, test_entry_source)
	}

	dbAccuracy := db_point_obj.Accuracy
	if dbAccuracy != randAccuracy {
		t.Errorf("Got different Accuracy in save position db (%f) than sent in request (%f)", dbAccuracy, randAccuracy)
	} else {
		t.Logf("Got same Accuracy in save position db (%f) as sent in request (%f)", dbAccuracy, randAccuracy)
	}
	dbSpeed := db_point_obj.Speed
	if dbSpeed != randSpeed {
		t.Errorf("Got different Speed in save position db (%f) than sent in request (%f)", dbSpeed, randSpeed)
	}
	dbHeading := db_point_obj.Heading
	if dbHeading != randHeading {
		t.Errorf("Got different Heading in save position db (%f) than sent in request (%f)", dbHeading, randHeading)
	}
	dbAltitude := db_point_obj.Altitude
	if dbAltitude != randAltitude {
		t.Errorf("Got different Altitude in save position db (%f) than sent in request (%f)", dbAltitude, randAltitude)
	}
	dbAltitudeAccuracy := db_point_obj.AltitudeAccuracy
	if dbAltitudeAccuracy != randAltitudeAccuracy {
		t.Errorf("Got different AltitudeAccuracy in save position db (%f) than sent in request (%f)", dbAltitudeAccuracy, randAltitudeAccuracy)
	}

}

func TestSavePositionNoTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	randLon := get_rand_lon()
	randLat := get_rand_lat()

	randLatLng := latlng.LatLng{Longitude: randLon, Latitude: randLat}
	geom := proto.Geometry{Coordinates: &randLatLng}

	point := proto.Point{Loc: &geom}
	test_request := &proto.SavePositionRequest{Coords: &point}

	test_entry_source := get_test_entry_source()
	test_response, err := savePosition(ctx, test_entry_source, test_request)
	if err == nil || !errors.Is(err, validationError) {
		t.Fatalf("Didn't get a validation error but expected one for no timestamp passed in (test_response: %v, %v)", test_response, err)
	}
}

func TestSavePositionOldTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	randLon := get_rand_lon()
	randLat := get_rand_lat()

	randLatLng := latlng.LatLng{Longitude: randLon, Latitude: randLat}
	geom := proto.Geometry{Coordinates: &randLatLng}

	testOldDate := "Jan 2, 2006 at 3:04pm (MST)"
	oldTime, err := time.Parse(testOldDate, testOldDate)
	if err != nil {
		t.Fatalf("Error parsing test date (%s): %v", testOldDate, err)
	}
	point := proto.Point{Loc: &geom, EntryDate: timestamppb.New(oldTime)}

	test_request := &proto.SavePositionRequest{Coords: &point}

	test_entry_source := get_test_entry_source()
	test_response, err := savePosition(ctx, test_entry_source, test_request)
	if err == nil || !errors.Is(err, validationError) {
		t.Fatalf("Didn't get a validation error but expected one for bad old timestamp passed in (test_response: %v, %v)", test_response, err)
	}
}

func TestSavePositionZeroFields(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	randLon := get_rand_lon()
	randLat := get_rand_lat()
	now := time.Now()

	nowts := timestamppb.New(now)

	randLatLng := latlng.LatLng{Longitude: randLon, Latitude: randLat}
	geom := proto.Geometry{Coordinates: &randLatLng}

	point := proto.Point{Loc: &geom, EntryDate: nowts}
	test_request := &proto.SavePositionRequest{Coords: &point}
	t.Log("test_request", test_request)

	test_entry_source := get_test_entry_source()
	test_response, err := savePosition(context.Background(), test_entry_source, test_request)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}

	// query DB For new point looking for existing fields that shouldn't exist

	t.Log("test_response", test_response)

	responseLat := test_response.GetLat()
	if responseLat != randLat {
		t.Errorf("Got different latitude in save position response (%f) than sent in request (%f)", responseLat, randLat)
	}

	responseLon := test_response.GetLon()
	if responseLon != randLon {
		t.Errorf("Got different latitude in save position response (%f) than sent in request (%f)", responseLon, randLon)
	}

	responseEntrySource := test_response.raw.SavedPoint.EntrySource
	if *responseEntrySource != test_entry_source {
		t.Errorf("Got different entry_source in save position response (%s) than sent in request (%s)", *responseEntrySource, test_entry_source)
	}

	client, collection, err := getPointsCollection()
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	defer client.Disconnect(ctx)

	oid, err := bson.ObjectIDFromHex(*test_response.raw.SavedPoint.Id)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	query := bson.M{"_id": oid}
	t.Log("query", query)

	db_point, err := collection.FindOne(ctx, query).Raw()
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	t.Log("db_point", db_point)

	zeroFields := []string{"speed", "activity_type", "altitude", "heading", "truncate"}

	for _, field := range zeroFields {
		query[field] = bson.M{"$exists": false}
	}
	t.Log("query", query)

	db_point, err = collection.FindOne(ctx, query).Raw()
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
	t.Log("db_point", db_point)
}

func get_test_entry_source() string {
	return fmt.Sprintf(test_entry_source_template, r.Float64())
}

func get_rand_lon() float64 {
	return r.Float64() + float64(r.Intn(175))
}
func get_rand_lat() float64 {
	return r.Float64() + float64(r.Intn(85))
}
