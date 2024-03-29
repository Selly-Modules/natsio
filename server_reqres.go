package natsio

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// Default timeout 10s
const requestTimeout = 10 * time.Second

// Request ...
func (sv Server) Request(subject string, payload []byte) (*nats.Msg, error) {
	timeout := requestTimeout
	if sv.Config.RequestTimeout > 0 {
		timeout = sv.Config.RequestTimeout
	}
	msg, err := sv.instance.Request(subject, payload, timeout)
	if errors.Is(err, nats.ErrNoResponders) {
		log.Printf("[NATS SERVER]: request - no responders for subject: %s", subject)
	}
	return msg, err
}

// Reply ...
func (sv Server) Reply(msg *nats.Msg, payload []byte) error {
	return sv.instance.Publish(msg.Reply, payload)
}

// Subscribe ...
func (sv Server) Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	sub, err := sv.instance.Subscribe(subject, cb)
	if err != nil {
		msg := fmt.Sprintf("[NATS SERVER] - subscribe subject %s error: %s", subject, err.Error())
		return nil, errors.New(msg)
	}
	return sub, nil
}

// NewJSONEncodedConn ...
func (sv Server) NewJSONEncodedConn() (*JSONEncoder, error) {
	enc, err := nats.NewEncodedConn(sv.instance, nats.JSON_ENCODER)
	if err != nil {
		log.Printf("natsio.NewJSONEncodedConn: err %v\n", err)
		return nil, err
	}
	return &JSONEncoder{encConn: enc}, nil
}
