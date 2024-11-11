package main

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	latlng "google.golang.org/genproto/googleapis/type/latlng"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"ernie.org/goe/proto"
)

const test_entry_source = "Test Entry Source"

func TestSavePosition(t *testing.T) {

	randLon := r.Float64()
	randLat := r.Float64()
	now := time.Now()

	nowts := timestamppb.New(now)
	// save_position.cgi?timestamp=752678609811&coords%5Blongitude%5D=-74.05715643152593&coords%5Blatitude%5D=40.711422093747366&coords%5Bspeed%5D=&coords%5Bheading%5D=&coords%5Baltitude%5D=&coords%5Baccuracy%5D=35&coords%5BaltitudeAccuracy%5D=
	// timestamp=752678609811
	// [ ] coords%5Blongitude%5D=-74.05715643152593
	// [ ] coords%5Blatitude%5D=40.711422093747366
	// [ ] coords%5Bspeed%5D=
	// [ ] coords%5Bheading%5D=
	// [ ] coords%5Baltitude%5D=
	// [ ] coords%5Baccuracy%5D=35
	// [ ] coords%5BaltitudeAccuracy%5D=

	randLatLng := latlng.LatLng{Longitude: randLon, Latitude: randLat}
	geom := proto.Geometry{Coordinates: &randLatLng}

	point := proto.Point{Loc: &geom, EntryDate: nowts}
	test_request := &proto.SavePositionRequest{Coords: &point}

	// test_response, err := savePosition(context.Background(), test_entry_source, test_request)
	test_response, err := savePosition(context.Background(), test_entry_source, test_request)
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

	randLon := r.Float64()
	randLat := r.Float64()

	randLatLng := latlng.LatLng{Longitude: randLon, Latitude: randLat}
	geom := proto.Geometry{Coordinates: &randLatLng}

	point := proto.Point{Loc: &geom}
	test_request := &proto.SavePositionRequest{Coords: &point}

	test_response, err := savePosition(context.Background(), test_entry_source, test_request)
	if err == nil || !errors.Is(err, validationError) {
		t.Fatalf("Didn't get a validation error but expected one for no timestamp passed in (test_response: %v, %v)", test_response, err)
	}
}
