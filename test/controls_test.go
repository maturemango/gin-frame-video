package test

import (
	"fmt"
	"gin-frame/build/cmd"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/model"
	"testing"
)

func TestVideoControl(t *testing.T) {
	cmd.LoadConfig()
	utils.InitConfig()

	var data model.UserVideoDetail
	no := "vn1709002620972380100o1ZADEvOPC"
	sql := `select count(*) from gf_video_triplet where video_no = ? and is_upvote = 1`
	if count, err := conn.GetEngine().SQL(sql, no).Count(); err != nil {
		t.Logf("count sql num failed:%v", err)
	} else {
		sql := `select * from gf_video where video_no = ?`
		if _, err := conn.GetEngine().SQL(sql, no).Get(&data); err != nil {
			t.Fatalf("query video failed:%v", err)
		}
		if count != int64(data.Upvote) {
			data.Upvote = int(count)
			if _, err := conn.GetEngine().Cols("upvote").Where("video_no = ?", no).Update(&data); err != nil {
				t.Fatalf("update video failed:%v", err)
			}
			fmt.Println("update video quvote num success")
		}
	}
}