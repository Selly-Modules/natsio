package natsio

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// QueueSubscribe ...
func QueueSubscribe(subject, queueName string, cb nats.MsgHandler) error {
	_, err := natsJS.QueueSubscribe(subject, queueName, cb)
	if err != nil {
		msg := fmt.Sprintf("queue subscribe with subject %s error: %s", subject, err.Error())
		return errors.New(msg)
	}
	return nil
}
