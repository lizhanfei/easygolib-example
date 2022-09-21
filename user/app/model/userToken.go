package model

import (
	"context"
	"user/helper"
)

var userTokenTableName = "userToken"

type UserToken struct {
	Id         uint64 `gorm:"column:id" json:"id" form:"id"`
	UserId     int    `gorm:"column:user_id" json:"user_id" form:"user_id"`
	Token      string `gorm:"column:token" json:"token" form:"token"`
	ExpireTime int64  `gorm:"column:expire_time" json:"expire_time" form:"expire_time"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`
}

func GetInfoByToken(ctx context.Context, token string, lastTime int64) (UserToken, error) {
	res := UserToken{}
	db := helper.UserClient.Table(userTokenTableName).WithContext(ctx).
		Where("token=?", token)
	if lastTime > 0 {
		db = db.Where("expire_time>?", lastTime)
	}
	err := db.Find(&res).Error
	return res, err
}

func CreateUserToken(ctx context.Context, userId int, token string, expireTime, createTime int64) error {
	newToken := UserToken{
		UserId:     userId,
		Token:      token,
		ExpireTime: expireTime,
		CreateTime: createTime,
	}
	err := helper.UserClient.Table(userTokenTableName).WithContext(ctx).
		Create(&newToken).Error
	return err
}

func CntUserToken(ctx context.Context, userId int, lastTime int64) (int64, error) {
	var cnt int64
	err := helper.UserClient.Table(userTokenTableName).WithContext(ctx).
		Where("user_id=?", userId).Where("expire_time>?", lastTime).
		Count(&cnt).Error
	return cnt, err
}
