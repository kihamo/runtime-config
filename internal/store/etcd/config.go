package etcd

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

// DefaultTimeout is the default dial timeout
const DefaultTimeout = 5 * time.Second

// Config is used for etcd client initialization
type Config struct {
	Endpoints   string
	ServiceCert string
	ServiceKey  string
	CACert      string
	Timeout     time.Duration

	tlsEnabled bool
	tlsConfig  *tls.Config
}

func (c *Config) setDefaults() {
	if c.Timeout == 0 {
		c.Timeout = DefaultTimeout
	}

	if c.ServiceCert != "" && c.ServiceKey != "" && c.CACert != "" {
		c.tlsEnabled = true
	}
}

func validateConfig(c *Config) error {
	if c == nil {
		return errors.New("config is empty")
	}

	if c.Endpoints == "" {
		return errors.New("endpoints are missing")
	}

	if c.ServiceCert == "" && c.ServiceKey != "" && c.CACert != "" {
		return errors.New("service certificate is missing")
	}

	if c.ServiceCert != "" && c.ServiceKey == "" && c.CACert != "" {
		return errors.New("service key is missing")
	}

	if c.ServiceCert != "" && c.ServiceKey != "" && c.CACert == "" {
		return errors.New("CA certificate is missing")
	}

	return nil
}

func newTLSConfig(serviceCertFile, serviceKeyFile, caCertFile string) (*tls.Config, error) {
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
