package main

import (
	"context"
	"fmt"
	"github.com/kihamo/runtime-config"
	"github.com/kihamo/runtime-config/store/etcd"
	"github.com/kihamo/runtime-config/updater"
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

	updaterConfig, err := updater.NewUpdater(context.Background(), config.NewVersion(123, "release-1"), storeEtcd)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(updaterConfig)

	//variable := config.NewVariable("test", nil)
	//err = updaterConfig.AddVariable(context.Background(), variable)

	//fmt.Println(variable, err)
}
