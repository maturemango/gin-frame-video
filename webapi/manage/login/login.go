package login

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"

	"github.com/gin-gonic/gin"
)

func ManageSysLogin(c *gin.Context) {
	var (
		data   model.LoginMessage
		usr    model.UserInfo
	)
	if err := c.BindJSON(&data); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	sql := `select %s from gf_user where phone = ? and password = ? and role_id <= ?`
	pas, err := handlers.EncodeCrypto(data.Password)
	if err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	if count, err := conn.GetEngine().SQL(fmt.Sprintf(sql, "count(1)"), data.Phone, pas, utils.Config.Manage.RoleId).Count(); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	} else if count > 1 {
		handlers.Base.Fail(c, 401, fmt.Errorf("user repeated error"))
		return
	} else {
		if _, err := conn.GetEngine().SQL(fmt.Sprintf(sql, "*"), data.Phone, pas, utils.Config.Manage.RoleId).Get(&usr); err != nil {
			handlers.Base.Fail(c, 401, err)
			return
		}
	}
	token, err := handlers.CreateToken(usr)
	if err != nil {
		handlers.Base.Fail(c, 401, err)
		return
	}
	handlers.Base.OK(c, token)
}
