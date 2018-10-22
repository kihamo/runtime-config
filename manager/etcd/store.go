package etcd

import (
	"github.com/kihamo/runtime-config/internal/store/etcd"
	"github.com/kihamo/runtime-config/manager"
	"go.etcd.io/etcd/clientv3"
)

func NewStore(client *clientv3.Client) manager.Store {
	return etcd.NewStore(client)
}
