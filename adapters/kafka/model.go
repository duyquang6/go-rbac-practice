package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

// Kafka define kafka config
type Kafka struct {
	Addrs              []string  `json:"addrs"`
	Topics             []string  `json:"topics"`
	Group              string    `json:"group"`
	Oldest             bool      `json:"oldest"`
	MaxMessageBytes    int       `json:"max_message_bytes"`
	Compress           bool      `json:"compress"`
	Newest             bool      `json:"newest"`
	PemFiles           *PemFiles `json:"pem_files"`
	InsecureSkipVerify bool      `json:"insecure_skip_verify"`
}

// PemFiles for tls
type PemFiles struct {
	ClientCert string `json:"client_cert"`
	ClientKey  string `json:"client_key"`
	CACert     string `json:"ca_cert"`
}

// NewTLSConfig for kafka
func NewTLSConfig(pemFiles PemFiles) (*tls.Config, error) {
	tlsConfig := tls.Config{}

	// Load client cert
	cert, err := tls.LoadX509KeyPair(pemFiles.ClientCert, pemFiles.ClientKey)
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	// Load CA cert
	caCert, err := ioutil.ReadFile(pemFiles.CACert)
	if err != nil {
		return &tlsConfig, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig.RootCAs = caCertPool

	tlsConfig.BuildNameToCertificate()
	return &tlsConfig, nil
}
