package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	HTTPListenAddr     string `default:"127.0.0.1:1234"`
	GRPCListenAddr     string `default:"127.0.0.1:9124"`
	MongoDB_Uri        string `default:"mongodb://127.0.0.1:27017"`
	Strava_MongoDB_Uri string `default:"mongodb://127.0.0.1:27017"`
	DB_Name            string `default:"test"`
	Strava_DB_Name     string `default:"strava"`
	Error_Ntfy_Topic   string `default:""` // blank for no notifications
}

func GetConfig() (*Specification, error) {
	var s Specification
	err := envconfig.Process("goe", &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
