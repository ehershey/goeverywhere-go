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

func getCollectionByName(collection_name string) (*mongo.Client, *mongo.Collection, error) {
	reg := bson.NewRegistry()

	reg.RegisterTypeDecoder(reflect.TypeOf(&timestamppb.Timestamp{}), bson.ValueDecoderFunc(timeDecoder))
	reg.RegisterTypeDecoder(reflect.TypeOf(timestamppb.Timestamp{}), bson.ValueDecoderFunc(timeDecoder))

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
	timeType := reflect.TypeOf(&timestamppb.Timestamp{})
	if !val.IsValid() || !val.CanSet() || val.Type() != timeType {
		return bson.ValueDecoderError{
			Name:     "timeDecoder",
			Types:    []reflect.Type{timeType},
			Received: val,
		}
	}

	var result time.Time
	switch vr.Type() {
	case bson.TypeDateTime:
		t, err := vr.ReadDateTime()
		//t, _, err := vr.ReadTimestamp()
		if err != nil {
			wrappedErr := fmt.Errorf("Error reading bson datetime: %w", err)
			return wrappedErr
		}
		result = time.Unix(t/1000, 0)
	case bson.TypeTimestamp:
		t, _, err := vr.ReadTimestamp()
		if err != nil {
			wrappedErr := fmt.Errorf("Error reading bson timestamp: %w", err)
			return wrappedErr
		}
		result = time.Unix(int64(t), 0)
	default:
		return fmt.Errorf(
			"received invalid BSON type to decode into time: %s",
			vr.Type())
	}

	tspbval := timestamppb.New(result)
	tspbrefval := reflect.ValueOf(tspbval)
	val.Set(tspbrefval)
	return nil
}
