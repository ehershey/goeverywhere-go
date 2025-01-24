package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const STRAVA_COLLECTION_NAME = "activities"

var timevar time.Time

func getCollectionByName(collection_name string) (*mongo.Client, *mongo.Collection, error) {
	reg := bson.NewRegistry()

	//codecOpt := bsonoptions.StringCodec().SetDecodeObjectIDAsHex(true)
	//strCodec := bsoncodec.NewStringCodec(codecOpt)

	//reg := bson.NewRegistryBuilder().RegisterDefaultDecoder(reflect.String, strCodec).Build()

	//dc := bsoncodec.DecodeContext{Registry: reg}

	//decoder := bson.NewDecoder(bson.NewDocumentReader(bytes.NewReader([]byte{})))
	//decoder.ObjectIDAsHexString()

	// reg := codecs.Register(bson.NewRegistryBuilder()).Build()
	//var oidvar bson.ObjectID
	//var oidtype = reflect.TypeOf(oidvar)
	//reg.RegisterTypeDecoder(oidtype, decoder)
	//reg.RegisterTypeDecoder(reflect.TypeOf(timevar),
	//bson.ValueDecoderFunc(timeDecoder))

	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	uri := config.MongoDB_Uri
	db_name := config.DB_Name
	if collection_name == STRAVA_COLLECTION_NAME {
		uri = config.Strava_MongoDB_Uri
		db_name = config.Strava_DB_Name
	}

	bson_options := options.BSONOptions{ObjectIDAsHexString: true}
	client, err := mongo.Connect(options.Client().ApplyURI(uri).SetRegistry(reg).SetBSONOptions(&bson_options))
	if err != nil {
		wrappedErr := fmt.Errorf("Error creating mongodb client: %w", err)
		return nil, nil, wrappedErr
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	// if err := client.Connect(ctx); err != nil {
	// wrappedErr := fmt.Errorf("Error connecting to mongodb: %w", err)
	// return nil, nil, wrappedErr
	// }
	collection := client.Database(db_name).Collection(collection_name)
	return client, collection, nil
}

func timeDecoder(
	_ bson.DecodeContext,
	vr bson.ValueReader,
	val reflect.Value,
) error {
	// All decoder implementations should check that val is valid, settable,
	// and is of the correct kind before proceeding.
	timeType := reflect.TypeOf(timevar)
	if !val.IsValid() || !val.CanSet() || val.Type() != timeType {
		return bson.ValueDecoderError{
			Name:     "timeDecoder",
			Types:    []reflect.Type{timeType},
			Received: val,
		}
	}

	var result timestamppb.Timestamp
	switch vr.Type() {
	case bson.TypeTimestamp:
		t, _, err := vr.ReadTimestamp()
		if err != nil {
			return err
		}
		ts := &timestamppb.Timestamp{Seconds: int64(t)}
		result = *ts
	default:
		return fmt.Errorf(
			"received invalid BSON type to decode into time: %s",
			vr.Type())
	}

	val.Set(reflect.ValueOf(result))
	return nil
}
