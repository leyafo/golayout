package main

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"golayout/internal/api"
	"golayout/pkg/daemon"
	"golayout/pkg/etcd"
	"golayout/pkg/httpctrl"
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
		logger.Fatalf("init etcd failed:", err)
		panic(err.Error())
	}

	s := httpctrl.NewServer(apiOpt.Server.Listen, apiOpt.Server.Port)
	err = api.InitApi(s, apiOpt.Etcd)
	if err != nil{
		logger.Fatal("init api failed:", err)
	}

	logger.Fatal(s.Run())
}
