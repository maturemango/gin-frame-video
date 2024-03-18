package cmd

import (
	"gin-frame/build/middleware"
	"gin-frame/build/router"
	"github.com/gin-gonic/gin"
)

func init() {
	middleWares = append(middleWares, gin.Recovery(),
		middleware.GINLog(), middleware.VerifyToken(), 
		middleware.VerifyPermission())
}

var middleWares []gin.HandlerFunc

func ginInit() *gin.Engine {
	g := gin.New()
	g.Use(middleWares...)

	router.AddRouter(g)
	router.ManageRouter(g.Group("/manageSystem"))

	return g
}
