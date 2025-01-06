package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const job_check_interval_seconds = 5
const job_max_run_time_seconds = 300 // includes db ops

// const job_max_retries = 5

// run as goroutine
func HandleJobs() {

	for {
		fmt.Printf("Checking for a queued job to run\n")
		err := HandleOneJob()
		if err != nil {
			wrappedErr := fmt.Errorf("Error trying to run one job: %w", err)
			log.Println("got an error:", wrappedErr)
		}
		fmt.Printf("Waiting before checking again\n")
		time.Sleep(job_check_interval_seconds * time.Second)
	}
}

func HandleOneJob() error {
	home, err := os.UserHomeDir()
	if err != nil {
		wrappedErr := fmt.Errorf("Error getting my homedir: %w", err)
		log.Println("got an error:", wrappedErr)
		return wrappedErr
	}
	command_argv0 := fmt.Sprintf("%s/simple_save_bound_nodes.sh", home)
	ctx, cancel := context.WithTimeout(context.Background(), job_max_run_time_seconds*time.Second)
	defer cancel()
	client, collection, err := getJobsCollection()
	if err != nil {
		wrappedErr := fmt.Errorf("Error getting Jobs Collection: %w", err)
		log.Println("got an error:", wrappedErr)
		return wrappedErr
	}
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Printf("Error disconnecting from db: %v\n", err)
		}
	}()

	var job RefreshNodesJob

	filter := bson.M{"job_status": "queued"}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "job_status", Value: "started"}, {Key: "status_time", Value: time.Now()}}}}
	update_opts := options.FindOneAndUpdate()
	update_opts.SetReturnDocument(options.After)
	err = collection.FindOneAndUpdate(ctx, filter, update, update_opts).Decode(&job)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			log.Println("No queued jobs")
			return nil
		}
		wrappedErr := fmt.Errorf("Error finding queued job in db: %w", err)
		log.Println("got an error:", wrappedErr)
		return wrappedErr
	}
	log.Printf("about to execute job: %v\n", job)
	cmd := exec.Command(command_argv0,
		fmt.Sprintf("min_lat=%f", job.MinLat),
		fmt.Sprintf("max_lat=%f", job.MaxLat),
		fmt.Sprintf("min_lon=%f", job.MinLon),
		fmt.Sprintf("max_lon=%f", job.MaxLon))
	out, err := cmd.CombinedOutput()
	if err != nil {
		if ntfyErr := errorNotify(err); ntfyErr != nil {
			log.Println("got ntfy error:", ntfyErr)
		}

		wrappedErr := fmt.Errorf("Error executing job command: %w", err)
		log.Println("got an error:", wrappedErr)
		job.Output = fmt.Sprintf("Error: %v\nOutput: %s", wrappedErr, out)
		failedErr := job.SetStatus("failed")
		if failedErr != nil {
			wrappedErr2 := fmt.Errorf("Error setting job status: %w", failedErr)
			log.Println("got an error:", wrappedErr2)
			return errors.Join(wrappedErr, wrappedErr2) // just abort if we get a second error while handling the command error
		}

		filter2 := bson.D{{Key: "_id", Value: job.Id}, {Key: "job_status", Value: "started"}}
		result, err := collection.ReplaceOne(ctx, filter2, job)
		if err != nil {
			wrappedErr2 := fmt.Errorf("Error saving job error to DB: %w", err)
			log.Println("got an error:", wrappedErr2)
			return errors.Join(wrappedErr, wrappedErr2) // just abort if we get a second error while handling the command error
		}
		if result.MatchedCount != 1 {
			wrappedErr2 := fmt.Errorf("bad Matched count after saving job error: %d", result.MatchedCount)
			log.Println("got an error:", wrappedErr2)
			return errors.Join(wrappedErr, wrappedErr2) // just abort if we get a second error while handling the command error
		}
		if result.ModifiedCount != 1 {
			wrappedErr2 := fmt.Errorf("bad modified count after saving job error: %d", result.ModifiedCount)
			log.Println("got an error:", wrappedErr2)
			return errors.Join(wrappedErr, wrappedErr2) // just abort if we get a second error while handling the command error
		}
		return wrappedErr
	}
	job.Output = string(out)
	err = job.SetStatus("completed")
	if err != nil {
		wrappedErr := fmt.Errorf("Error setting job status: %w", err)
		log.Println("got an error:", wrappedErr)
		return wrappedErr
	}

	filter3 := bson.D{{Key: "_id", Value: job.Id}, {Key: "job_status", Value: "started"}}
	result, err := collection.ReplaceOne(ctx, filter3, job)
	if err != nil {
		wrappedErr := fmt.Errorf("Error saving job output to DB: %w", err)
		log.Println("got an error:", wrappedErr)
		return wrappedErr
	}
	if result.MatchedCount != 1 {
		wrappedErr := fmt.Errorf("bad Matched count after saving job error: %d", result.MatchedCount)
		log.Println("got an error:", wrappedErr)
		return wrappedErr
	}
	if result.ModifiedCount != 1 {
		wrappedErr := fmt.Errorf("bad modified count after saving job error: %d", result.ModifiedCount)
		log.Println("got an error:", wrappedErr)
		return wrappedErr
	}

	return nil
}
