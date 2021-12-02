package natsio

import (
	"errors"
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/nats-io/nats.go"
)

var (
	natsClient *nats.Conn
	natsJS     nats.JetStreamContext
)

// Connect ...
func Connect(cfg Config) error {
	if cfg.URL == "" {
		return errors.New("connect URL is required")
	}

	// Connect options
	opts := make([]nats.Option, 0)

	// Has authentication
	if cfg.User != "" {
		opts = append(opts, nats.UserInfo(cfg.User, cfg.Password))
	}

	// If it has TLS
	if cfg.TLS != nil {
		opts = append(opts, nats.ClientCert(cfg.TLS.CertFilePath, cfg.TLS.KeyFilePath))
		opts = append(opts, nats.RootCAs(cfg.TLS.RootCAFilePath))
	}

	nc, err := nats.Connect(cfg.URL, opts...)
	if err != nil {
		msg := fmt.Sprintf("error when connecting to NATS: %s", err.Error())
		return errors.New(msg)
	}

	fmt.Println(aurora.Green("*** CONNECTED TO NATS: " + cfg.URL))

	// Set client
	natsClient = nc

	// Create jet stream context
	js, err := natsClient.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		msg := fmt.Sprintf("error when create NATS JetStream: %s", err.Error())
		return errors.New(msg)
	}
	natsJS = js

	return nil
}

// GetClient ...
func GetClient() *nats.Conn {
	return natsClient
}

// GetJetStream ...
func GetJetStream() nats.JetStreamContext {
	return natsJS
}
