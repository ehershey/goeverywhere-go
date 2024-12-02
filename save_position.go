package main

import (
	"context"
	"errors"
	"fmt"
	"log"

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
	log.Printf("savePosition point: %v\n", point)

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
	fmt.Printf("gps_log_point: %v\n", gps_log_point)
	result, err := collection.InsertOne(ctx, gps_log_point)
	if err != nil {
		wrappedErr := fmt.Errorf("Error inserting point: %w", err)
		log.Println("got an error:", wrappedErr)
		return nil, wrappedErr
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	point.Id = result.InsertedID.(primitive.ObjectID).Hex()

	proto := &proto.SavePositionResponse{Status: "OK", SavedPoint: point}
	fmt.Printf("proto: %v\n", proto)
	r := savePositionResponse{raw: proto}
	fmt.Printf("r: %v\n", r)
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
	return nil
}

// copied from livetrack_db.go
func (request_point *savePositionPoint) to_gps_log_point() (*gps_log_point, error) {

	point := request_point.raw

	return &gps_log_point{
		Entry_source: point.EntrySource,
		Altitude:     point.Altitude,
		Speed:        point.Speed,
		Entry_date:   point.EntryDate.AsTime(),
		Loc:          geopoint{Type: "Point", Coordinates: []float64{request_point.GetLon(), request_point.GetLat()}},
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
