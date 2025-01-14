package main

import (
	"context"
	"fmt"
	"log"

	"git.sr.ht/~ernie/livetrackreceiver/livetrack_slack"

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
	sample_poly := "ezqlFv{luM?MEDBGAB@C?UZAl@AFAHYCaA?sECoAFc@?{@v@Ed@?RCnB@xADtAAf@E^@NCTi@LE\\?BKHEDJECPo@?AFJC?d@dABNCGE?CEB@CSA@GSGE?MB@?F@A@DDEIQ@DEAA@D@GMGBNPCKBBIG?CBFAAFBAA?BDLGAADL@GOC@@ECE@CEGCBDJB??CCIEGBHACCAD?AH@A@FDFFAFBf@vAeAqBNTCMC@AGCB?BAEFPDAGG?HBKGEA@BDCE@??J@EEE"

	url, ok := session["url"].(string)
	log.Printf("Weird url in session: %v", url)
	if !ok {
		log.Printf("Weird url in session: %v", url)
	}
	resp := &proto.GetLivetrackResponse{Polyline: &sample_poly,
		Url: &url}

	return resp, nil
}

type getLivetrackResponse struct {
	raw *proto.GetLivetrackResponse
}
