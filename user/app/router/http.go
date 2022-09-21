package router

import (
	"github.com/gin-gonic/gin"
	"user/app/controller"
)

func RegisterRouter(engine *gin.Engine) {
	router := engine.Group("/user")

	router.POST("/user/add", controller.AddUser)
	router.POST("/user/rm", controller.RmUser)
	router.POST("/user/login", controller.Login)
	router.GET("/user/auth", controller.Auth)
}
