package natsio

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// AddStream add new stream, with default config
func AddStream(name string, subjects []string) error {
	// Get info about the stream
	stream, _ := natsJS.StreamInfo(name)
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
