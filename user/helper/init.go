package helper

import "user/conf"

func Init() {
	conf.InitConf()
	InitLog()
	initDb()
}