package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
	servertiming "github.com/mitchellh/go-server-timing"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const nodes_collection_name = "nodes"

const default_limit = 1000

const default_timeout_seconds = 60

func getNodesCollection() (*mongo.Client, *mongo.Collection, error) {
	return getCollectionByName(nodes_collection_name)
}

type GetNodesOptions struct {
	MinLon          float64 `schema:"min_lon"`
	MinLat          float64 `schema:"min_lat"`
	MaxLon          float64 `schema:"max_lon"`
	MaxLat          float64 `schema:"max_lat"`
	FromLat         float64 `schema:"from_lat"`
	FromLon         float64 `schema:"from_lon"`
	AllowIgnored    bool    `schema:"allow_ignored"`
	RequirePriority bool    `schema:"require_priority"`
	MaxDistance     float64 `schema:"max_distance"`
	Limit           int
	Exclude         string
	Ts              string
	BoundString     string `schema:"bound_string"`
	Rind            string
	NodeId          int `schema:"node_id"`
}

type Point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Node struct {
	Id              bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	CreationDate    time.Time     `bson:"creation_date" json:"creation_date"`
	Loc             Point         `bson:"loc" json:"loc"`
	ExternalId      int           `bson:"external_id,omitempty" json:"external_id"`
	Priority        bool          `bson:"priority" json:"priority"`
	Ignored         bool          `bson:"ignored" json:"ignored"`
	StreetName      string        `bson:"street_name" json:"street_name"`
	StreetNodeCount int           `bson:"street_node_count" json:"street_node_count"`
	CityName        string        `bson:"city_name" json:"city_name"`
	StreetID        int           `bson:"street_id" json:"street_id"`
	CityID          int           `bson:"city_id" json:"city_id"`
	Deactivated     bool          `bson:"deactivated" json:"deactivated"`
}

func (n *Node) GetLat() float64 {
	return n.Loc.GetLat()
}

func (p *Point) GetLat() float64 {
	return p.Coordinates[1]
}

func (n *Node) GetLon() float64 {
	return n.Loc.GetLon()
}

func (p *Point) GetLon() float64 {
	return p.Coordinates[0]
}

type getNodesResponse struct {
	MinLon      float64 `json:"min_lon"`
	MinLat      float64 `json:"min_lat"`
	MaxLon      float64 `json:"max_lon"`
	MaxLat      float64 `json:"max_lat"`
	FromLon     float64 `json:"from_lon"`
	FromLat     float64 `json:"from_lat"`
	Rid         string  `json:"rid"`
	BoundString string  `json:"bound_string"`
	Count       int     `json:"count"`
	Limit       int     `json:"limit"`
	MaxDistance float64 `json:"max_distance"`
	Setsize     int     `json:"setsize"`
	Points      []Node  `json:"points"`
}

var decoder = schema.NewDecoder()

// GetNodesHandlerWithTiming wraps our handler with
// the server timing middleware
var GetNodesHandlerWithTiming = servertiming.Middleware(http.HandlerFunc(GetNodesHandler), nil)

// GetNodesHandler returns nodes based on an HTTP request
// without server timing headers
func GetNodesHandler(w http.ResponseWriter, r *http.Request) {

	timing := servertiming.FromContext(r.Context())

	metric := timing.NewMetric("translate input for query").Start()

	err := r.ParseForm()
	var querymap map[string][]string

	if r.Method == "GET" {
		querymap = r.URL.Query()
	} else {
		querymap = r.PostForm
	}

	if err != nil {
		log.Println("got an error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var roptions GetNodesOptions

	err = decoder.Decode(&roptions, querymap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	metric.Stop()

	// for excludes:
	// https://stackoverflow.com/a/37533144/408885
	// func arrayToString(a []int, delim string) string {
	// return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	// //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	// //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
	// }
	nodes, err := getNodes(roptions.sanitize())
	if err != nil {
		log.Println("got an error calling getNodes(roptions.sanitize()):", err)
		return
	}

	h := md5.New()
	if _, err := io.WriteString(h, fmt.Sprintf("%s%f%f%f%f%f%f%f%d", roptions.BoundString, roptions.MinLon, roptions.MaxLon, roptions.MinLat, roptions.MaxLat, roptions.FromLat, roptions.FromLon, roptions.MaxDistance, roptions.NodeId)); err != nil {
		fmt.Printf("Error writing hash string: %v\n", err)
	}
	requestHash := fmt.Sprintf("%x", h.Sum(nil))

	totalcount, err := getTotalCount()
	if err != nil {
		log.Println("got an error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := &getNodesResponse{
		MinLon:      roptions.MinLon,
		MinLat:      roptions.MinLat,
		MaxLon:      roptions.MaxLon,
		MaxLat:      roptions.MaxLat,
		FromLon:     roptions.FromLon,
		FromLat:     roptions.FromLat,
		MaxDistance: roptions.MaxDistance,
		Rid:         requestHash,
		Points:      nodes,
		BoundString: roptions.BoundString,
		Count:       len(nodes),
		Limit:       roptions.Limit,
		Setsize:     int(totalcount)}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}

}

func DecodeResponse(jsondata []byte) (error, *getNodesResponse) {
	var response getNodesResponse
	if err := json.Unmarshal(jsondata, &response); err != nil {
		wrappedErr := fmt.Errorf("Error unmarshaling json data: %w\n", err)
		return wrappedErr, &response
	}

	return nil, &response
}

func (roptions *GetNodesOptions) sanitize() GetNodesOptions {
	if roptions.Limit == 0 {
		roptions.Limit = default_limit
	}
	return *roptions
}

func getTotalCount() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), default_timeout_seconds*time.Second)
	defer cancel()

	client, collection, err := getNodesCollection()
	if err != nil {
		log.Println("got an error:", err)
		return 0, err
	}

	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Printf("Error disconnecting from db: %v\n", err)
		}
	}()

	return collection.EstimatedDocumentCount(ctx)
}

