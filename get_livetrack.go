package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"git.sr.ht/~ernie/livetrackreceiver/livetrack_slack"
	"github.com/twpayne/go-polyline"
	"go.mongodb.org/mongo-driver/bson/primitive"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"ernie.org/goe/proto"
)

// type session struct {
// URL          string
// Distance     float32
// ActivityType string
// Start        time.Time
// End          time.Time
// trackPoints livetrack_slack.Livetrack_pointset
// }

func getLivetrack(ctx context.Context, req *proto.GetLivetrackRequest) (*proto.GetLivetrackResponse, error) {

	session, err := livetrack_slack.Get_session(ctx)
	if err != nil {
		wrappedErr := fmt.Errorf("got an error calling Get_session(): %w", err)
		return nil, wrappedErr
	}
	// sample_poly := "ezqlFv{luM?MEDBGAB@C?UZAl@AFAHYCaA?sECoAFc@?{@v@Ed@?RCnB@xADtAAf@E^@NCTi@LE\\?BKHEDJECPo@?AFJC?d@dABNCGE?CEB@CSA@GSGE?MB@?F@A@DDEIQ@DEAA@D@GMGBNPCKBBIG?CBFAAFBAA?BDLGAADL@GOC@@ECE@CEGCBDJB??CCIEGBHACCAD?AH@A@FDFFAFBf@vAeAqBNTCMC@AGCB?BAEFPDAGG?HBKGEA@BDCE@??J@EEE"

	url, ok := session["url"].(string)
	log.Printf("url: %v", url)
	if !ok {
		log.Printf("Weird url in session: %v", url)
	}

	start, ok := session["start"].(string)
	log.Printf("start: %v", start)
	if !ok {
		log.Printf("Weird start in session: %v", start)
	}
	end, ok := session["end"].(string)
	log.Printf("end: %v", end)
	if !ok {
		log.Printf("Weird end in session: %v", end)
	}
	layout := time.RFC3339 // "2025-01-16T20:49:18.000Z"
	start_parsed, err := time.Parse(layout, start)
	if err != nil {
		wrappedErr := fmt.Errorf("got an error parsing start time: %w", err)
		log.Printf("Weird start date in session: %v", wrappedErr)
	}
	end_parsed, err := time.Parse(layout, end)
	if err != nil {
		wrappedErr := fmt.Errorf("got an error parsing end time: %w", err)
		log.Printf("Weird end date in session: %v", wrappedErr)
	}
	start_converted := timestamppb.New(start_parsed)
	end_converted := timestamppb.New(end_parsed)

	trackPoints := session["trackPoints"]
	trackPoints_asserted, ok := trackPoints.(primitive.A)
	if !ok {
		log.Printf("Weird trackPoints in session: %v (type %T)", trackPoints, trackPoints)
	}

	trackPoints_converted := []interface{}(trackPoints_asserted)

	var distanceMeters float32 = 0.0
	activityType := "Unknown"

	all_coords := [][]float64{{}}

	for _, point := range trackPoints_converted {
		point_asserted, ok := point.(map[string]interface{})
		if !ok {
			log.Printf("Weird point in session: %v (type %T)", point, point)
		}

		position := point_asserted["position"]
		position_asserted, ok := position.(map[string]interface{})
		if !ok {
			log.Printf("Weird position in session: %v (type: %T)", position, position)
		} else {
			coords := []float64{position_asserted["lat"].(float64), position_asserted["lon"].(float64)}
			//log.Printf("coords: %v", coords)

			all_coords = append(all_coords, coords)

		}
	}
	encodedPolyBytes := polyline.EncodeCoords(all_coords)
	encoded_poly := string(encodedPolyBytes)
	if len(trackPoints_converted) > 0 {
		last_point := trackPoints_converted[len(trackPoints_converted)-1]
		log.Printf("last_point: %v", last_point)

		last_point_asserted, ok := last_point.(map[string]interface{})
		if !ok {
			log.Printf("Weird last_point in session: %v (type %T)", last_point, last_point)
		}

		fitnessPointData := last_point_asserted["fitnessPointData"]
		fitnessPointData_asserted, ok := fitnessPointData.(map[string]interface{})
		if !ok {
			log.Printf("Weird fitnessPointData in session: %v (type: %T)", fitnessPointData, fitnessPointData)
		} else {
			log.Printf("Non-weird fitnessPointData_asserted in session: %v (type: %T)", fitnessPointData_asserted, fitnessPointData)
			distanceMeters_unasserted := fitnessPointData_asserted["totalDistanceMeters"]
			distanceMeters_asserted, ok := distanceMeters_unasserted.(float64)
			if !ok {
				log.Printf("Weird distanceMeters_unasserted in session: %v (type: %T)", distanceMeters_unasserted, distanceMeters_unasserted)
			} else {
				distanceMeters = float32(distanceMeters_asserted)
			}
			activityType_unasserted := fitnessPointData_asserted["activityType"]
			activityType_asserted, ok := activityType_unasserted.(string)
			if !ok {
				log.Printf("Weird activityType_unasserted in session: %v (type: %T)", activityType_unasserted, activityType_unasserted)
			} else {
				activityType = activityType_asserted
			}
		}
	}

	resp := &proto.GetLivetrackResponse{Polyline: &encoded_poly,
		Url:          &url,
		Distance:     &distanceMeters,
		ActivityType: &activityType,
		Start:        start_converted,
		End:          end_converted}

	return resp, nil
}

type getLivetrackResponse struct {
	raw *proto.GetLivetrackResponse
}
