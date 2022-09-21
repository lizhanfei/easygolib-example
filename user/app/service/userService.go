package service

import "user/app/entity"

type User interface {
	AddUser(userName, password string) (*entity.User, error)
	RemoveUser(userId int) error
	Auth(token string) (bool, *entity.User, error)
	Login(userName, password string) (string, *entity.User, error)
}
