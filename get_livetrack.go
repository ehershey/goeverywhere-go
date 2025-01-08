package main

import (
	"context"

	"ernie.org/goe/proto"
)

func getLivetrack(ctx context.Context, req *proto.GetLivetrackRequest) (*proto.GetLivetrackResponse, error) {

	sample_poly := "ezqlFv{luM?MEDBGAB@C?UZAl@AFAHYCaA?sECoAFc@?{@v@Ed@?RCnB@xADtAAf@E^@NCTi@LE\\?BKHEDJECPo@?AFJC?d@dABNCGE?CEB@CSA@GSGE?MB@?F@A@DDEIQ@DEAA@D@GMGBNPCKBBIG?CBFAAFBAA?BDLGAADL@GOC@@ECE@CEGCBDJB??CCIEGBHACCAD?AH@A@FDFFAFBf@vAeAqBNTCMC@AGCB?BAEFPDAGG?HBKGEA@BDCE@??J@EEE"
	resp := &proto.GetLivetrackResponse{Polyline: &sample_poly}

	return resp, nil
}

type getLivetrackResponse struct {
	raw *proto.GetLivetrackResponse
}
