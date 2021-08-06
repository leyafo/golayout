package api

import (
	"context"
	"golayout/pkg/daemon"
	"golayout/pkg/etcd"
	"golayout/pkg/httpctrl"
	"google.golang.org/grpc"
	"path"
)

var (
	ctrlServer *httpctrl.Server
	etcdCfg    *daemon.EtcdOption
)

func Init(s *httpctrl.Server, etcdConfig *daemon.EtcdOption) error {
	etcdCfg = etcdConfig
	etcd.Init(etcdConfig.Endpoints)
	if s == nil { //for test usage
		return nil
	}
	return s.RoutersRegister("v1", []httpctrl.Router{
		{Method: "GET", Path: "/version", Handler: Version, Doc: VersionDoc},
		{Method: "GET", Path: "/doc", Handler: Doc, Doc: nil},
	})
}

const (
	serviceMonitor = "monitor"
)

func getServiceConnection(service string) (*grpc.ClientConn, error) {
	var err error
	lb, err := etcd.NewLoadBalancer(path.Clean(etcdCfg.Key+"/"+service), etcd.RoundRobin)
	if err != nil {
		panic(err)
	}
	return etcd.DialGrpc(context.TODO(), lb, grpc.WithInsecure())
}
