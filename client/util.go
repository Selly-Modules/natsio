package client

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func toBytes(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("natsio/client.toBytes: marshal_json %v", err)
	}
	return b
}

// bsonToBytes ...
func bsonToBytes(data interface{}) []byte {
	b, _ := bson.Marshal(data)
	return b
}
