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
// js := natsio.GetJetStream()
//
// sub, err := js.PullSubscribe("A_SUBJECT", "A_DURABLE")
//
// for {
//   messages, err := sub.Fetch(10)
//   // process each messages
// }
//
func (js JetStream) PullSubscribe(stream, subject, durable, consumer string) (*nats.Subscription, error) {
	channel := combineStreamAndSubjectName(stream, subject)

	// Check if consumer existed
	con, err := js.instance.ConsumerInfo(stream, consumer)
	fmt.Println("con", con)
	fmt.Println("err", err)
	if con == nil {
		info, err := js.instance.AddConsumer(stream, &nats.ConsumerConfig{
			Durable:       durable,
			AckPolicy:     nats.AckExplicitPolicy,
			FilterSubject: subject,
		})
		if err == nil {
			fmt.Println("CONSUMER INFO", info)
		} else {
			fmt.Println("ADD CONSUMER ERROR", err)
		}
	}

	sub, err := js.instance.PullSubscribe(subject, durable)
	if err != nil {
		msg := fmt.Sprintf("[NATS JETSTREAM] - pull subscribe subject #%s - durable #%s error: %s", channel, durable, err.Error())
		return nil, errors.New(msg)
	}

	return sub, nil
}
