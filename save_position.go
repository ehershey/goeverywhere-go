package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"ernie.org/goe/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func savePosition(ctx context.Context, entry_source string, req *proto.SavePositionRequest) (*savePositionResponse, error) {

	_, collection, err := getPointsCollection()
	if err != nil {
		wrappedErr := fmt.Errorf("Error getting points collection: %w", err)
		log.Println("got an error:", wrappedErr)
		return nil, wrappedErr
	}

	point := req.Coords

	if point == nil {
		err = errors.New("missing point in savePosition()")
		log.Println("got an error:", err)
		return nil, err
	}

	if err := validatePoint(point); err != nil {
		wrappedErr := fmt.Errorf("Error validating point: %w", err)
		log.Println("got an error:", wrappedErr)
		return nil, wrappedErr
	}

	point.EntrySource = entry_source
	p := savePositionPoint{raw: point}
	gps_log_point, err := p.to_gps_log_point()
	if err != nil {
		wrappedErr := fmt.Errorf("Error converting request point to db point: %w", err)
		log.Println("got an error:", wrappedErr)
		return nil, wrappedErr
	}
	result, err := collection.InsertOne(ctx, gps_log_point)
	if err != nil {
		wrappedErr := fmt.Errorf("Error inserting point: %w", err)
		log.Println("got an error:", wrappedErr)
		return nil, wrappedErr
	}
	point.Id = result.InsertedID.(primitive.ObjectID).Hex()

	proto := &proto.SavePositionResponse{Status: "OK", SavedPoint: point}
	r := savePositionResponse{raw: proto}
	return &r, nil

}

type savePositionResponse struct {
	raw *proto.SavePositionResponse
}

type savePositionPoint struct {
	raw *proto.Point
}

func validatePoint(point *proto.Point) error {
	if point.Loc == nil {
		return fmt.Errorf("Missing Loc in Point: %w", validationError)
	}
	// log.Fatalf("point.EntryDate: %v", point.EntryDate)
	if point.EntryDate == nil {
		return fmt.Errorf("Missing EntryDate in Point: %w", validationError)
	}
	MAX_LOCATION_AGE := time.Second * 86400
	now := time.Now()
	fmt.Printf("point.EntryDate: %v\n", point.EntryDate)
	pointTimestamp := point.EntryDate.AsTime()
	fmt.Printf("pointTimestamp: %v\n", pointTimestamp)
	pointTimestampAge := now.Sub(pointTimestamp)
	//oldest_possible_timestamp := now.Sub(time.Duration(int(time.Second) * MAX_LOCATION_AGE_SECS))
	//if point.EntryDate.AsTime().Before(oldest_possible_timestamp) {
	if pointTimestampAge > MAX_LOCATION_AGE {
		return fmt.Errorf("EntryDate too old in Point (%v - %v > %v): %w", now, pointTimestamp, MAX_LOCATION_AGE, validationError)
	}

	return nil
}

// copied from livetrack_db.go
func (request_point *savePositionPoint) to_gps_log_point() (*gps_log_point, error) {

	point := request_point.raw

	return &gps_log_point{
		EntrySource:      point.EntrySource,
		Altitude:         point.Altitude,
		AltitudeAccuracy: point.AltitudeAccuracy,
		Accuracy:         point.Accuracy,
		Speed:            point.Speed,
		Heading:          point.Heading,
		EntryDate:        point.EntryDate.AsTime(),
		Loc:              geopoint{Type: "Point", Coordinates: []float64{request_point.GetLon(), request_point.GetLat()}},
	}, nil
}

func (r *savePositionResponse) GetLon() float64 {
	proto := r.raw
	savedpoint := savePositionPoint{raw: proto.SavedPoint}
	return savedpoint.GetLon()
}
func (p *savePositionPoint) GetLon() float64 {
	loc := p.raw.Loc
	coords := loc.Coordinates
	lon := coords.GetLongitude()
	return lon
	//return r.raw.SavedPoint.Loc.Coordinates.GetLongitude()
}

func (r *savePositionResponse) GetLat() float64 {
	proto := r.raw
	savedpoint := savePositionPoint{raw: proto.SavedPoint}
	return savedpoint.GetLat()
}
func (p *savePositionPoint) GetLat() float64 {
	loc := p.raw.Loc
	coords := loc.Coordinates
	lat := coords.GetLatitude()
	return lat
	//return r.raw.SavedPoint.Loc.Coordinates.GetLongitude()
}
