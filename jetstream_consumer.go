package natsio

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// GetConsumerInfo ...
func (js JetStream) GetConsumerInfo(stream, name string) (*nats.ConsumerInfo, error) {
	return js.instance.ConsumerInfo(stream, name)
}

// AddConsumer ...
func (js JetStream) AddConsumer(stream, subject, name string) error {
	channel := combineStreamAndSubjectName(stream, subject)

	// Add
	_, err := js.instance.AddConsumer(stream, &nats.ConsumerConfig{
		Durable:       name,
		AckPolicy:     nats.AckExplicitPolicy,
		FilterSubject: channel,
	})

	if err != nil {
		msg := fmt.Sprintf("[NATS JETSTREAM] - add consumer %s for stream #%s error: %s", name, stream, err.Error())
		return errors.New(msg)
	}
	return nil
}
