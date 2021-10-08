package natsio

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// Publish ...
func Publish(subject string, data interface{}) error {
	b, _ := json.Marshal(data)
	_, err := natsJS.PublishAsync(subject, b)
	if err != nil {
		msg := fmt.Sprintf("publish message error: %s", err.Error())
		return errors.New(msg)
	}
	return nil
}

// Subscribe ...
func Subscribe(subject string, cb nats.MsgHandler) error {
	_, err := natsJS.Subscribe(subject, cb)
	if err != nil {
		msg := fmt.Sprintf("subscribe subject %s error: %s", subject, err.Error())
		return errors.New(msg)
	}
	return nil
}
