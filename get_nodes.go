/*



# strip out empties
#
exclude = [int(i) for i in exclude.split('|') if i]
# also works: list(filter(lambda x: len(x) > 0, exclude.split('|')))


# bounds of map displayed when tile info was requested
#
bound_string = form.getfirst('bound_string', '')

debug("read form data")

# sql = "SELECT * FROM gps_log WHERE longitude > %s
# AND longitude < %s AND latitude > %s AND latitude < %s LIMIT 1" \
# % ( min_lon, max_lon, min_lat, max_lat)

query = {'$and': []}

box_query = {"loc":
             {"$geoIntersects":
              {"$geometry":
               {"type": "Polygon",
                "coordinates":
                [[[float(min_lon),
                    float(min_lat)],
                    [float(min_lon),
                     float(max_lat)],
                    [float(max_lon),
                     float(max_lat)],
                    [float(max_lon),
                     float(min_lat)],
                    [float(min_lon),
                     float(min_lat)]]]
                }
               }
              }
             }

query['$and'].append(box_query)

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
	"time"
)

const db_name = "ernie_org"
const collection_name = "nodes"
const db_uri = "mongodb://localhost:27017"

const default_limit = 1000

func getCollection() (*mongo.Client, *mongo.Collection, error) {
	fmt.Println("vim-go")
	client, err := mongo.NewClient(options.Client().ApplyURI(db_uri))
	if err != nil {
		fmt.Println("got an error:", err)
		return nil, nil, err
	}
	fmt.Println("got client:", client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("connecting")
	err = client.Connect(ctx)
	collection := client.Database(db_name).Collection(collection_name)
	return client, collection, nil
}

type GetNodesOptions struct {
	min_lon      int64
	min_lat      int64
	max_lon      int64
	max_lat      int64
	from_lat     string
	from_lon     string
	ignored      string
	priority     string
	limit        int64
	exclude      string
	rind         string
	ts           string
	bound_string string
}

func main() {
	options := GetNodesOptions{
		ignored:      "false",
		priority:     "true",
		exclude:      "294876208|4245240|294876209|294876210",
		limit:        4,
		from_lat:     "40.5900973",
		from_lon:     "-73.997701",
		bound_string: "%28%2840.58934490420493%2C%20-74.00047944472679%29%2C%20%2840.591811709253925%2C%20-73.99345205645294%29%29",
		rind:         "1/1",
		ts:           "1612114799249",
	}
	err := get_nodes(options)
	if err != nil {
		fmt.Println("got an error:", err)
	}
}
func get_nodes(query GetNodesOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, collection, err := getCollection()
	if err != nil {
		fmt.Println("got an error:", err)
		return err
	}
	defer client.Disconnect(ctx)

	fmt.Println("query.limit:", query.limit)
	limit := query.limit

	if limit == 0 {
		limit = default_limit
	}

	min_lon := query.min_lon
	min_lat := query.min_lat
	max_lon := query.min_lon
	max_lat := query.max_lat
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

	//exclude_query = {"external_id": {"$nin": exclude}}
	//query['$and'].append(exclude_query)

	// fmt.Println("collection:", collection)
	//filter := bson.D{{"collection", collection}}
	filter := bson.D{{}}
	find_opts := options.Find()
	find_opts.SetLimit(limit)
	find_opts.SetSort(bson.D{{"duration", -1}})

	cursor, err := collection.Find(ctx, filter, find_opts)
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
