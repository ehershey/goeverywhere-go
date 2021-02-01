/*
if ignored.lower() == 'true':
    query['$and'].append({"ignored": True})
elif ignored.lower() == 'false':
    query['$and'].append({"ignored": {"$ne": True}})  # Field generally only present when == true

if priority.lower() == 'true':
    query['$and'].append({"priority": True})
elif priority.lower() == 'false':
    query['$and'].append({"priority": {"$ne": True}})  # Field generally only present when == true

if from_lat and from_lon:
    current_location = {"type": "Point", "coordinates": [float(from_lon), float(from_lat)]}
    near_query = {"loc": {"$near": current_location}}
    query['$and'].append(near_query)

if exclude:
    exclude_query = {"external_id": {"$nin": exclude}}
    query['$and'].append(exclude_query)

debug("query")

count = 0

points = []
last_included_entry_date = None
cursor = nodes.find(query).limit(limit)
for point in cursor:
    count = count + 1
    points.append(point)
response = {
  'min_lon': min_lon,
  'min_lat': min_lat,
  'max_lon': max_lon,
  'max_lat': max_lat,

  'rid': hashlib.md5(
      ("%s%s%s%s%s" %
          (bound_string, min_lon, max_lon, min_lat, max_lat)).encode(encoding='UTF-8',
                                                                     errors='strict')).hexdigest(),
  'bound_string': bound_string,
  'count': count,
  'limit': limit,
  'setsize': nodes.estimated_document_count(),
  'points': points
}


debug("executed")
response['count'] = count

print("Content-Type: text/plain")
print("")


print(json.dumps(response, default=json_util.default))
debug("response")
debug("dumped output")
debug("ending")
*/
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

const db_name = "ernie_org"
const collection_name = "nodes"
const db_uri = "mongodb://localhost:27017"

const default_limit = 1000

func getCollection() (*mongo.Client, *mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(db_uri))
	if err != nil {
		fmt.Println("got an error:", err)
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	collection := client.Database(db_name).Collection(collection_name)
	return client, collection, nil
}

type GetNodesOptions struct {
	min_lon      float64
	min_lat      float64
	max_lon      float64
	max_lat      float64
	from_lat     float64
	from_lon     float64
	ignored      bool
	priority     bool
	limit        int64
	exclude      string
	rind         string
	ts           int64
	bound_string string
}

func GetNodesHandler(http.ResponseWriter, *http.Request) {
	getNodes(GetNodesOptions{})
}
func getNodes(roptions GetNodesOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, collection, err := getCollection()
	if err != nil {
		fmt.Println("got an error:", err)
		return err
	}
	defer client.Disconnect(ctx)

	limit := roptions.limit

	if limit == 0 {
		limit = default_limit
	}

	min_lon := roptions.min_lon
	min_lat := roptions.min_lat
	max_lon := roptions.min_lon
	max_lat := roptions.max_lat
	if min_lon == 0 {
		min_lon = -80
	}
	if min_lat == 0 {
		min_lat = -80
	}
	if max_lon == 0 {
		max_lon = 80
	}
	if max_lat == 0 {
		max_lat = 80
	}

	var ands []bson.M

	box_query := bson.M{"loc": bson.M{"$geoIntersects": bson.M{"$geometry": bson.M{"type": "Polygon",
		"coordinates": bson.A{bson.A{bson.A{min_lon,
			min_lat},
			bson.A{min_lon,
				max_lat},
			bson.A{max_lon,
				max_lat},
			bson.A{max_lon,
				min_lat},
			bson.A{min_lon,
				min_lat}}},
	},
	},
	},
	}

	fmt.Println("ands:", ands)
	ands = append(ands, box_query)

	// query := bson.E{Key: "$and", Value: ands}
	query := bson.M{"$and": ands}

	//exclude_query = {"external_id": {"$nin": exclude}}
	//query['$and'].append(exclude_query)

	fmt.Println("ands:", ands)
	fmt.Println("box_query:", box_query)
	fmt.Println("query:", query)
	// query = bson.M{"$and": []bson.M{bson.M{"storyID": 123}, bson.M{"parentID": 432}}}
	find_opts := options.Find()
	find_opts.SetLimit(limit)

	cursor, err := collection.Find(ctx, query, find_opts)
	if err != nil {
		fmt.Println("got an error:", err)
		return err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var node bson.M
		if err = cursor.Decode(&node); err != nil {
			fmt.Println("got an error:", err)
			return err
		}
		fmt.Println(node)
	}

	return nil
}
