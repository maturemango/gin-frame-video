package user

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"
	"time"

	"github.com/gin-gonic/gin"
)

func UpadteAccountName(c *gin.Context) {
	var usr model.UserInfo
	var cols []string
	if err := c.BindJSON(&usr); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}

	if len(usr.UserName) <= 0 {
		handlers.Base.Fail(c, 400, fmt.Errorf("username not null"))
		return
	} else {
		cols = append(cols, "user_name")
	}
	i, err := VerfiyAccountName(usr.UserName)
	if err != nil {
		handlers.Base.Fail(c, 500, fmt.Errorf("service error"))
		return
	} else if i == 0 {
		usr.UpdatedTime = time.Now()
		cols = append(cols, "updated_time")
	}
	if _, err := conn.GetEngine().Where("id = ?", handlers.Identity()).Cols(cols...).Update(&usr); err != nil {
		handlers.Base.Fail(c, 500, fmt.Errorf("service error"))
		return
	}

	handlers.Base.OK(c, "update name success")
}

// 验证修改的用户名是否有变化
func VerfiyAccountName(name string) (int, error) {
	var username string
	sql := "select user_name from gf_user where id = ?"
	if _, err := conn.GetEngine().SQL(sql, handlers.Identity()).Get(&username); err != nil {
		return 0, err
	}
	if username == name {
		return 0, nil
	} else {
		return 1, nil
	}
}
