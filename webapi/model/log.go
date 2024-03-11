package model

import "time"

type OperateLog struct {
	UserId        int64         `xorm:"user_id"`
	Addr          string        `xorm:"addr"`
	LogLevel      string        `xorm:"log_level"` 
	OperateTime   time.Time     `xorm:"operate_time"`
	OperateDesc   string        `xorm:"operate_desc"` // 操作描述
	Detail        string        `xorm:"detail"`   // 操作详情
}

func (ol OperateLog) TableName() string { return "gf_log" }

type LogList struct {
	Account       string    `json:"phone" xorm:"phone"`
	Addr          string    `json:"addr" xorm:"addr"`
	LogLevel      string    `json:"logLevel" xorm:"log_level"`
	OperateTime   string    `json:"operateTime" xorm:"operate_time"`
	OperateDesc   string    `json:"operateDesc" xorm:"operate_desc"`
	Detail        string    `json:"detail" xorm:"detail"`
}