package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/kihamo/runtime-config/client"
	"github.com/kihamo/runtime-config/client/environment"
	"github.com/kihamo/runtime-config/internal"
)

const (
	Endpoints = "http://localhost:2379"
	ProjectID = 123
	VersionId = "release-1"
)

func main() {
	ctx := context.Background()
	version := internal.NewVersion(strconv.FormatUint(ProjectID, 10), VersionId)

	/*
		clientEtcd, err := clientv3.New(clientv3.Config{
			Endpoints: strings.Split(Endpoints, ";"),
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		storeEtcd := etcd.NewStore(clientEtcd)
	*/
	storeEnv := environment.NewStore("config_")

	clientConfig, err := client.NewClient(ctx, version, storeEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(clientConfig.Values(ctx))

	val, err := clientConfig.Value(ctx, "tEst")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(val.String(), val.IsNil())
}
