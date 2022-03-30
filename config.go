package natsio

import "time"

// Config ...
type Config struct {
	// Connect url
	URL string

	// Auth user
	User string

	// Auth password
	Password string

	// TLS config
	TLS *TLSConfig

	// RequestTimeout
	RequestTimeout time.Duration
}

// TLSConfig ...
type TLSConfig struct {
	// Cert file
	CertFilePath string

	// Key file
	KeyFilePath string

	// Root CA
	RootCAFilePath string
}
