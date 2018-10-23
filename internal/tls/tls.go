package tls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

// NewConfig sets up TLS config for usage with config providers
func NewConfig(serviceCertFile, serviceKeyFile, caCertFile string) (*tls.Config, error) {
	// Load client cert
	serviceCert, err := tls.LoadX509KeyPair(serviceCertFile, serviceKeyFile)
	if err != nil {
		return nil, fmt.Errorf("unable to load X509 key pair: %v", err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read CA certificate: %v", err)
	}

	// Create TLS config
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serviceCert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()

	return tlsConfig, nil
}
