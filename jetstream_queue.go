package natsio

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// QueueSubscribe ...
func (js JetStream) QueueSubscribe(stream, subject, queueName string, cb nats.MsgHandler) error {
	channel := combineStreamAndSubjectName(stream, subject)

	_, err := js.instance.QueueSubscribe(channel, queueName, cb)
	if err != nil {
		msg := fmt.Sprintf("[NATS JETSTREAM] - queue subscribe with subject %s error: %s", channel, err.Error())
		return errors.New(msg)
	}
	return nil
}
