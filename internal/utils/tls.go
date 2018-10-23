package utils

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/pkg/errors"
)

// CreateTLSConfig sets up TLS config for usage with config providers
func CreateTLSConfig(serviceCertFile, serviceKeyFile, caCertFile string) (*tls.Config, error) {
	// Load client cert
	serviceCert, err := tls.LoadX509KeyPair(serviceCertFile, serviceKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load X509 key pair")
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read CA certificate")
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
