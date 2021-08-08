package main

import (
	"fmt"
	"golayout/internal/monitor"
	"golayout/pkg/daemon"
	"golayout/pkg/etcd"
	"golayout/pkg/logger"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func main() {
	flags, err := daemon.ParseFlags()
	if err != nil {
		panic(err.Error())
	}

	monitorOpt := daemon.MonitorOption{}
	err = daemon.ParseConfig(flags.ConfigFile, &monitorOpt)
	if err != nil {
		panic(err.Error())
	}

	err = logger.Init(logger.NewDefaultOption(monitorOpt.Log.Debug, monitorOpt.Log.Path))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Infof("configuration is %+v", monitorOpt)

	addr := fmt.Sprintf("%s:%d", monitorOpt.Server.Listen, monitorOpt.Server.Port)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		logger.Fatal("listen error:", e)
	}
	logger.Infof("listening %s ...", addr)

	server := grpc.NewServer()
	monitor.RegisterServerStub(server)
	go server.Serve(l)

	err = etcd.Init(monitorOpt.Etcd.Endpoints)
	if err != nil {
		panic(err)
	}

	registerAddr := addr
	if flags.RegisterAddr != ""{
		 registerAddr =  flags.RegisterAddr
	}
	err = etcd.ServiceAdd(monitorOpt.Etcd.Key, registerAddr)
	if err != nil {
		panic(err)
	}
	defer etcd.ServiceDelete(monitorOpt.Etcd.Key)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
