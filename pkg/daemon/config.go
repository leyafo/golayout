package daemon

import (
	"fmt"
	"golayout/pkg/path"
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
	if globalApiOption == nil {
		panic("globalApiOption is nil")
	}
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
	Name   string
	Schema string
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

//DBString dbType is "mysql", "pg", "sqlite"
func (dbOption DatabaseOption) DBString(dbType string) string {
	if dbType == "mysql" {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbOption.User, dbOption.Password, dbOption.Host,
			dbOption.Port, dbOption.Name)
	}
	return ""
}

type OutsourceOption struct {
	Logo string
}
