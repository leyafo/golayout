package main

import (
	"golayout/internal/api"
	"golayout/pkg/daemon"
	"golayout/pkg/etcd"
	"golayout/pkg/httpctrl"
	"golayout/pkg/logger"
)

func main() {
	flags, err := daemon.ParseFlags()
	if err != nil {
		panic(err.Error())
	}

	apiOpt := daemon.ApiOption{}
	err = daemon.ParseConfig(flags.ConfigFile, &apiOpt)
	if err != nil {
		panic(err.Error())
	}

	err = logger.Init(logger.NewDefaultOption(apiOpt.Log.Debug, apiOpt.Log.Path))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Infof("configuration is: %+v", apiOpt)

	err = etcd.Init(apiOpt.Etcd.Endpoints)
	if err != nil {
		logger.Fatalf("init etcd failed:", err)
	}

	s := httpctrl.NewServer(apiOpt.Server.Listen, apiOpt.Server.Port)
	err = api.Init(s, &apiOpt.Etcd)
	if err != nil {
		logger.Fatal("init api failed:", err)
	}

	logger.Fatal(s.Run())
}
