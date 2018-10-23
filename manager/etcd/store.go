package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"github.com/kihamo/runtime-config/internal/store/etcd"
	"github.com/kihamo/runtime-config/manager"
)

func NewStore(client *clientv3.Client) manager.Store {
	return etcd.NewStore(client)
}
