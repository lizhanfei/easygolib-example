package controller

type CreateReq struct {
	UserName string `form:"userName" json:"userName"`
	Password string `form:"passWord" json:"passWord"`
}

type RmReq struct {
	UserId int `form:"userId" json:"userId"`
}

type AuthReq struct {
	Token string `form:"token" json:"token"`
}

type LoginReq struct {
	UserName string `form:"userName" json:"userName"`
	Password string `form:"passWord" json:"passWord"`
}
