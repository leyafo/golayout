package etcd

import (
	"fmt"
	"golayout/pkg/daemon"
	"golayout/pkg/logger"
	"os"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	etcdOption = daemon.EtcdOption{
		Endpoints: []string{"http://172.16.238.100:2380", "http://172.16.238.101:2380", "http://172.16.238.102:2380"},
		Key:       "service/golayout/test",
	}
)

func TestMain(m *testing.M) {
	etcdCfg := clientv3.Config{
		Endpoints: etcdOption.Endpoints,
	}
	err := Init(etcdCfg.Endpoints)
	if err != nil {
		fmt.Printf("init etcd failed: %v", err)
		os.Exit(-1)
	}
	err = logger.Init(logger.NewDefaultOption(false, ""))
	if err != nil {
		fmt.Printf("init logger failed: %v", err)
		os.Exit(-1)
	}
	code := m.Run()
	os.Exit(code)
}

func TestServiceAdd(t *testing.T) {
	registerAddr := "1.2.3.4"
	err := ServiceAdd(etcdOption.Key, registerAddr)
	if err != nil {
		t.Error(err)
	}
	list, err := ServiceList(etcdOption.Key)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
	arrayHas := func(arr []string, str string) bool {
		for _, e := range arr {
			if e == str {
				return true
			}
		}
		return false
	}
	if !arrayHas(list, registerAddr) {
		t.Errorf("list service is not equal register, want=%s, out=%s", registerAddr, list[0])
	}

	ServiceDelete(etcdOption.Key)
	if err != nil {
		t.Error(err)
	}
	list, err = ServiceList(etcdOption.Key)
	if err != nil {
		t.Error(err)
	}
	if arrayHas(list, registerAddr) {
		t.Errorf("call service delete failed: %v", err)
	}
}
