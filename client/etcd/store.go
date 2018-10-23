package etcd

import (
	"fmt"

	"github.com/kihamo/runtime-config/client"
	"github.com/kihamo/runtime-config/internal/store/etcd"
	"go.etcd.io/etcd/clientv3"
)

// Config is used for etcd client initialization
type Config etcd.Config

// NewStore creates new instance of Store from Config
func NewStore(config *Config) (client.Store, error) {
	c := etcd.Config(*config)
	client, err := etcd.NewEtcdClient(&c)
	if err != nil {
		return nil, fmt.Errorf("unable to create new etcd store: %v", err)
	}

	return etcd.NewStore(client), nil
}

// NewStoreFromClient creates new instance of Store with existing etcd client
func NewStoreFromClient(client *clientv3.Client) client.Store {
	return etcd.NewStore(client)
}
