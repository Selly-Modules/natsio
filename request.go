package natsio

import (
	"time"

	"github.com/nats-io/nats.go"
)

// Default timeout 10s
const requestTimeout = 10 * time.Second

// Request ...
func Request(subject string, payload []byte) (*nats.Msg, error) {
	return natsClient.Request(subject, payload, requestTimeout)
}
