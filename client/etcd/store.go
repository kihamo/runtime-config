package etcd

import (
	"github.com/kihamo/runtime-config/client"
	"github.com/kihamo/runtime-config/internal/store/etcd"
	"go.etcd.io/etcd/clientv3"
)

// Config is used for etcd client initialization
type Config etcd.Config

// NewStore creates new instance of Store from Config
func NewStore(c *Config) (client.Store, error) {
	config := etcd.Config(*c)
	return etcd.NewStore(&config)
}

// NewStoreFromClient creates new instance of Store with existing etcd client
func NewStoreFromClient(client *clientv3.Client) client.Store {
	return etcd.NewStoreFromClient(client)
}
