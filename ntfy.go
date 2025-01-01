package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func errorNotify(jobErr error) error {
	log.Printf("errorNotify(%v) start", jobErr)
	defer log.Printf("errorNotify(%v) end", jobErr)
	config, err := GetConfig()
	if err != nil {
		wrappedErr := fmt.Errorf("Error getting config for error notifications: %w", err)
		log.Println("Got an error:", wrappedErr)
		return wrappedErr
	}

	topic := config.Error_Ntfy_Topic
	if topic == "" {
		log.Printf("No ntfy topic defined - not sending error notification\n")
		return nil
	}
	log.Printf("Sending error notification to ntfy topic: %v\n", topic)

	ntfy_url := fmt.Sprintf("https://ntfy.sh/%v", topic)

	bodyString := fmt.Sprintf("%v", jobErr)
	body := strings.NewReader(bodyString)
	title := "GOE Job Error"

	Debug := true
	Verbose := true
	DryRun := false
	req, _ := http.NewRequest("POST", ntfy_url, body)
	// req.Header.Set("Click", url)
	// req.Header.Set("Actions", fmt.Sprintf("view, View on Strava, %s", url))
	req.Header.Set("Title", title)

	log.Printf("ntfy_url: %v\n", ntfy_url)
	log.Printf("body: %v\n", body)

	if Debug {
		log.Printf("req:\n")
		spew.Dump(req)
	}
	if !DryRun {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			wrappedErr := fmt.Errorf("error contacting ntfy.sh: %w\n", err)
			log.Println("Got an error:", wrappedErr)
			return wrappedErr
		}
		log.Printf("res.StatusCode: %v\n", res.StatusCode)
		if res.StatusCode != 200 {
			// TODO clean this up / verify logic
			//
			err := fmt.Errorf("non-OK response from ntfy.sh: %v\nBody:\n", res.StatusCode)
			log.Println("Got an error:", err)
			bytesWritten, err := io.Copy(os.Stdout, res.Body)
			if err != nil {
				wrappedErr := fmt.Errorf("error copying http error to stdout: %w\n", err)
				log.Println("Got an error:", wrappedErr)
			}
			if bytesWritten < 20 {
				log.Printf("odd amount of data copied from http error to stdount (%d)\n", bytesWritten)
			}
			log.Printf("%v\n", res)
			return err
		}
	} else {
		if Verbose {
			fmt.Printf("Skipping ntfy server call in dry run mode\n")
		}
	}
	return nil
}