func getNodes(roptions GetNodesOptions) ([]Node, error) {
	log.Printf("start getNodes(%v)\n", roptions)
	defer func() { log.Printf("end getNodes(%v)\n", roptions) }()
	ctx, cancel := context.WithTimeout(context.Background(), default_timeout_seconds*time.Second)
	defer cancel()

	client, collection, err := getNodesCollection()
	if err != nil {
		log.Println("got an error:", err)
		return nil, err
	}
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Printf("Error disconnecting from db: %v\n", err)
		}
	}()

	if roptions.MinLon > roptions.MaxLon {
		return nil, errors.New("min_lon must be <= max_lon")
	}
	if roptions.MinLat > roptions.MaxLat {
		return nil, errors.New("min_lat must be <= max_lat")
	}

	var ands []bson.M

	if roptions.NodeId != 0 {
		ands = append(ands, bson.M{"external_id": roptions.NodeId})
	}

	if roptions.FromLat != 0 && roptions.FromLon != 0 && roptions.FromLat != -1 && roptions.FromLon != -1 {
		coords := make([]float64, 2)
		coords[0] = roptions.FromLon
		coords[1] = roptions.FromLat
		current_location := Point{Type: "Point", Coordinates: coords}
		// var loc_doc []bson.M
		var loc_doc bson.D
		loc_doc = append(loc_doc, bson.E{Key: "$near", Value: current_location})

		if roptions.MaxDistance > 0 {
			// near_query["loc"].(map[string]interface{})["$maxDistance"] = roptions.MaxDistance
			//near_query["loc"].(bson.D)["$maxDistance"] = roptions.MaxDistance
			loc_doc = append(loc_doc, bson.E{Key: "$maxDistance", Value: roptions.MaxDistance})
		}
		near_query := bson.M{"loc": loc_doc}

		ands = append(ands, near_query)
	}

	if roptions.MinLon != 0 || roptions.MinLat != 0 || roptions.MaxLon != 0 || roptions.MaxLat != 0 {
		box_query := bson.M{"loc": bson.M{"$geoIntersects": bson.M{"$geometry": bson.M{"type": "Polygon",
			"coordinates": bson.A{bson.A{bson.A{roptions.MinLon,
				roptions.MinLat},
				bson.A{roptions.MinLon,
					roptions.MaxLat},
				bson.A{roptions.MaxLon,
					roptions.MaxLat},
				bson.A{roptions.MaxLon,
					roptions.MinLat},
				bson.A{roptions.MinLon,
					roptions.MinLat}}},
		},
		},
		},
		}

		ands = append(ands, box_query)
	}

	// Fields generally only present in db when == true
	// defaults are to return only priority nodes that aren't ignored
	if !roptions.AllowIgnored {
		ands = append(ands, bson.M{"ignored": bson.M{"$ne": true}})
	}

	// Can make this more interesting later.. for now only even acknowledge
	// deactivated nodes exist if searching for a specific one by ID
	//
	if roptions.NodeId == 0 {
		ands = append(ands, bson.M{"deactivated": bson.M{"$ne": true}})
	}

	if roptions.RequirePriority {
		ands = append(ands, bson.M{"priority": true})
	}

	if len(roptions.Exclude) > 0 {
		exclude_array := make([]int64, 0)
		for _, exclude_id := range strings.Split(roptions.Exclude, "|") {
			var exclude_id_int int64
			if exclude_id_int, err = strconv.ParseInt(exclude_id, 10, 64); err != nil {
				return nil, fmt.Errorf("Error parsing exclude id into int64: %w", err)
			}
			exclude_array = append(exclude_array, exclude_id_int)
		}

		exclude_query := bson.M{"external_id": bson.M{"$nin": exclude_array}}
		ands = append(ands, exclude_query)
	}

	// query := bson.E{Key: "$and", Value: ands}
	query := bson.M{}
	if len(ands) > 0 {
		query["$and"] = ands
	}

	// log.Println("ands:", ands)
	// log.Println("query:", query)

	// query = bson.M{"$and": []bson.M{bson.M{"storyID": 123}, bson.M{"parentID": 432}}}
	find_opts := options.Find()
	find_opts.SetLimit(int64(roptions.Limit))

	var nodes []Node
	log.Printf("roptions.Limit: %v\n", roptions.Limit)
	cursor, err := collection.Find(ctx, query, find_opts)
	if err != nil {
		log.Println("got an error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &nodes); err != nil {
		return nil, err
	}
	log.Printf("Returning %v nodes in getNodes(%v)\n", len(nodes), roptions)
	return nodes, nil
}
