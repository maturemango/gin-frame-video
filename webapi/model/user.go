package model

import "time"

type LoginMessage struct {
	Account  string `json:"account" xorm:"account"`
	Password string `json:"password" xorm:"password"`
}

type RegisterData struct {
	Account  string `json:"account" xorm:"account"`
	Password string `json:"password" xorm:"password"`
	RoleId   int    `json:"-" xorm:"role_id"`
}

func (rd RegisterData) TableName() string { return "gf_user" }

type UserInfo struct {
	Id             int64
	CreatedTime    time.Time    `json:"createdTime" xorm:"created_time"`
	UpdatedTime    time.Time    `json:"updatedTime" xorm:"updated_time"`
	UserName       string       `json:"userName" xorm:"user_name"`
	Password       string       `json:"password" xorm:"password"`
	Account        string       `json:"account" xorm:"account"`
}

func (ui UserInfo) TableName() string { return "gf_user" }