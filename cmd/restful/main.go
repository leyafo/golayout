package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"golayout/internal/api"
	"golayout/pkg/daemon"
	"golayout/pkg/logger"
)

func main() {
	flags := daemon.NewCmdFlags()
	err := flags.Parse()
	if err != nil{
		panic(err.Error())
	}

	apiOpt := daemon.ApiOption{}
	err = daemon.ParseConfig(flags.ConfigFile, &apiOpt)
	if err != nil{
		panic(err.Error())
	}

	logOpt, err := logger.NewDefaultOption(apiOpt.Log.Debug, apiOpt.Log.Path)
	if err != nil{
		panic(err)
	}
	err = logger.InitLog(logOpt)
	if err != nil{
		panic(err)
	}
	defer logger.Sync()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/version", api.Version)

	listenAddr := fmt.Sprintf("%s:%d", apiOpt.Server.Listen, apiOpt.Server.Port)
	logger.Infof("serving http://%s", listenAddr)
	logger.Fatal(http.ListenAndServe(listenAddr, r))
}
