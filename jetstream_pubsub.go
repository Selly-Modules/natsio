package natsio

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// Publish ...
func (js JetStream) Publish(stream, subject string, payload []byte) error {
	channel := combineStreamAndSubjectName(stream, subject)

	_, err := js.instance.PublishAsync(channel, payload)
	if err != nil {
		msg := fmt.Sprintf("[NATS JETSTREAM] - publish message to subject #%s error: %s", channel, err.Error())
		return errors.New(msg)
	}
	return nil
}

// Subscribe ...
func (js JetStream) Subscribe(stream, subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	channel := combineStreamAndSubjectName(stream, subject)

	sub, err := js.instance.Subscribe(channel, cb)
	if err != nil {
		msg := fmt.Sprintf("[NATS JETSTREAM] - subscribe subject %s error: %s", channel, err.Error())
		return nil, errors.New(msg)
	}
	return sub, nil
}

// PullSubscribe ...
//
// Example:
//
// sub, err := natsio.PullSubscribe("A_STREAM", "A_OBJECT")
//
// for {
//   messages, err := sub.Fetch(10)
//   // process each messages
// }
//
func (js JetStream) PullSubscribe(stream, subject string) (*nats.Subscription, error) {
	channel := combineStreamAndSubjectName(stream, subject)

	sub, err := js.instance.PullSubscribe(stream, subject)
	if err != nil {
		msg := fmt.Sprintf("[NATS JETSTREAM] - pull subscribe subject #%s error: %s", channel, err.Error())
		return nil, errors.New(msg)
	}

	return sub, nil
}
