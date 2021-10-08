package natsio

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// GetStreamInfo ...
func GetStreamInfo(name string) (*nats.StreamInfo, error) {
	return natsJS.StreamInfo(name)
}

// AddStream add new stream, with default config
// Due to subject must have a unique name, subject name will be combined with stream name
// E.g: stream name is "DEMO", subject name is "Subject-1", so final name in NATS will be: DEMO.Subject-1
func AddStream(name string, subjects []string) error {
	// Get info about the stream
	stream, _ := GetStreamInfo(name)

	// If stream not found, create new
	if stream == nil {
		subjectNames := generateSubjectNames(name, subjects)
		_, err := natsJS.AddStream(&nats.StreamConfig{
			Name:     name,
			Subjects: subjectNames,
		})
		if err != nil {
			msg := fmt.Sprintf("add stream error: %s", err.Error())
			return errors.New(msg)
		}
	}

	return nil
}

// DeleteStream ...
func DeleteStream(name string) error {
	if err := natsJS.DeleteStream(name); err != nil {
		msg := fmt.Sprintf("delete stream error: %s", err.Error())
		return errors.New(msg)
	}
	return nil
}

// AddStreamSubjects ...
func AddStreamSubjects(name string, subjects []string) error {
	// Get info about the stream
	stream, _ := GetStreamInfo(name)
	if stream == nil {
		msg := fmt.Sprintf("error when adding stream %s subjects: stream not found", name)
		return errors.New(msg)
	}

	// Merge current and new subjects
	subjectNames := generateSubjectNames(name, subjects)
	newSubjects := mergeAndUniqueArrayStrings(subjectNames, stream.Config.Subjects)

	_, err := natsJS.UpdateStream(&nats.StreamConfig{
		Name:     name,
		Subjects: newSubjects,
	})
	if err != nil {
		msg := fmt.Sprintf("add stream error: %s", err.Error())
		return errors.New(msg)
	}
	return nil
}
