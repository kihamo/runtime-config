package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/etcd-io/etcd/clientv3"
	"github.com/kihamo/runtime-config/internal"
	"github.com/kihamo/runtime-config/internal/store/etcd"
	"github.com/kihamo/runtime-config/internal/updater"
)

const (
	Endpoints = "http://localhost:2379"
	ProjectID = 123
	VersionId = "release-1"
)

func main() {
	clientEtcd, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(Endpoints, ";"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	storeEtcd := etcd.NewStore(clientEtcd)

	updaterConfig, err := updater.NewUpdater(context.Background(), internal.NewVersion(strconv.FormatUint(ProjectID, 10), VersionId), storeEtcd)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(updaterConfig)

	variable := internal.NewVariable("test", nil)
	err = updaterConfig.AddVariable(context.Background(), variable)

	fmt.Println(variable, err)
}
