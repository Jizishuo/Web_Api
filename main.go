package main

import (
	"Web_Api/config"
	"Web_Api/router"
	"github.com/gin-gonic/gin"
	"go-admin/models/gorm"
	"log"
	orm "Web_Api/database"
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
}
