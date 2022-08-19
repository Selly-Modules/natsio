package natsio

import (
	"log"

	"github.com/nats-io/nats.go"
)

// JSONEncoder ...
type JSONEncoder struct {
	encConn *nats.EncodedConn
}

// Subscribe ...
func (e JSONEncoder) Subscribe(subject string, cb nats.Handler) (*nats.Subscription, error) {
	sub, err := e.encConn.Subscribe(subject, cb)
	if err != nil {
		log.Printf("natsio.JSONEncoder.Subscribe err: %v\n", err)
	}
	return sub, err
}

// Publish ...
func (e JSONEncoder) Publish(reply string, data interface{}) error {
	return e.encConn.Publish(reply, data)
}
