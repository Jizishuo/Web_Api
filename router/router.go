package router

import (
	"Web_Api/handler/sd"
	"Web_Api/middleware"
	"github.com/gin-gonic/gin"
	"Web_Api/handler"
	"log"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.LoggerToFile())
	r.Use(middleware.CustomError)
	r.Use(middleware.NoCache)
	r.Use(middleware.Options)
	r.Use(middleware.Secure)
	r.Use(middleware.RequestId())
	r.Use(middleware.DemoEnv())

	r.Static("/static", "./stctic")
	r.GET("/info", handler.Ping)

	// 监控信息
	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
		svcd.GET("/os", sd.OSCheck)
	}
	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	if err!=nil {
		log.Fatalln("JWT Error", err.Error())
	}
	// Refresh time can be longer than token timeout

}