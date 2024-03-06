package model

import (
	"time"
)

type UploadUserVideo struct {
	VideoNo        string       `form:"-" xorm:"video_no"`
	UserId         int64        `form:"-" xorm:"user_id"`
	Title          string       `form:"title" xorm:"title"`
	Introduction   string       `form:"introduction" xorm:"introduction"`
	UploadTime     time.Time    `form:"-" xorm:"upload_time"`
	// SuccessTime    time.Time    `form:"-" xorm:"success_time"`
	Video          string       `form:"video" xorm:"-"`
}

func (uuv UploadUserVideo) TableName() string { return "gf_video" }

type UserVideoList struct {
	VideoNo        string     `json:"videoNo" xorm:"video_no"`
	Title          string     `json:"title" xorm:"title"`
}

func (uvl UserVideoList) TableName() string { return "gf_video" }

type UserVideoDetail struct {
	VideoNo        string     `json:"videoNo" xorm:"video_no"`
	Title          string     `json:"title" xorm:"title"`
	Introduction   string     `json:"introduction" xorm:"introduction"`
	UploadTime     string     `json:"uploadTime" xorm:"upload_time"`
	IsDel          int        `json:"-" xorm:"is_del"`
	Upvote         int        `json:"upvote" xorm:"upvote"`
	Disagree       int        `json:"disagree" xorm:"disagree"`
	Coins          int        `json:"coins" xorm:"coins"`
	Collect        int        `json:"collect" xorm:"collect"`
}

func (uvd UserVideoDetail) TableName() string { return "gf_video" }