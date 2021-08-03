package api

import (
	"golayout/pkg/daemon"
	"golayout/pkg/httpctrl"
)

func InitApi(s *httpctrl.Server, etcdCfg daemon.EtcdOption)error{
	return s.RoutersRegister("v1", []httpctrl.Router{
		{ Method: "GET", Path:"/version", Handler: Version, Doc: VersionDoc},
	})
}