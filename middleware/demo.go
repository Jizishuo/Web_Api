package middleware

import (
	"Web_Api/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DemoEnv() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.ApplicationConfig.Env == "demo" {
			method := c.Request.Method
			if method == "GET" || method == "OPTIONS" || c.Request.RequestURI == "/login" || c.Request.RequestURI == "/api/v1/logout" {
				c.Next()
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  config.ApplicationConfig.EnvMsg,
				})
				c.Abort()
				return
			}
		} else {
			c.Next()
		}
	}
}