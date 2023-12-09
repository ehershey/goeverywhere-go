package main

import (
	// "context"
	"crypto/md5"
	"encoding/json"

	// "errors"
	"fmt"
	// "github.com/gorilla/schema"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"

	servertiming "github.com/mitchellh/go-server-timing"
	// "strconv"
	// "strings"
	// "time"
)

// RefreshNodesOptions is options for refreshing nodes
type RefreshNodesOptions struct {
	MinLon      float64 `schema:"min_lon"`
	MinLat      float64 `schema:"min_lat"`
	MaxLon      float64 `schema:"max_lon"`
	MaxLat      float64 `schema:"max_lat"`
	BoundString string  `json:"bound_string"`
}

// RefreshNodesResponse is a response from refreshing nodes
// It mostly contains the input parameters.
// It also contains a job id for callers to check progress.
type RefreshNodesResponse struct {
	MinLon      float64 `json:"min_lon"`
	MinLat      float64 `json:"min_lat"`
	MaxLon      float64 `json:"max_lon"`
	MaxLat      float64 `json:"max_lat"`
	BoundString string  `json:"bound_string"`
	Rid         string  `json:"rid"`
	JobID       string  `json:"job_id"`
}

const maxSpan = .3

// RefreshNodesHandlerWithTiming wraps our handler with
// the server timing middleware
var RefreshNodesHandlerWithTiming = servertiming.Middleware(http.HandlerFunc(RefreshNodesHandler), nil)

// RefreshNodesHandler returns nodes based on an HTTP request
// without server timing headers
func RefreshNodesHandler(w http.ResponseWriter, r *http.Request) {

	timing := servertiming.FromContext(r.Context())

	metric := timing.NewMetric("translate input for query").Start()

	err := r.ParseForm()

	if err != nil {
		log.Println("got an error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var querymap map[string][]string

	if r.Method == "GET" {
		querymap = r.URL.Query()
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

	var roptions RefreshNodesOptions

	err = decoder.Decode(&roptions, querymap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	metric.Stop()

	if roptions.MinLat == 0 && roptions.MaxLat == 0 && roptions.MinLon == 0 && roptions.MaxLon == 0 {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
	}

	// account for swapped mins and maxes
	if roptions.MinLat >= roptions.MaxLat {
		roptions.MinLat, roptions.MaxLat = roptions.MaxLat, roptions.MinLat
	}
	if roptions.MinLon >= roptions.MaxLon {
		roptions.MinLon, roptions.MaxLon = roptions.MaxLon, roptions.MinLon
	}

	if roptions.MaxLat-roptions.MinLat >= maxSpan {
		http.Error(w, "Latitude span is too large", http.StatusBadRequest)
	}

	if roptions.MaxLon-roptions.MinLon >= maxSpan {
		http.Error(w, "Longitude span is too large", http.StatusBadRequest)
	}

	// for excludes:
	// https://stackoverflow.com/a/37533144/408885
	// func arrayToString(a []int, delim string) string {
	// return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	// //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	// //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
	// }
	jobid, err := RefreshNodes(roptions)
	if err != nil {
		log.Println("got an error:", err)
		return
	}

	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%s%f%f%f%f", roptions.BoundString, roptions.MinLon, roptions.MaxLon, roptions.MinLat, roptions.MaxLat))
	requestHash := fmt.Sprintf("%x", h.Sum(nil))

	if err != nil {
		log.Println("got an error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := &RefreshNodesResponse{
		MinLon:      roptions.MinLon,
		MinLat:      roptions.MinLat,
		MaxLon:      roptions.MaxLon,
		MaxLat:      roptions.MaxLat,
		BoundString: roptions.BoundString,
		Rid:         requestHash,
		JobID:       jobid,
	}

	json.NewEncoder(w).Encode(response)

}

func decodeRefreshNodesResponse(jsondata []byte) (*RefreshNodesResponse, error) {
	var response RefreshNodesResponse
	json.Unmarshal(jsondata, &response)
	return &response, nil
}

// RefreshNodes kicks off a node refresh with the given options
// it returns a jobid
func RefreshNodes(roptions RefreshNodesOptions) (string, error) {
	jobid := fmt.Sprintf("%f / %f / %f / %f", roptions.MinLon,
		roptions.MinLat,
		roptions.MaxLon,
		roptions.MaxLat)
	return jobid, nil
}
