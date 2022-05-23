package main

import (
	"fmt"
	"golayout/internal/api"
	"golayout/pkg/daemon"
	"golayout/pkg/logger"
	"net/http"
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
	daemon.SetGlobalApiOption(&apiOpt)

	err = logger.Init(logger.NewDefaultOption(apiOpt.Log.Debug, apiOpt.Log.Path))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Infof("configuration is: %+v", apiOpt)

	addr := fmt.Sprintf("%s:%d", apiOpt.Server.Listen, apiOpt.Server.Port)
	logger.Infof("Starting server on %v\n", addr)
	logger.Fatal(http.ListenAndServe(addr, api.Router()))
}
