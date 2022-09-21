package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lizhanfei/easygolib/utils/encrypt"
	"math/rand"
	"strconv"
	"time"
	"user/app/entity"
	"user/app/model"
)

const (
	maxToken        = 10
	tokenExpireTime = 30 * 24 * 3600
)

func NewUserServiceImplV1(ctx *gin.Context) User {
	return &UserImplV1{ctx: ctx}
}

type UserImplV1 struct {
	ctx *gin.Context
}

func (this *UserImplV1) AddUser(userName, password string) (*entity.User, error) {
	salt := this.getSalt()
	newUser := model.User{
		Name:       userName,
		Password:   encrypt.Md5(password + salt),
		Salt:       salt,
		Status:     1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	newUser, err := model.CreateUser(this.ctx, newUser)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		Id:     newUser.Id,
		Name:   newUser.Name,
		Status: newUser.Status,
	}, err
}

func (this *UserImplV1) RemoveUser(userId int) error {
	_, err := model.UpdateStatusById(this.ctx, model.UserUnValid, userId)
	return err
}

func (this *UserImplV1) Auth(token string) (bool, *entity.User, error) {
	userToken, err := model.GetInfoByToken(this.ctx, token, time.Now().Unix())
	if err != nil {
		return false, nil, err
	}
	if userToken.Id == 0 {
		return false, nil, nil
	}
	userInfo, err := model.GetUserInfoById(this.ctx, userToken.UserId)
	if err != nil {
		return false, nil, err
	}
	if userInfo.Id == 0 {
		return false, nil, nil
	}
	return true, &entity.User{
		Id:     userInfo.Id,
		Name:   userInfo.Name,
		Status: userInfo.Status,
	}, nil
}

func (this *UserImplV1) Login(userName, password string) (string, *entity.User, error) {
	userInfo, err := model.GetUserInfoByName(this.ctx, userName)
	if err != nil {
		return "", nil, err
	}
	if userInfo.Id == 0 || userInfo.Status == model.UserUnValid {
		return "", nil, nil
	}
	if encrypt.Md5(password+userInfo.Salt) != userInfo.Password {
		return "", nil, errors.New("login fail")
	}
	totalToken, err := model.CntUserToken(this.ctx, userInfo.Id, time.Now().Unix())
	if err != nil {
		return "", nil, err
	}
	if totalToken >= maxToken {
		return "", nil, errors.New("token over max")
	}
	newToken := this.makeUserToken(userInfo.Id)
	err = model.CreateUserToken(this.ctx, userInfo.Id, newToken, time.Now().Unix()+tokenExpireTime, time.Now().Unix())
	return newToken, &entity.User{
		Id:     userInfo.Id,
		Name:   userInfo.Name,
		Status: userInfo.Status,
	}, err
}

func (this *UserImplV1) getSalt() string {
	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Intn(10000) + 10000)
}

func (this *UserImplV1) makeUserToken(userId int) string {
	return encrypt.Md5(strconv.Itoa(userId) + this.getSalt() + strconv.FormatInt(time.Now().Unix(), 10))
}
