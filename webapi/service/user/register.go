package user

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/webapi/model"
	"gin-frame/webapi/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var r model.RegisterData
	if err := c.BindJSON(&r); err != nil {
		service.Svc.Fail(c, 400, err)
		return
	}
	err := verfiyRegisterData(r)
	if err != nil {
		service.Svc.Fail(c, 400, err)
		return
	}
	r.RoleId = model.User
	if _, err := conn.GetEngine().InsertOne(&r); err != nil {
		service.Svc.Fail(c, 400, err)
		return
	}
	service.Svc.OK(c, "register success")
}

func verfiyRegisterData(d model.RegisterData) error {
	if len(d.Account) <= 0 || len(d.Password) <= 0 {
		return fmt.Errorf("account or password not null")
	}
	return nil
}