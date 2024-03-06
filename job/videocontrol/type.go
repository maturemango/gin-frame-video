package videocontrol

type VideoTripletData struct {
	VideoNo      string     `xorm:"video_no"`
	Upvote       int        `xorm:"upvote"`
	Disagree     int        `xorm:"disagree"`
	Coins        int        `xorm:"coins"`
	Collect      int        `xorm:"collect"`
	Watch        int        `xorm:"watch"`
}