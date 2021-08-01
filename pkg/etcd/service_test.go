package etcd

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golayout/pkg/daemon"
	"golayout/pkg/logger"
	"os"
	"testing"
)

var(
	EndpointsKey = "service/golayout/test"
	etcdOption = daemon.EtcdOption{
		Endpoints:  []string{"http://172.16.238.100:2380","http://172.16.238.101:2380","http://172.16.238.102:2380"},
		EndpointsKey: EndpointsKey,
	}
)

func TestMain(m *testing.M) {
	etcdCfg := clientv3.Config{
		Endpoints:            etcdOption.Endpoints,
	}
	err := InitEtcd(etcdCfg)
	if err != nil{
		fmt.Printf("init etcd failed: %v", err)
		os.Exit(-1)
	}
	err = logger.InitLog(logger.NewDefaultOption(false, ""))
	if err != nil{
		fmt.Printf("init logger failed: %v", err)
		os.Exit(-1)
	}
	code := m.Run()
	os.Exit(code)
}

func TestServiceAdd(t *testing.T){
	registerAddr := "1.2.3.4"
	err := ServiceAdd(EndpointsKey, registerAddr)
	if err != nil{
		t.Error(err)
	}
	list, err := ServiceList(EndpointsKey)
	if err != nil{
		t.Error(err)
	}
	t.Log(list)
	if list[0] != registerAddr{
		t.Errorf("list service is not equal register, want=%s, out=%s", registerAddr, list[0])
	}
}

