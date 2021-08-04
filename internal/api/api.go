package api

import (
	"golayout/pkg/daemon"
	"golayout/pkg/httpctrl"
)

var (
	ctrlServer *httpctrl.Server
	etcdCfg *daemon.EtcdOption
)

func InitApi(s *httpctrl.Server, etcdConfig daemon.EtcdOption)error{
	ctrlServer = s
	etcdCfg = &etcdConfig
	return s.RoutersRegister("v1", []httpctrl.Router{
		{ Method: "GET", Path:"/version", Handler: Version, Doc: VersionDoc},
		{ Method: "GET", Path:"/doc", Handler: Doc, Doc: nil},
	})
}