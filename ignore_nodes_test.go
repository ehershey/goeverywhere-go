package main

import (
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
