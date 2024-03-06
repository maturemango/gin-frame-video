package model

type VideoTriplet struct {
	Type        int      `json:"type"`   // 根据传入的type判断执行的操作
	Controller  int      `json:"controller"`   // 根据传入的controller判断执行还是取消操作
	VideoNo     string   `json:"videoNo"`
}

type UserTripletControls struct {
	VideoNo       string      `json:"videoNo" xorm:"video_no"`
	UserId        int64       `json:"userId" xorm:"user_id"`
	IsUpvote      int         `json:"isUpvote" xorm:"is_upvote"`
	IsDisagree    int         `json:"isDisagree" xorm:"is_disagree"`
	IsCoins       int         `json:"isCoins" xorm:"is_coins"`
	IsCollect     int         `json:"isCollect" xorm:"is_collect"`
	IsTriplet     int         `json:"isTriplet" xorm:"is_triplet"`
	IsWatch       int         `json:"isWatch" xorm:"is_watch"`
}

func (utc UserTripletControls) TableName() string { return "gf_video_triplet" }

type UserWatchCountVideo struct {
	VideoNo      string      `json:"videoNo"`
	Duration     int         `json:"duration"`  // 传入的时长以秒为单位
}