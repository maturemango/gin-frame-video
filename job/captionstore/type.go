package captionstore

import "time"

var VideoInputKey string = "danmu:Video:No.%s;Node:" // 视频弹幕

type CacheInputData struct {
	VideoNo     string     `json:"videoNo" xorm:"video_no"`
	UserId      int64      `json:"userId" xorm:"user_id"`
	SendTime    time.Time  `json:"sendTime" xorm:"send_time"`
	Caption     string     `json:"caption" xorm:"caption"`
	Second      int64      `json:"second" xorm:"second"`
}

func (cid CacheInputData) TableName() string { return "gf_video_caption" }