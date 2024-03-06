package sysLog

import (
	"gin-frame/build/conn"
	"gin-frame/webapi/model"
	"gin-frame/webapi/service"

	"github.com/gin-gonic/gin"
)

// 暂时不设置用户管理员，后续补上
func GetLogList(c *gin.Context) {
	var list []model.LogList
	sql := `select gu.account, gl.addr, gl.log_level, gl.operate_time, gl.operate_desc, gl.detail from 
	gf_log gl inner join gf_user gu on gl.user_id = gu.id order by gl.operate_time desc`
	if err := conn.GetEngine().SQL(sql).Find(&list); err != nil {
		service.Svc.Fail(c, 400, err)
		return
	}
	service.Svc.OK(c, list)
}

