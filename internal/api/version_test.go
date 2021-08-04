package api

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"golayout/pkg/daemon"
	"golayout/pkg/etcd"
	"golayout/pkg/logger"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	etcdCfg = &daemon.EtcdOption{
		Endpoints: []string{"http://172.16.238.100:2380","http://172.16.238.101:2380","http://172.16.238.102:2380"},
		Key:       "service/golayout/",
	}
	logger.InitLog(logger.NewDefaultOption(true, ""))

	err := etcd.InitEtcd(clientv3.Config{Endpoints: etcdCfg.Endpoints})
	if err != nil{
		os.Exit(-1)
	}
	m.Run()
}

func TestVersion(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Version))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil{
		t.Fatal("get request error: ", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK{
		t.Error("http response is not OK")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil{
		t.Error(err)
	}
	t.Log(body)
}
