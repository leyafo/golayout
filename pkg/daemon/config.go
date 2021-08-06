package daemon

import (
	"github.com/spf13/viper"
)

func ParseConfig(configFile string, opt interface{}) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		return nil
	}
	return v.Unmarshal(opt)
}

type Log struct {
	Path  string
	Debug bool
}

type Server struct {
	Listen string
	Port   int
}

type EtcdOption struct {
	Endpoints []string
	Key       string
}

type ApiOption struct {
	Log    Log
	Server Server

	Etcd EtcdOption
}

type MonitorOption struct {
	Log    Log
	Server Server
	Etcd   EtcdOption
}
