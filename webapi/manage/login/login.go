package login

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"
	"image/color"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var (
	store = base64Captcha.DefaultMemStore
	captchaid string
	mu sync.RWMutex
	requests = make(map[string]*requestInfo)
)

func ManageSysLogin(c *gin.Context) {
	var (
		data model.LoginMessage
		usr  model.UserInfo
	)
	if err := c.BindJSON(&data); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	if err := verfiyCaptcha(data.Code); err != nil {
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

func verfiyCaptcha(answer string) error {
	if ok := store.Verify(captchaid, answer, true); !ok {
		return fmt.Errorf("captcha invaild")
	} else {
		return nil
	}
}

// 该接口的answer在实际使用时不并不需要
func SysLoginCode(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	ip := c.ClientIP()
	info, exist := requests[ip]
	if !exist || time.Since(info.fristRequest) > time.Duration(utils.Config.Manage.TimeLimit)*time.Minute {
		info = &requestInfo{
			count:          1,
			fristRequest:   time.Now(),
		}
		requests[ip] = info
	} else {
		info.count++
		if info.count > utils.Config.Manage.Bucket {
			handlers.Base.Fail(c, 429, fmt.Errorf("refresh too many"))
			return
		}
	}
	captcha := base64Captcha.NewCaptcha(base64Captcha.NewDriverString(60, 200, 2, 5, utils.Config.Manage.CodeLength, model.Characters, &color.RGBA{0, 0, 0, 0}, 
		base64Captcha.NewEmbeddedFontsStorage(model.FontFS), []string{"simhei.ttf"}), store)
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		handlers.Base.Fail(c, 400, fmt.Errorf("generate captha failed"))
		return
	}
	var resp SysLoginCaptcha
	captchaid = id
	resp.Id = id
	resp.B64s = b64s
	resp.Answer = answer
	handlers.Base.OK(c, resp)
}
