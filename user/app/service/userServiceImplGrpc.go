package service

import (
	"context"
	"time"
	"user/app/entity"
	"user/app/model"
	"user/helper"
)

type UserServiceImplGrpc struct {
}

func (this *UserServiceImplGrpc) Auth(ctx context.Context, req *entity.AuthRequest) (*entity.AuthResponse, error) {
	helper.Logger.Infof(ctx, "auth begin token:%s", req.Token)
	userToken, err := model.GetInfoByToken(ctx, req.Token, time.Now().Unix())
	if err != nil {
		helper.Logger.Warnf(ctx, "getTokenInfo fail, err:%s", err)
		return &entity.AuthResponse{}, err
	}
	if userToken.Id == 0 {
		helper.Logger.Warnf(ctx, "getTokenInfo res is empty")
		return &entity.AuthResponse{}, nil
	}
	if userToken.ExpireTime < time.Now().Unix() {
		helper.Logger.Warnf(ctx, "getTokenInfo res expire")
		return &entity.AuthResponse{}, nil
	}
	helper.Logger.Infof(ctx, "getTokenInfo res, res:%#v", userToken)
	userInfo, err := model.GetUserInfoById(ctx, userToken.UserId)
	if err != nil {
		helper.Logger.Warnf(ctx, "GetUserInfoById fail, err:%s", err)
		return &entity.AuthResponse{}, err
	}
	if userInfo.Id == 0 || userInfo.Status == model.UserUnValid {
		helper.Logger.Warnf(ctx, "GetUserInfoById res unvalid, res:%#v", userInfo)
		return &entity.AuthResponse{}, nil
	}
	helper.Logger.Infof(ctx, "GetUserInfoById res, res:%#v", userInfo)
	return &entity.AuthResponse{
		Id:     int32(userInfo.Id),
		Name:   userInfo.Name,
		Status: int32(userInfo.Status),
	}, nil
}
