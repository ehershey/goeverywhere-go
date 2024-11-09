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
