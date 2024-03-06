package model

import "time"

type VideoInput struct {
	Sec     int64  `json:"sec"`
	Caption string `json:"caption"`
	VideoNo string `json:"videoNo"`
}

type CacheVideoInput struct {
	VideoNo    string     `json:"videoNo" xorm:"video_no"`
	UserId     int64      `json:"userId" xorm:"user_id"`
	SendTime   time.Time  `json:"sendTime" xorm:"send_time"`
	Caption    string     `json:"caption" xorm:"caption"`
	Second     int64      `json:"second" xorm:"second"`
}

func (cvi CacheVideoInput) TableName() string { return "gf_video_caption" }

type NodeInputData struct {
	VideoNo   string   `json:"videoNo"`
	Second    int64    `json:"second"`
}