package conf

import (
	"github.com/lizhanfei/easygolib/utils/conf"
	"user/env"
)

var (
	DbConf   Resource
	HttpConf Server
	LogConf  LogConfig
	GrpcConf GrpcServerConf
)

type Resource struct {
	Mysql map[string]MysqlConf `yaml:"mysql"`
}

type Server struct {
	Server ServerConf `yaml:"server"`
}

func InitConf() {
	confYaml := conf.NewImplYaml()
	_ = confYaml.Load(env.ConfPath+"app/db.yaml", &DbConf)
	_ = confYaml.Load(env.ConfPath+"app/server.yaml", &HttpConf)
	_ = confYaml.Load(env.ConfPath+"app/log.yaml", &LogConf)
	_ = confYaml.Load(env.ConfPath+"app/grpc.yaml", &GrpcConf)

	return
}
