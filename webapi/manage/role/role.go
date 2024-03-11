package role

import (
	"gin-frame/build/conn"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"

	"github.com/gin-gonic/gin"
)

func GetRoleList(c *gin.Context) {
	var list []model.RoleList
	if err := conn.GetEngine().Find(&list); err != nil {
		handlers.Base.Fail(c, 500, err)
		return
	}
	handlers.Base.OK(c, list)
}

func UpdateRoleStatus(c *gin.Context) {
	var role model.RoleList
	if err := c.BindJSON(&role); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}

	if _, err := conn.GetEngine().Cols("status").Where("id = ?", role.RoleId).Update(&role); err != nil {
		handlers.Base.Fail(c, 500, err)
		return
	}
	handlers.Base.OK(c, "update role status success")
}
