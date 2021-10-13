package natsio

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// Publish ...
func Publish(stream, subject string, data interface{}) error {
	channel := combineStreamAndSubjectName(stream, subject)

	b, _ := json.Marshal(data)
	_, err := natsJS.PublishAsync(channel, b)
	if err != nil {
		msg := fmt.Sprintf("publish message to subject %s error: %s", channel, err.Error())
		return errors.New(msg)
	}
	return nil
}

// Subscribe ...
func Subscribe(stream, subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	channel := combineStreamAndSubjectName(stream, subject)

	sub, err := natsJS.Subscribe(channel, cb)
	if err != nil {
		msg := fmt.Sprintf("subscribe subject %s error: %s", channel, err.Error())
		return nil, errors.New(msg)
	}
	return sub, nil
}
