package helper

import (
	"user/conf"
	"github.com/gin-gonic/gin"
	"github.com/lizhanfei/easygolib/server/http"
)

func InitHttpServer(ginEngin *gin.Engine) {

	l, _ := http.NewListenerTcp(conf.HttpConf.Server.Address)

	s := http.NewServer(ginEngin, l, conf.HttpConf.Server.ReadTimeout, conf.HttpConf.Server.WriteTimeout, conf.HttpConf.Server.CloseTimeout)
	s.Run()
}
