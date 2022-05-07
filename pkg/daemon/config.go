package daemon

import (
	"bridgeswap/pkg/path"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	globalApiOption *ApiOption = nil
)

func SetGlobalApiOption(opt *ApiOption) {
	if globalApiOption != nil { //init once
		return
	}
	globalApiOption = opt
}

func GetGlobalApiOption() *ApiOption {
	return globalApiOption
}

func ParseConfig(configFile string, opt interface{}) error {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	v.SetConfigType("yaml")
	if configFile != "" {
		if !path.PathIsExist(configFile) {
			return fmt.Errorf("%s is not exist", configFile)
		}
		v.SetConfigFile(configFile)
	}
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

//注意尽量用一个单词，经过测试后用两个单词无法解析配置文件。
type ApiOption struct {
	Log     Log
	Server  Server
	Rsakey  string
	Statics string

	Etcd     EtcdOption
	Database DatabaseOption

	Outsource OutsourceOption
}

type MonitorOption struct {
	Log    Log
	Server Server
	Etcd   EtcdOption
}

type DatabaseOption struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

type OutsourceOption struct {
	Logo string
}
