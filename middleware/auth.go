package middleware

import (
	config2 "Web_Api/config"
	jwt "Web_Api/pkg/jwtauth"
	"Web_Api/handler"
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
		IdentityHandler: handler.IdentityHandler,
		Authenticator:   handler.Authenticator,
		Authorizator:    handler.Authorizator,
		Unauthorized:    handler.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}