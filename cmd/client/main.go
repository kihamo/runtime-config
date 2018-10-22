package main

import (
	"context"
	"fmt"
	"github.com/kihamo/runtime-config"
	"github.com/kihamo/runtime-config/client"
	"github.com/kihamo/runtime-config/store/etcd"
	"go.etcd.io/etcd/clientv3"
	"log"
)

func main() {
	clientEtcd, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://localhost:2379"},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	storeEtcd := etcd.NewStore(clientEtcd)

	clientConfig, err := client.NewClient(context.Background(), config.NewVersion(123, "release-1"), storeEtcd)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(clientConfig)
}
