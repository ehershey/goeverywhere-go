package main

import (
	"testing"
)

func TestGetStats(t *testing.T) {
	request := StatsRequest{}
	stats, err := getStats(&request)
	if err != nil {
		t.Errorf("got an error calling getStats(&request): %v", err)
	}

	if len(stats.EntrySources) < 5 {
		t.Errorf("len(stats.EntrySources) = %d; want >=5", len(stats.EntrySources))
	}
}
