package client

import (
	"encoding/json"
	"log"
)

func toBytes(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("natsio/client.toBytes: marshal_json %v", err)
	}
	return b
}
