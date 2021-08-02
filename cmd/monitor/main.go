package monitor

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golayout/internal/monitor"
	"golayout/pkg/daemon"
	"golayout/pkg/etcd"
	"golayout/pkg/logger"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	flags, err := daemon.ParseFlags()
	if err != nil{
		panic(err.Error())
	}

	monitorOpt := daemon.MonitorOption{}
	err = daemon.ParseConfig(flags.ConfigFile, &monitorOpt)
	if err != nil{
		panic(err.Error())
	}

	err = logger.InitLog(logger.NewDefaultOption(monitorOpt.Log.Debug, monitorOpt.Log.Path))
	if err != nil{
		panic(err)
	}
	defer logger.Sync()

	addr := fmt.Sprintf("%s:%d", monitorOpt.Server.Listen, monitorOpt.Server.Port)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}

	err = etcd.InitEtcd(clientv3.Config{
		Endpoints: monitorOpt.Etcd.Endpoints,
	})
	if err != nil{
		log.Fatal("init etcd failed: ", err)
	}

	server := grpc.NewServer()
	monitor.RegisterRpc(server)
	go log.Fatal(server.Serve(l))

	if err = etcd.ServiceAdd(monitorOpt.Etcd.EndpointsKey, addr); err != nil{
		log.Fatal("add etcd service failed: ", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<- c
}
