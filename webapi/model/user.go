package model

import "time"

type LoginMessage struct {
	Phone      string   `json:"phone" xorm:"phone"`
	Password   string   `json:"password" xorm:"password"`
	Code       string   `json:"code" xorm:"-"`
}

type RegisterData struct {
	Phone      string   `json:"phone" xorm:"phone"`
	Password   string   `json:"password" xorm:"password"`
	RoleId     int      `json:"-" xorm:"role_id"`
}

func (rd RegisterData) TableName() string { return "gf_user" }

type UserInfo struct {
	Id             int64
	CreatedTime    time.Time    `json:"createdTime" xorm:"created_time"`
	UpdatedTime    time.Time    `json:"updatedTime" xorm:"updated_time"`
	UserName       string       `json:"userName" xorm:"user_name"`
	Password       string       `json:"password" xorm:"password"`
	Phone          string       `json:"phone" xorm:"phone"`
}

func (ui UserInfo) TableName() string { return "gf_user" }

type UpdateLoginPsw struct {
	OldPssword    string    `json:"oldPassword"`
	NewPssword    string    `json:"newPassword"`
	Code          string    `json:"code"`
}