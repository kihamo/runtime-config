package etcd

import (
	"crypto/tls"
	"time"

	"github.com/pkg/errors"
)

// DefaultTimeout is the default dial timeout
const DefaultTimeout = 5 * time.Second

// Config describes configuration for the etcd provider
type Config struct {
	Endpoints   string
	Timeout     time.Duration
	ServiceCert string
	ServiceKey  string
	CACert      string

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
