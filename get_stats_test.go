package main

import (
	"context"
	"testing"
	"time"

	"ernie.org/goe/proto"
)

func TestGetStats(t *testing.T) {
	request := proto.StatsRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stats, err := getStats(ctx, &request)
	if err != nil {
		t.Errorf("got an error calling getStats(ctx, &request): %v", err)
	}

	if len(stats.EntrySources) < 5 {
		t.Errorf("len(stats.EntrySources) = %d; want >=5", len(stats.EntrySources))
	}
}

func TestElevationFloat(t *testing.T) {
	request := proto.StatsRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stats, err := getStats(ctx, &request)
	if err != nil {
		t.Errorf("got an error calling getStats(ctx, &request): %v", err)
	}

	if len(stats.EntrySources) < 5 {
		t.Errorf("len(stats.EntrySources) = %d; want >=5", len(stats.EntrySources))
	}
}
