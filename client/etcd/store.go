package etcd

import (
	"strings"

	"github.com/kihamo/runtime-config/client"
	"github.com/kihamo/runtime-config/internal/store/etcd"
	"github.com/kihamo/runtime-config/internal/utils"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
)

// NewStore creates new instance of Store from Config
func NewStore(c *Config) (client.Store, error) {
	return newStoreFromConfig(c)
}

// NewStoreFromClient creates new instance of Store with existing etcd client
func NewStoreFromClient(client *clientv3.Client) client.Store {
	return etcd.NewStore(client)
}

func newStoreFromConfig(c *Config) (client.Store, error) {
	// Validate our config
	if err := validateConfig(c); err != nil {
		return nil, errors.Wrap(err, "unable to validate config")
	}

	// Setup etcd config
	c.setDefaults()
	endpoints := strings.Split(c.Endpoints, ",")
	etcdConfig := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: c.Timeout,
	}

	// Create secure connection if certs were provided
	if c.tlsEnabled {
		tlsConfig, err := utils.CreateTLSConfig(c.ServiceCert, c.ServiceKey, c.CACert)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create TLS config")
		}

		etcdConfig.TLS = tlsConfig
	}

	// Setup client
	etcdClient, err := clientv3.New(etcdConfig)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create new etcd client")
	}

	return etcd.NewStore(etcdClient), nil
}
