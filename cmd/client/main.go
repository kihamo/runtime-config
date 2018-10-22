package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/kihamo/runtime-config"
	"github.com/kihamo/runtime-config/client"
	"github.com/kihamo/runtime-config/internal/store/etcd"
	"go.etcd.io/etcd/clientv3"
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

	clientConfig, err := client.NewClient(context.Background(), config.NewVersion(ProjectID, VersionId), storeEtcd)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(clientConfig)
}
