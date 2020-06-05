package middleware

import (
	"Web_Api/pkg"
	"Web_Api/pkg/jwtauth"
	"github.com/gin-gonic/gin"
	mycasbin "Web_Api/pkg/casbin"
	"log"
	"net/http"
)

//权限检查中间件
func AuthChecRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("JWT_PAYLOAD")
		v := data.(jwtauth.MapClaims)
		e, err := mycasbin.Casbin()
		pkg.HasError(err, "", 500)
		// 检查权限
		res, err := e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)

		log.Println("----------------", v["rolekey"], c.Request.URL.Path, c.Request.Method)

		pkg.HasError(err,"",500)
		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 401,
				"msg":    "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}
	}
}