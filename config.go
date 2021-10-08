package natsio

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
