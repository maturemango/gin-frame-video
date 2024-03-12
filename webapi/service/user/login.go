package user

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var login model.LoginMessage
	if err := c.BindJSON(&login); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}

	if code, err := conn.Get(fmt.Sprintf(model.LoginCode, login.Phone)); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	} else {
		if !strings.EqualFold(code, login.Code) {
			handlers.Base.Fail(c, 400, fmt.Errorf("code incorrect"))
			return
		}
	}

	if code, err := verfiyUserLogin(login); err != nil {
		if code == 400 {
			handlers.Base.Fail(c, code, err)
			return
		} else if code == 401 {
			handlers.Base.Fail(c, code, err)
			return
		}
	}

	var token string
	if info, err := getUserInfoByPhone(login.Phone); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	} else {
		if token, err = handlers.CreateToken(info); err != nil {
			handlers.Base.Fail(c, 400, err)
			return
		}
	}

	conn.Del(fmt.Sprintf(model.LoginCode, login.Phone))

	if err := handlers.AddSystemLog(handlers.Identity(), c.ClientIP(), model.Info, model.Login, model.LoginSuccess); err != nil {
		log.Printf("user:%v login addlog failed:%v", handlers.Identity(), err)
	}
	handlers.Base.OK(c, token)
}

func verfiyUserLogin(data model.LoginMessage) (int, error) {
	var pas string
	sql := `select password from gf_user where phone = ?`
	if _, err := conn.GetEngine().SQL(sql, data.Phone).Get(&pas); err != nil {
		return 400, err
	}
	if pas != data.Password {
		return 401, fmt.Errorf("password error")
	}
	return 200, nil
}

func getUserInfoByPhone(phone string) (model.UserInfo, error) {
	var info model.UserInfo
	sql := `select * from gf_user where phone = ? and role_id = 3`
	if _, err := conn.GetEngine().SQL(sql, phone).Get(&info); err != nil {
		return info, err
	}
	return info, nil
}

// 简易获取登录验证码的接口，具体实现需要短信模板以及通过短信模板将验证码发送到对应的手机号 限时一分钟
func GetLoginCode(c *gin.Context) {
	var data model.LoginMessage
	if err := c.BindJSON(&data); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}

	ok, err := regexp.MatchString("^1[3-9]\\d{9}$", data.Phone)
	if len(data.Phone) != 11 || !ok {
		handlers.Base.Fail(c, 400, fmt.Errorf("invalid phone"))
		return
	} else if err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	phoneKey := fmt.Sprintf(model.LoginCode, data.Phone)
	if conn.Exist(phoneKey) {
		handlers.Base.Fail(c, 400, fmt.Errorf("send logincode repeat"))
		return
	}
	code := utils.RandomCode(6)
	conn.Set(phoneKey, code, time.Minute * 1)
	handlers.Base.OK(c, code)
}
