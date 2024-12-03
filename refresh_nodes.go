package main

import (
	"context"
	"encoding/json"
	"time"

	"fmt"

	"log"
	"net/http"

	servertiming "github.com/mitchellh/go-server-timing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const jobs_collection_name = "jobs"

func getJobsCollection() (*mongo.Client, *mongo.Collection, error) {
	return getCollectionByName(jobs_collection_name)
}

type RefreshNodesOptions struct {
	Action string  `schema:"action"`
	JobId  string  `schema:"job_id"`
	MinLon float64 `schema:"min_lon"`
	MinLat float64 `schema:"min_lat"`
	MaxLon float64 `schema:"max_lon"`
	MaxLat float64 `schema:"max_lat"`
}

type RefreshNodesJob struct {
	// automatic _id generated by mongodb
	Id primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	// "manual" id semi-readable
	JobId      string    `json:"job_id" bson:"job_id"`
	Status     string    `json:"job_status" bson:"job_status"`
	StatusTime time.Time `json:"status_time" bson:"status_time"`
	Output     string    `json:"output" bson:"output"`
	MinLon     float64   `json:"min_lon" bson:"min_lon"`
	MinLat     float64   `json:"min_lat" bson:"min_lat"`
	MaxLon     float64   `json:"max_lon" bson:"max_lon"`
	MaxLat     float64   `json:"max_lat" bson:"max_lat"`
}

// RefreshNodesResponse is a response to clients
// It contains a job id and status text
type RefreshNodesResponse struct {
	JobId     string `json:"job_id"`
	JobStatus string `json:"job_status"`
	JobOutput string `json:"job_output,omitempty"`
}

var validPreviousStatusByNewStatus = map[string]string{
	"queued":    "",
	"started":   "queued",
	"failed":    "started",
	"completed": "started",
}

func (job *RefreshNodesJob) SetStatus(new_status string) error {
	if new_status != "queued" && new_status != "started" && new_status != "failed" && new_status != "completed" {
		return fmt.Errorf("Invalid status for job: %s", new_status)
	}
	old_status := job.Status
	if old_status != validPreviousStatusByNewStatus[new_status] {
		return fmt.Errorf("Invalid previous status for new status: old: %s, new: %s", old_status, new_status)
	}
	job.StatusTime = time.Now()
	job.Status = new_status
	return nil
}

func (job *RefreshNodesJob) GetStatus() string {
	return job.Status
}

func (job *RefreshNodesJob) GetStatusTime() time.Time {
	return job.StatusTime
}

// type Job interface {
// SetStatus(status string) error
// GetStatus() string
// GetStatusTime() time.Time
// }

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
		wrappedErr := fmt.Errorf("Error parsing input parameters: %w", err)
		log.Println("got an error:", wrappedErr)
		http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
		return
	}

	var querymap map[string][]string

	if r.Method == "GET" {
		querymap = r.URL.Query()
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var roptions RefreshNodesOptions

	err = decoder.Decode(&roptions, querymap)
	if err != nil {
		wrappedErr := fmt.Errorf("Error decoding input parameters: %w", err)
		log.Println("got an error:", wrappedErr)
		http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
		return
	}
	metric.Stop()

	if roptions.Action != "start" && roptions.Action != "check" {
		http.Error(w, "Invalid or missing action parameter", http.StatusBadRequest)
		return
	}

	response := &RefreshNodesResponse{}

	if roptions.Action == "check" {
		if roptions.JobId == "" {
			http.Error(w, "Invalid or missing job id parameter", http.StatusBadRequest)
			return
		}
		job, err := GetJob(roptions.JobId)
		if err != nil {
			wrappedErr := fmt.Errorf("Error getting job: %w", err)
			log.Println("got an error:", wrappedErr)
			http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
			return
		}
		response.JobStatus = job.Status
		response.JobId = roptions.JobId
		response.JobOutput = job.Output
	}
	if roptions.Action == "start" {
		if roptions.MinLat == 0 && roptions.MaxLat == 0 && roptions.MinLon == 0 && roptions.MaxLon == 0 {
			http.Error(w, "Invalid or missing coordinate parameters", http.StatusBadRequest)
			return
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
			return
		}

		if roptions.MaxLon-roptions.MinLon >= maxSpan {
			http.Error(w, "Longitude span is too large", http.StatusBadRequest)
			return
		}

		jobid, err := RefreshNodes(roptions)
		if err != nil {
			wrappedErr := fmt.Errorf("Error creating job: %w", err)
			log.Println("got an error:", wrappedErr)
			http.Error(w, wrappedErr.Error(), http.StatusBadRequest)
			return
		}

		response.JobId = jobid
		response.JobStatus = "queued"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}

}

// func decodeRefreshNodesResponse(jsondata []byte) (*RefreshNodesResponse, error) {
// var response RefreshNodesResponse
// if err := json.Unmarshal(jsondata, &response); err != nil {
// wrappedErr := fmt.Errorf("Error unmarshaling response: %w\n", err)
// return &response, wrappedErr
// }
// return &response, nil
// }

// RefreshNodes kicks off a node refresh with the given options
// it returns a jobid
func RefreshNodes(roptions RefreshNodesOptions) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), default_timeout_seconds*time.Second)
	defer cancel()

	client, collection, err := getJobsCollection()
	if err != nil {
		log.Println("got an error:", err)
		return "", err
	}
	defer client.Disconnect(ctx)

	jobid := fmt.Sprintf("%f-%f-%f-%f-%s",
		roptions.MinLon,
		roptions.MinLat,
		roptions.MaxLon,
		roptions.MaxLat, time.Now().Format("2006-01-02T15:04"),
	)

	job := &RefreshNodesJob{
		JobId:  jobid,
		MinLon: roptions.MinLon,
		MinLat: roptions.MinLat,
		MaxLon: roptions.MaxLon,
		MaxLat: roptions.MaxLat,
		Status: "queued",
	}
	fmt.Printf("Inserting document %v\n", job)
	result, err := collection.InsertOne(ctx, job)
	if err != nil {
		fmt.Printf("Error inserting job: %v\n", err)
		return "", fmt.Errorf("error inserting job: %w", err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return jobid, nil
}

func GetJob(jobid string) (*RefreshNodesJob, error) {
	ctx, cancel := context.WithTimeout(context.Background(), default_timeout_seconds*time.Second)
	defer cancel()

	client, collection, err := getJobsCollection()
	if err != nil {
		log.Println("got an error:", err)
		return nil, err
	}
	defer client.Disconnect(ctx)

	filter := bson.D{{Key: "job_id", Value: jobid}}

	var job RefreshNodesJob
	err = collection.FindOne(ctx, filter, options.FindOne()).Decode(&job)
	if err != nil {
		log.Println("got an error:", err)
		return nil, err
	}
	return &job, nil
}
