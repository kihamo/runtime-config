package etcd

import (
	"github.com/kihamo/runtime-config/client"
	"github.com/kihamo/runtime-config/internal/store/etcd"
	"go.etcd.io/etcd/clientv3"
)

func NewStore(client *clientv3.Client) client.Store {
	return etcd.NewStore(client)
}
