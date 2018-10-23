package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kihamo/runtime-config/internal"
	"github.com/kihamo/runtime-config/manager"
	"github.com/kihamo/runtime-config/manager/etcd"
	"go.etcd.io/etcd/clientv3"
)

const (
	Endpoints = "http://localhost:2379"
	ProjectID = 123
	VersionId = "release-1"
)

func main() {
	ctx := context.Background()
	version := internal.NewVersion(strconv.FormatUint(ProjectID, 10), VersionId)

	clientEtcd, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(Endpoints, ";"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	storeEtcd := etcd.NewStore(clientEtcd)

	managerConfig := manager.NewManager(storeEtcd)

	fmt.Println(managerConfig.Values(ctx, version))
}
