package main

import (
	"Web_Api/config"
	"Web_Api/router"
	"context"
	"github.com/gin-gonic/gin"
	"go-admin/models/gorm"
	"log"
	orm "Web_Api/database"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main()  {
	gin.SetMode(gin.DebugMode)
	log.Println(config.DatabaseConfig.Port)

	err := gorm.AutoMigrate(orm.Eloquent)
	if err != nil {
		log.Fatal("数据库基础数据初始化失败")
	} else {
		config.SetApplicationIsInit()
	}

	r := router.InitRouter()


	defer orm.Eloquent.Close()

	srv := &http.Server{
		Addr: config.ApplicationConfig.Host+":"+ config.ApplicationConfig.Port,
		Handler: r,
	}
	go func() {
		//服务连接
		if err := srv.ListenAndServe();err!=nil&&err!=http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server Run ", config.ApplicationConfig.Host+":"+ config.ApplicationConfig.Port)
	log.Println("Enter Control + C Shutdown Server")
	// 等待中断信号  优雅的关闭服务器 超时5秒
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shoutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx);err!=nil{
		log.Fatal("Server Shutdown", err)
	}
	log.Println("Server exiting")
}
