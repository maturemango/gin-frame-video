package user

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var r model.RegisterData
	if err := c.BindJSON(&r); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	err := verfiyRegisterData(r)
	if err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	r.RoleId = model.User
	if _, err := conn.GetEngine().InsertOne(&r); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	handlers.Base.OK(c, "register success")
}

func verfiyRegisterData(d model.RegisterData) error {
	if len(d.Phone) <= 0 || len(d.Password) <= 0 {
		return fmt.Errorf("account or password not null")
	}
	return nil
}
