package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	ListenPort  int    `default:"1234"`
	MongoDB_Uri string `default:"localhost:27017"`
}

func GetConfig() (*Specification, error) {
	var s Specification
	err := envconfig.Process("goe", &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
