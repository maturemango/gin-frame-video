package log

import (
	"gin-frame/build/conn"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"

	"github.com/gin-gonic/gin"
)

func GetLogList(c *gin.Context) {
	var list []model.LogList
	sql := `select gu.phone, gl.addr, gl.log_level, date_format(gl.operate_time, '%Y-%m-%d %H:%i:%s') operate_time, gl.operate_desc, gl.detail from 
	gf_log gl inner join gf_user gu on gl.user_id = gu.id order by gl.operate_time desc`
	if err := conn.GetEngine().SQL(sql).Find(&list); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	handlers.Base.OK(c, list)
}
