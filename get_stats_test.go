package main

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

var r *rand.Rand

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	os.Exit(m.Run())
}

func TestGetStats(t *testing.T) {
	request := StatsRequest{}
	stats, err := getStats(request)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	if len(stats.EntrySources) < 5 {
		t.Errorf("len(stats.EntrySources) = %d; want >=5", len(nodes))
	}
}
