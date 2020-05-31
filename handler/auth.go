package handler

import (
	jwt "Web_Api/pkg/jwtauth"
	"github.com/mojocn/base64Captcha"
	"Web_Api/models"
)

var store = base64Captcha.DefaultMemStore

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{});ok {
		u, _ := v["user"].(models.SysUser)
	}
}