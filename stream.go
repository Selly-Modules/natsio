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
func AddStream(name string, subjects []string) error {
	// Get info about the stream
	stream, _ := GetStreamInfo(name)
	// If stream not found, create new
	if stream == nil {
		_, err := natsJS.AddStream(&nats.StreamConfig{
			Name:     name,
			Subjects: subjects,
			Storage:  nats.FileStorage,
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
	newSubjects := mergeAndUniqueArrayStrings(subjects, stream.Config.Subjects)

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
