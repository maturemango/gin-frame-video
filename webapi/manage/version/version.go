package version

import (
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"

	"github.com/gin-gonic/gin"
)

// 动态设置并且获取系统版本信息 
//在程序编译或者运行时使用 -ldflags "-X main.Version=v1.0.0 -X main.GoVersion=1.20" 
func GetSystemMessage(c *gin.Context) {
	handlers.Base.OK(c, model.Version)
}