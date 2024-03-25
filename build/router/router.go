package router

import (
	"gin-frame/webapi/service/caption"
	"gin-frame/webapi/service/controls"
	sysLog "gin-frame/webapi/manage/log"
	"gin-frame/webapi/manage/role"
	"gin-frame/webapi/service/user"
	mamageVer "gin-frame/webapi/manage/version"
	"gin-frame/webapi/manage/login"
	"github.com/gin-gonic/gin"
)

func AddRouter(g *gin.Engine) {
	r := g.Group("/user")
	{
		r.POST("/login", user.UserLogin)                      // 用户登录
		r.POST("/register", user.UserRegister)                // 用户注册
		r.POST("/loginCode", user.GetLoginCode)               // 登录验证码(6位数)
		r.POST("/update/username", user.UpadteAccountName)    // 用户修改昵称
		r.POST("/update/psw", user.UpdateUserPsw)             // 用户修改密码
		r.POST("/update/send", user.SendPswCode)              // 修改密码验证码
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

func ManageRouter(g *gin.RouterGroup) {
	g.POST("/version", mamageVer.GetSystemMessage) // 获取系统的版本相关信息
	loginRouter(g)
	logRouter(g.Group("/log"))
	roleRouter(g.Group("/role"))
}

func logRouter(g *gin.RouterGroup) {
	g.POST("/list", sysLog.GetLogList) // 获取日志列表
}

func roleRouter(g *gin.RouterGroup) {
	g.POST("/list", role.GetRoleList)               // 获取角色列表
	g.POST("/update/status", role.UpdateRoleStatus) // 修改角色状态
}

func loginRouter(g *gin.RouterGroup) {
	g.POST("/login", login.ManageSysLogin)      // 管理员登录控制台
	g.POST("/login/captcha", login.SysLoginCode)    // 控制台登录验证码（锁+map）
	// g.POST("/login/captcha1", login.SysLoginCode1)  // 控制台登录验证码（sync.Map）
	g.POST("/loginOut", login.SysLoginOut)  // 控制台登出清除一些必要数据
}