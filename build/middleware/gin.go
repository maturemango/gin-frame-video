package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

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