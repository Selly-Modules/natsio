package natsio

import (
	"errors"
	"fmt"
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
	return sv.instance.Request(subject, payload, timeout)
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
