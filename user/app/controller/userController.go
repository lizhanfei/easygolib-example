package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user/app/service"
)

func AddUser(ctx *gin.Context) {
	var (
		params = new(CreateReq)
	)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	userService := service.NewUserServiceImplV1(ctx)
	res, err := userService.AddUser(params.UserName, params.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]int{
		"id": res.Id,
	})
}

func RmUser(ctx *gin.Context) {
	var (
		params = new(RmReq)
	)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	userService := service.NewUserServiceImplV1(ctx)
	err := userService.RemoveUser(params.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"res": "succ",
	})
}

func Auth(ctx *gin.Context) {
	var (
		params = new(AuthReq)
	)
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	userService := service.NewUserServiceImplV1(ctx)
	res, userEntity, err := userService.Auth(params.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if !res {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"info": userEntity,
	})
}

func Login(ctx *gin.Context) {
	var (
		params = new(LoginReq)
	)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	userService := service.NewUserServiceImplV1(ctx)
	token, userEntity, err := userService.Login(params.UserName, params.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"id":    userEntity.Id,
	})
}
