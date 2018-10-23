package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"github.com/kihamo/runtime-config/client"
	"github.com/kihamo/runtime-config/internal/store/etcd"
)

func NewStore(client *clientv3.Client) client.Store {
	return etcd.NewStore(client)
}
