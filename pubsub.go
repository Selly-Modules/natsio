package natsio

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// Publish ...
func Publish(stream, subject string, payload []byte) error {
	channel := combineStreamAndSubjectName(stream, subject)

	_, err := natsJS.PublishAsync(channel, payload)
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
