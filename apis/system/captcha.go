package system

import (
	"Web_Api/pkg"
	"Web_Api/pkg/app"
	"Web_Api/pkg/captcha"
	"github.com/gin-gonic/gin"
)

func GenerateCaptchaHandler(c *gin.Context)  {
	id, b64s, err := captcha.DriverDigitFunc()
	pkg.HasError(err, "验证码获取失败", 500)
	app.Custum(c, gin.H{
		"code":200,
		"date":b64s,
		"id":id,
		"msg":"success",
	})
}