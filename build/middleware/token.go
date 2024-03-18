package middleware

import (
	"gin-frame/webapi/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		claim, err := handlers.VerfiyToken(token)
		if err != nil {
			log.Printf("token invalid: %s\n", err)
			c.JSON(403, gin.H{
				"code": 403,
				"message": "token valid",
			})
			c.Abort()
		}

		handlers.NewIdentity(*claim)
	}
}