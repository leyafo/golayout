package main

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"golayout/internal/api"
	"golayout/pkg/daemon"
	"golayout/pkg/etcd"
	httpServer "golayout/pkg/http"
	"golayout/pkg/logger"
)

func main() {
	flags, err := daemon.ParseFlags()
	if err != nil{
		panic(err.Error())
	}

	apiOpt := daemon.ApiOption{}
	err = daemon.ParseConfig(flags.ConfigFile, &apiOpt)
	if err != nil{
		panic(err.Error())
	}

	err = logger.InitLog(logger.NewDefaultOption(apiOpt.Log.Debug, apiOpt.Log.Path))
	if err != nil{
		panic(err)
	}
	defer logger.Sync()

	err = etcd.InitEtcd(clientv3.Config{
		Endpoints: apiOpt.Etcd.Endpoints,
	})
	if err != nil {
		logger.Fatalf("init etcd failed: ", err)
		panic(err.Error())
	}

	businessLogical, err := api.NewServer("monitor", &apiOpt)
	if err != nil{
		logger.Fatal("init api server: ", err)
	}

	s := httpServer.NewServer(apiOpt.Server.Listen, apiOpt.Server.Port)
	s.RouterRegister(businessLogical)
	logger.Fatal(s.Run())
}
