package cmd

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/service/caption"
	"gin-frame/webapi/service/controls"
	sysLog "gin-frame/webapi/manage/log"
	"gin-frame/webapi/manage/role"
	"gin-frame/webapi/service/user"
	mamageVer "gin-frame/webapi/manage/version"
	"log"
	"strings"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func ginInit() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery(), GINLog(),
		verifyToken(), verifyPermission())

	addRouter(g)
	manageRouter(g.Group("/manageSystem"))

	return g
}

func GINLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		timestmp := end.Sub(start)
		path := c.Request.URL.Path
		clientIp := c.ClientIP()
		method := c.Request.Method
		code := c.Writer.Status()

		log.Printf("| %3d | %10v | %12s | %s  %s ",
			code,
			timestmp,
			clientIp,
			method, path)
	}
}

func addRouter(g *gin.Engine) {
	r := g.Group("/user")
	{
		r.POST("/login", user.UserLogin)                      // 用户登录
		r.POST("/register", user.UserRegister)                // 用户注册
		r.POST("/loginCode", user.GetLoginCode)               // 登录验证码(6位数)
		r.POST("/update/username", user.UpadteAccountName)    // 用户修改昵称
		r.POST("/upload/video", user.UploadUserVideo)         // 用户上传视频
		r.POST("/manage/videoList", user.GetUserVideoList)    // 获取用户视频列表
		r.GET("/manage/videoDetail", user.GetUserVideoDetail) // 获取用户视频信息
		r.GET("/manage/videoDel", user.DelUserVideo)          // 用户删除视频
	}
	r1 := g.Group("/caption")
	{
		r1.POST("/input", caption.InputWord) // 弹幕输入
		// r1.GET("/get/input", caption.GetInputWord)
		// r1.POST("/get/node/data", caption.GetNodeInputData)
		// r1.POST("/save/data", caption.SaveCacheInputData)
	}
	r2 := g.Group("/controls")
	{
		r2.POST("/watchVideo", controls.WatchAndCountVideo)    // 观看视频、统计观看次数
		r2.POST("/videoTriplet", controls.VideoTripletControl) // 视频三连相关操作
	}
}

func manageRouter(g *gin.RouterGroup) {
	g.POST("/version", mamageVer.GetSystemMessage) // 获取系统的版本相关信息
	logRouter(g.Group("/log"))
	roleRouter(g.Group("/role"))
}

func logRouter(g *gin.RouterGroup) {
	g.POST("/list", sysLog.GetLogList)  // 获取日志列表
}

func roleRouter(g *gin.RouterGroup) {
	g.POST("/list", role.GetRoleList)  // 获取角色列表
	g.POST("/update/status", role.UpdateRoleStatus)  // 修改角色状态
}

func verifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		claim, err := handlers.VerfiyToken(token)
		if err != nil {
			log.Printf("token invalid: %s\n", err)
			c.JSON(403, "token valid")
			c.Abort()
		}

		handlers.NewIdentity(*claim)
	}
}

type roleData struct {
	Sub      int    `xorm:"role_id"`
	Status   int    `xorm:"status"`
}

func verifyPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data roleData
		sql := `select gu.role_id, gr.status from 
		gf_user gu inner join gf_role gr on gu.role_id = gr.id 
		where gu.id = ?`
		if _, err := conn.GetEngine().SQL(sql, handlers.Identity()).Get(&data); err != nil {
			log.Printf("get user info failed: %s\n", err)
			c.JSON(403, "get user info failed")
			c.Abort()
		} else if data.Status != 1 {
			c.JSON(403, "role disable")
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
				c.JSON(403, "auth failed")
				c.Abort()
			} else if !ok {
				c.JSON(403, "user no permission")
				c.Abort()
			}
		}
	}
}