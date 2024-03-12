package user

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"
	"log"
	"strings"
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

func UpdateUserPsw(c *gin.Context) {
	var info model.UpdateLoginPsw
	if err := c.BindJSON(&info); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	} else if len(info.OldPssword) <= 0 || len(info.NewPssword) <= 0 {
		handlers.Base.Fail(c, 400, fmt.Errorf("password not null"))
		return
	}

	oldPsw, _ := handlers.EncodeCrypto(info.OldPssword)
	usr, err := handlers.GetUserInfoById(handlers.Identity())
	if err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	} else if !strings.EqualFold(usr.Password, oldPsw) {
		handlers.Base.Fail(c, 400, fmt.Errorf("oldPassword incorrect"))
		return
	}
	phoneKey := fmt.Sprintf(model.ModfiyPswCode, usr.Phone)
	code, _ := conn.Get(phoneKey)
	if !strings.EqualFold(code, info.Code) {
		handlers.Base.Fail(c, 400, fmt.Errorf("code incorrect"))
		return
	}
	usr.Password, _ = handlers.EncodeCrypto(info.NewPssword)
	if _, err := conn.GetEngine().Where("id = ?", handlers.Identity()).Cols("password").Update(&usr); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	conn.Del(phoneKey)
	if err := handlers.AddSystemLog(handlers.Identity(), c.ClientIP(), model.Info, model.UpdateUserPsw, model.UpdatePswSuccess); err != nil {
		log.Printf("user:%v update psw addlog failed:%v", handlers.Identity(), err)
	}
	handlers.Base.OK(c, "update password success")
}

// 通接口/user/loginCode一样
func SendPswCode(c *gin.Context) {
	usr, err := handlers.GetUserInfoById(handlers.Identity())
	if err != nil {
		handlers.Base.Fail(c, 400, fmt.Errorf(err.Error()))
		return
	}
	
	code := utils.RandomCode(6)
	phoneKey := fmt.Sprintf(model.ModfiyPswCode, usr.Phone)
	if conn.Exist(phoneKey) {
		handlers.Base.Fail(c, 400, fmt.Errorf("send logincode repeat"))
		return
	}
	conn.Set(phoneKey, code, 1 * time.Minute)
	handlers.Base.OK(c, code)
}