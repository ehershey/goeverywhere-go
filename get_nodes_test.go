package main

import (
	"fmt"
	"testing"
)

func TestGetNodes(t *testing.T) {
	options := GetNodesOptions{
		ignored:      false,
		priority:     true,
		exclude:      "294876208|4245240|294876209|294876210",
		limit:        4,
		from_lat:     40.5900973,
		from_lon:     -73.997701,
		bound_string: "%28%2840.58934490420493%2C%20-74.00047944472679%29%2C%20%2840.591811709253925%2C%20-73.99345205645294%29%29",
		rind:         "1/1",
		ts:           1612114799249,
	}
	err := getNodes(options)
	if err != nil {
		fmt.Println("got an error:", err)
	}
}
