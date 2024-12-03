package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
)

func TestIgnoreNode(t *testing.T) {
	roptions := IgnoreNodesOptions{NodeId: 243030850, Action: "ignore"}
	node, err := ignoreNodes(roptions)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	if node.CityName != "Oaxaca de Juárez" {
		t.Errorf("Node city name = %v; want Oaxaca de Juárez", node.CityName)
	}
	if node.Ignored != true {
		t.Errorf("Node ignored = %v; want true", node.Ignored)
	}
}

func TestUnignoreNode(t *testing.T) {
	roptions := IgnoreNodesOptions{NodeId: 243030850, Action: "unignore"}
	node, err := ignoreNodes(roptions)
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
	if node.CityName != "Oaxaca de Juárez" {
		t.Errorf("Node city name = %v; want Oaxaca de Juárez", node.CityName)
	}
	if node.Ignored != false {
		t.Errorf("Node ignored = %v; want false", node.Ignored)
	}
}

func TestIgnoreNodesHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:1234/ignore_nodes?node_id=243030850&action=ignore", nil)
	w := httptest.NewRecorder()
	IgnoreNodesHandler(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	// log.Println("body: ", string(body))

	if resp.StatusCode != 200 {
		t.Errorf("resp.StatusCode = %d; want 200", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("resp.Header.Get(\"Content-Type\") = %v; want \"application/json\"", resp.Header.Get("Content-Type"))
	}

	err, response := DecodeIgnoreResponse(body)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	node := response

	if node.CityName != "Oaxaca de Juárez" {
		t.Errorf("Node city name = %v; want Oaxaca de Juárez", node.CityName)
	}
	if node.Ignored != true {
		t.Errorf("Node ignored = %v; want true", node.Ignored)
	}
}

func TestIgnoreNodesHandlerUnignore(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:1234/ignore_nodes?node_id=243030850&action=unignore", nil)
	w := httptest.NewRecorder()
	IgnoreNodesHandler(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	// log.Println("body: ", string(body))

	if resp.StatusCode != 200 {
		t.Errorf("resp.StatusCode = %d; want 200", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("resp.Header.Get(\"Content-Type\") = %v; want \"application/json\"", resp.Header.Get("Content-Type"))
	}

	err, response := DecodeIgnoreResponse(body)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	node := response

	if node.CityName != "Oaxaca de Juárez" {
		t.Errorf("Node city name = %v; want Oaxaca de Juárez", node.CityName)
	}
	if node.Ignored != false {
		t.Errorf("Node ignored = %v; want false", node.Ignored)
	}
}

func DecodeIgnoreResponse(jsondata []byte) (error, *IgnoreNodesResponse) {
	var response IgnoreNodesResponse
	if err := json.Unmarshal(jsondata, &response); err != nil {
		wrappedErr := fmt.Errorf("Error unmarshaling response: %w", err)
		return wrappedErr, &response
	}
	return nil, &response
}
