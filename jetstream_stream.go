package natsio

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// GetStreamInfo ...
func (js JetStream) GetStreamInfo(name string) (*nats.StreamInfo, error) {
	return js.instance.StreamInfo(name)
}

// AddStream add new stream, with default config
// Due to subject must have a unique name, subject name will be combined with stream name
// E.g: stream name is "DEMO", subject name is "Subject-1", so final name in NATS will be: DEMO.Subject-1
func (js JetStream) AddStream(name string, subjects []string) error {
	// Get info about the stream
	stream, _ := js.GetStreamInfo(name)

	// If stream not found, create new
	if stream == nil {
		subjectNames := generateSubjectNames(name, subjects)
		_, err := js.instance.AddStream(&nats.StreamConfig{
			Name:         name,
			Subjects:     subjectNames,
			Retention:    nats.WorkQueuePolicy,
			MaxConsumers: -1,
			MaxMsgSize:   -1,
			MaxMsgs:      -1,
			NoAck:        false,
		})
		if err != nil {
			msg := fmt.Sprintf("[NATS JETSTREAM] - add stream error: %s", err.Error())
			return errors.New(msg)
		}
	}

	return nil
}

// DeleteStream ...
func (js JetStream) DeleteStream(name string) error {
	if err := js.instance.DeleteStream(name); err != nil {
		msg := fmt.Sprintf("[NATS JETSTREAM] - delete stream error: %s", err.Error())
		return errors.New(msg)
	}
	return nil
}

// AddStreamSubjects ...
func (js JetStream) AddStreamSubjects(name string, subjects []string) error {
	// Get info about the stream
	stream, _ := js.GetStreamInfo(name)
	if stream == nil {
		msg := fmt.Sprintf("[NATS JETSTREAM] - error when adding stream %s subjects: stream not found", name)
		return errors.New(msg)
	}

	// Merge current and new subjects
	subjectNames := generateSubjectNames(name, subjects)
	newSubjects := mergeAndUniqueArrayStrings(subjectNames, stream.Config.Subjects)

	_, err := js.instance.UpdateStream(&nats.StreamConfig{
		Name:     name,
		Subjects: newSubjects,
	})
	if err != nil {
		msg := fmt.Sprintf("[NATS JETSTREAM] - add stream error: %s", err.Error())
		return errors.New(msg)
	}
	return nil
}
