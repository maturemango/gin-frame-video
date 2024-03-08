package user

import (
	"gin-frame/build/conn"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"

	"github.com/gin-gonic/gin"
)

func GetUserVideoList(c *gin.Context) {
	var data []model.UserVideoList
	sql := `select video_no, title from gf_video where user_id = ? and is_del = 0`
	if err := conn.GetEngine().SQL(sql, handlers.Identity()).Find(&data); err != nil {
		handlers.Base.Fail(c, 500, err)
		return
	}
	handlers.Base.OK(c, data)
}

func GetUserVideoDetail(c *gin.Context) {
	var detail model.UserVideoDetail
	k := c.Query("no")
	sql := `select video_no, title, introduction, date_format(upload_time, '%Y-%m-%d %H:%i:%s') upload_time, upvote, disagree, coins, collect 
	from gf_video where video_no = ?`
	if _, err := conn.GetEngine().SQL(sql, k).Get(&detail); err != nil {
		handlers.Base.Fail(c, 500, err)
		return
	}
	handlers.Base.OK(c, detail)
}

func DelUserVideo(c *gin.Context) {
	var detail model.UserVideoDetail
	detail.IsDel = 1
	k := c.Query("no")
	if _, err := conn.GetEngine().Cols("is_del").Where("video_no = ?", k).Update(&detail); err != nil {
		handlers.Base.Fail(c, 500, err)
		return
	}
	handlers.Base.OK(c, "delete video success")
}
