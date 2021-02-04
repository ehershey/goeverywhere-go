package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetNodesIncludeIgnored(t *testing.T) {
	roptions := GetNodesOptions{AllowIgnored: true}
	nodes, err := getNodes(roptions.sanitize())
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	seenIgnored := false
	for _, node := range nodes {
		if node.Ignored == true {
			seenIgnored = true
		}
	}
	if !seenIgnored {
		t.Errorf("No ignored nodes returned when AllowIignored: true in query")
	}

	roptions.AllowIgnored = false
	nodes, err = getNodes(roptions)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	for _, node := range nodes {
		if node.Ignored == true {
			t.Errorf("Ignored node returned when AllowIgnored == false")
		}
	}

}
func TestGetNodes(t *testing.T) {
	options := GetNodesOptions{
		AllowIgnored:    false,
		RequirePriority: true,
		Exclude:         "294876208|4245240|294876209|294876210",
		Limit:           4,
		FromLat:         40.5900973,
		FromLon:         -73.997701,
	}
	nodes, err := getNodes(options.sanitize())
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	if len(nodes) != 4 {
		t.Errorf("len(nodes) = %d; want 4", len(nodes))
	}
	options.Limit = 0
	nodes, err = getNodes(options)

	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	for index, node := range nodes {
		if node.Ignored == true {
			t.Errorf("nodes[%d].Ignored == true; want none true (external_id: %d)", index, node.ExternalId)
		}
		if node.Priority != true {
			t.Errorf("nodes[%d].Priority != true; want all true (external_id: %d)", index, node.ExternalId)
		}
	}

	options.MinLon = -73.9920391530954
	options.MinLat = 40.67437941796339
	options.MaxLon = -73.93582004690457
	options.MaxLat = 40.69511656898346

	nodes, err = getNodes(options)

	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	for index, node := range nodes {
		if node.GetLat() < options.MinLat {
			t.Errorf("nodes[%d].GetLat() < %f (min_lat) (external_id: %d)", index, node.GetLat(), node.ExternalId)
		}
		if node.GetLon() < options.MinLon {
			t.Errorf("nodes[%d].GetLon() < %f (min_lon) (external_id: %d)", index, node.GetLon(), node.ExternalId)
		}
		if node.GetLat() > options.MaxLat {
			t.Errorf("nodes[%d].GetLat() > %f (max_lat) (external_id: %d)", index, node.GetLat(), node.ExternalId)
		}
		if node.GetLon() > options.MaxLon {
			t.Errorf("nodes[%d].GetLon() > %f (max_lon) (external_id: %d)", index, node.GetLon(), node.ExternalId)
		}

	}
	// should error

	options.MinLon = 80
	options.MinLat = 0
	options.MaxLon = -80
	options.MaxLat = 0

	nodes, err = getNodes(options)

	if err == nil {
		t.Errorf("expected an error with invalid min_lon(%f)/max_lon(%f) but got %d nodes in result", options.MinLon, options.MaxLon, len(nodes))
	}

	options.MinLon = 0
	options.MinLat = 20
	options.MaxLon = 0
	options.MaxLat = 10

	nodes, err = getNodes(options)

	if err == nil {
		t.Errorf("expected an error with invalid min_lat(%f)/max_lat(%f) but got %d nodes in result", options.MinLat, options.MaxLat, len(nodes))
	}

	// test for exact result

	options.MinLon = 0
	options.MinLat = 0
	options.MaxLon = 0
	options.MaxLat = 0
	options.Limit = 4

	nodes, err = getNodes(options.sanitize())

	if err != nil {
		t.Errorf("got an error: %v", err)
	}

	if len(nodes) == 0 {
		t.Errorf("No nodes in response")
		return
	}
	if nodes[0].ExternalId != 4245239 {
		t.Errorf("nodes[0].ExternalId = %d; want 4245239", nodes[0].ExternalId)
	}

}

func TestGetNodesHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:1234/nodes?allow_ignored=false&require_priority=true&exclude=294876208|4245240|294876209|294876210&limit=4&from_lat=40.5900973&from_lon=-73.997701&bound_string=%28%2840.58934490420493%2C%20-74.00047944472679%29%2C%20%2840.591811709253925%2C%20-73.99345205645294%29%29&rind=1/1&ts=1612114799249", nil)
	w := httptest.NewRecorder()
	GetNodesHandler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("resp.StatusCode = %d; want 200", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("resp.Header.Get(\"Content-Type\") = %v; want \"application/json\"", resp.Header.Get("Content-Type"))
	}
	var responsejson map[string]interface{}
	json.Unmarshal([]byte(body), &responsejson)

	if responsejson["min_lon"] != -80.0 {
		t.Errorf("Response JSON doesn't contain 'min_lon' key set to -80: (%f) / %s", responsejson["min_lon"], body[0:80])
	}

	if responsejson["min_lat"] != -80.0 {
		t.Errorf("Response JSON doesn't contain 'min_lat' key set to -80: (%f) / %s", responsejson["min_lat"], body[0:80])
	}

	if responsejson["max_lon"] != 80.0 {
		t.Errorf("Response JSON doesn't contain 'max_lon' key set to 80: (%f) / %s", responsejson["max_lon"], body[0:80])
	}

	if responsejson["max_lat"] != 80.0 {
		t.Errorf("Response JSON doesn't contain 'max_lat' key set to 80: (%f) / %s", responsejson["max_lat"], body[0:80])
	}

	if responsejson["bound_string"] != "((40.58934490420493, -74.00047944472679), (40.591811709253925, -73.99345205645294))" {
		t.Errorf("Incorrect bound string in response: %s", responsejson["bound_string"])
	}

	point := (responsejson["points"].([]interface{})[0]).(map[string]interface{})

	log.Println("point: ", point)
	loc := point["loc"].(map[string]interface{})
	loc_type := loc["type"]
	if loc_type != "Point" {
		t.Errorf("loc type in json is not \"Point\": %s", loc_type)
	}
	coordinates := loc["coordinates"].([]interface{})
	lat := coordinates[0]
	lon := coordinates[1]
	if lat != -73.9962469 {
		t.Errorf("lat in node in json != -73.9962469: %f", lat)
	}
	if lon != 40.5908497 {
		t.Errorf("lon in node in json != 40.5908497: %f", lat)
	}

}

func TestGetNodesExclude(t *testing.T) {
	options := GetNodesOptions{}
	nodes, err := getNodes(options.sanitize())
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	excluded_ids := make([]int, 3)

	if len(nodes) == 0 {
		t.Errorf("No nodes in response")
		return
	}
	excluded_ids[0] = nodes[0].ExternalId
	options.Exclude = strings.Trim(strings.Replace(fmt.Sprint(excluded_ids), " ", "|", -1), "[]")

	nodes, err = getNodes(options)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	for index, node := range nodes {
		for _, excluded_id := range excluded_ids {
			if node.ExternalId == excluded_id {
				t.Errorf("nodes[%d].ExternalId in excluded list (%d)", index, excluded_id)
			}
		}
	}

	// add a second excluded

	excluded_ids[1] = nodes[0].ExternalId
	options.Exclude = strings.Trim(strings.Replace(fmt.Sprint(excluded_ids), " ", "|", -1), "[]")

	nodes, err = getNodes(options)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	for index, node := range nodes {
		for _, excluded_id := range excluded_ids {
			if node.ExternalId == excluded_id {
				t.Errorf("nodes[%d].ExternalId in excluded list (%d)", index, excluded_id)
			}
		}
	}

	// and a third

	excluded_ids[2] = nodes[0].ExternalId
	options.Exclude = strings.Trim(strings.Replace(fmt.Sprint(excluded_ids), " ", "|", -1), "[]")

	nodes, err = getNodes(options)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	for index, node := range nodes {
		for _, excluded_id := range excluded_ids {
			if node.ExternalId == excluded_id {
				t.Errorf("nodes[%d].ExternalId in excluded list (%d)", index, excluded_id)
			}
		}
	}
}
