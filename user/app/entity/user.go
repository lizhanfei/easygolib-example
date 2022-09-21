package entity

type User struct {
	Id     int    `gorm:"column:id" json:"id" form:"id"`
	Name   string `gorm:"column:name" json:"name" form:"name"`
	Status uint8  `gorm:"column:status" json:"status" form:"status"` //1->可用；2->不可用
}

type Usertoken struct {
	UserId     uint   `gorm:"column:user_id" json:"user_id" form:"user_id"`
	Token      string `gorm:"column:token" json:"token" form:"token"`
	ExpireTime uint   `gorm:"column:expire_time" json:"expire_time" form:"expire_time"`
}
