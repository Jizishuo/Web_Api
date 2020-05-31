package middleware

import (
	config2 "Web_Api/config"
	jwt "Web_Api/pkg/jwtauth"
	"go-admin/handler"
	"time"
)

func AuthInit() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm: "test zone",
		Key:[]byte("secret key"),
		Timeout: time.Hour,
		MaxRefresh: time.Hour,
		IdentityKey: config2.ApplicationConfig.JwtSecret,
		PayloadFunc: handler.PayloadFunc,
	})
}