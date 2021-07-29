package monitor

import (
	"fmt"
	"golayout/internal/monitor"
	"golayout/pkg/daemon"
	"golayout/pkg/logger"
	"google.golang.org/grpc"
	"log"
	"net"
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

	l, e := net.Listen("tcp", fmt.Sprintf("%s:%d", monitorOpt.Server.Listen, monitorOpt.Server.Port))
	if e != nil {
		log.Fatal("listen error:", e)
	}

	server := grpc.NewServer()
	monitor.RegisterRpc(server)
	log.Fatal(server.Serve(l))
}
