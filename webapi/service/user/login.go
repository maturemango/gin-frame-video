package user

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"
	"gin-frame/webapi/service"
	"log"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var login model.LoginMessage
	if err := c.BindJSON(&login); err != nil {
		service.Svc.Fail(c, 400, err)
		return
	}
	
	if code, err := verfiyUserLogin(login); err != nil {
		if code == 400 {
			service.Svc.Fail(c, code, err)
			return
		} else if code == 401 {
			service.Svc.Fail(c, code, err)
			return
		}
	}
	if err := handlers.AddSystemLog(handlers.Identity(), c.ClientIP(), model.Info, model.Login, model.LoginSuccess); err != nil {
		log.Printf("user:%v login addlog failed:%v", handlers.Identity(), err)
	}
	service.Svc.OK(c, "login success")
}

func verfiyUserLogin(data model.LoginMessage) (int, error) {
	var pas string
	sql := `select password from gf_user where account = ?`
	if _, err := conn.GetEngine().SQL(sql, data.Account).Get(&pas); err != nil {
		return 400, err
	}
	if pas != data.Password {
		return 401, fmt.Errorf("password error")
	}
	return 200, nil
}