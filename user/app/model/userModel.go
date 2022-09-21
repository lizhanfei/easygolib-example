package model

import (
	"context"
	"time"
	"user/helper"
)

var userTableName = "users"

const (
	UserValid   = 1 //可用
	UserUnValid = 2 //不可用
)

type User struct {
	Id         int    `gorm:"column:id" json:"id" form:"id"`
	Name       string `gorm:"column:name" json:"name" form:"name"`
	Password   string `gorm:"column:password" json:"password" form:"password"`
	Salt       string `gorm:"column:salt" json:"salt" form:"salt"`
	Status     uint8  `gorm:"column:status" json:"status" form:"status"` //1->可用；2->不可用
	CreateTime int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time" form:"update_time"`
}

func GetUserInfoByName(ctx context.Context, userName string) (User, error) {
	res := User{}
	err := helper.UserClient.Table(userTableName).WithContext(ctx).
		Where("name=?", userName).Where("status=?", UserValid).Find(&res).Error
	return res, err
}

func GetUserInfoById(ctx context.Context, id int) (User, error) {
	res := User{}
	err := helper.UserClient.Table(userTableName).WithContext(ctx).
		Where("id=?", id).Find(&res).Error
	return res, err
}

func CreateUser(ctx context.Context, newUser User) (User, error) {
	err := helper.UserClient.Table(userTableName).WithContext(ctx).
		Create(&newUser).Error
	return newUser, err
}

func UpdateStatusById(ctx context.Context, statusNew, id int) (int64, error) {
	newData := map[string]interface{}{
		"status":      statusNew,
		"update_time": time.Now().Unix(),
	}
	res := helper.UserClient.Table(userTableName).WithContext(ctx).
		Where("id=?", id).
		UpdateColumns(newData)
	return res.RowsAffected, res.Error
}
