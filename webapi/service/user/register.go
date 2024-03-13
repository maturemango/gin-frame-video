package user

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"
	"strings"

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
	r.UserName = generateRandomName(r.Phone)
	if r.Password, err = handlers.EncodeCrypto(r.Password); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	_, err = conn.GetEngine().InsertOne(&r)
	if err == nil {
		handlers.Base.OK(c, "register success")
	} else if strings.Contains(err.Error(), "gf_user.phone_role") {
		handlers.Base.Fail(c, 400, fmt.Errorf("user exist"))
		return
	} else if err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
}

func verfiyRegisterData(d model.RegisterData) error {
	if len(d.Phone) <= 0 || len(d.Password) <= 0 {
		return fmt.Errorf("account or password not null")
	}
	return nil
}

func generateRandomName(phone string) (random string) {
	random = utils.Config.Login.No + utils.RandomNo(4) + phone
	return random
}