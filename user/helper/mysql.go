package helper

import (
	"github.com/lizhanfei/easygolib/log"
	"gorm.io/gorm/logger"
	"user/conf"
	mysqlLib "github.com/lizhanfei/easygolib/db/mysql"
	"gorm.io/gorm"
	"user/env"
)

var (
	MysqlClient  map[string]*gorm.DB
	UserClient *gorm.DB
)

func initDb() {
	MysqlClient = make(map[string]*gorm.DB)
	dbLog := log.DbLog{LogLevel: logger.Info,}
	dbLog.SetLogger(Logger)
	dbLog.SetCkRequestId(env.ContextKeyRequestID)
	dbLog.SetAppName(env.AppName)

	for name, mysqlConf := range conf.DbConf.Mysql {
		clientTmp, err := mysqlLib.InitMysql(&mysqlLib.MysqlConf{
			DataBase:        mysqlConf.DataBase,
			Addr:            mysqlConf.Addr,
			User:            mysqlConf.User,
			Password:        mysqlConf.Password,
			Charset:         mysqlConf.Charset,
			MaxIdleConns:    mysqlConf.MaxIdleConns,
			MaxOpenConns:    mysqlConf.MaxOpenConns,
			ConnMaxIdlTime:  mysqlConf.ConnMaxIdlTime,
			ConnMaxLifeTime: mysqlConf.ConnMaxLifeTime,
			ConnTimeOut:     mysqlConf.ConnTimeOut,
			WriteTimeOut:    mysqlConf.WriteTimeOut,
			ReadTimeOut:     mysqlConf.ReadTimeOut,
		}, &dbLog)
		if err != nil {
			panic(err)
		}
		if "user" == name {
			UserClient = clientTmp
		}
		MysqlClient[name] = clientTmp
	}
}
