package middleware

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/handlers"
	"log"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type roleData struct {
	Sub      int    `xorm:"role_id"`
	Status   int    `xorm:"status"`
}

func VerifyPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.URL.Path
		if strings.Contains(method, "/login") {
			return
		}
		var data roleData
		sql := `select gu.role_id, gr.status from 
		gf_user gu inner join gf_role gr on gu.role_id = gr.id 
		where gu.id = ?`
		if _, err := conn.GetEngine().SQL(sql, handlers.Identity()).Get(&data); err != nil {
			log.Printf("get user info failed: %s\n", err)
			c.JSON(403, gin.H{
				"code": 403,
				"message": "get user info failed",
			})
			c.Abort()
		} else if data.Status != 1 {
			c.JSON(403, gin.H{
				"code": 403,
				"message": "role disable",
			})
			c.Abort()
		}
		obj := c.Request.URL.Path
		act := c.Request.Method

		if strings.Contains(obj, "manageSystem") {
			e, err := casbin.NewEnforcer(utils.Config.Casbin.ModelPath, utils.Config.Casbin.PolicyPath)
			if err != nil {
				log.Printf("failed to create enforcer: %s\n", err)
			}
			if ok, err := e.Enforce(fmt.Sprintf("%d", data.Sub), obj, act); err != nil {
				log.Printf("enforce failed: %s\n", err)
				c.JSON(403, gin.H{
					"code": 403,
					"message": "auth failed",
				})
				c.Abort()
			} else if !ok {
				c.JSON(403, gin.H{
					"code": 403,
					"message": "user no perssion",
				})
				c.Abort()
			}
		}
	}
}