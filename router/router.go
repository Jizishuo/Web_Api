package router

import (
	"Web_Api/apis/system/dict"
	. "Web_Api/apis/tools"
	"Web_Api/handler/sd"
	"Web_Api/middleware"
	"github.com/gin-gonic/gin"
	"Web_Api/handler"
	"Web_Api/apis/monitor"
	"Web_Api/apis/system"
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

	r.POST("/login", authMiddleware.LoginHandler)
	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	r.GET("/routes", Dashboard)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("monitor/server", monitor.ServerInfo)

		apiv1.GET("/getCaptcha", system.GenerateCaptchaHandler)
		apiv1.GET("/db/tables/page", GetDBTableList)
		apiv1.GET("/db/columns/page", GetDBColumnList)
		apiv1.GET("/sys/tables/page", GetSysTableList)
		apiv1.POST("/sys/tables/info", InsertSysTable)
		apiv1.PUT("/sys/tables/info", UpdateSysTable)
		apiv1.DELETE("/sys/tables/info/:tableId", DeleteSysTables)
		apiv1.GET("/sys/tables/info/:tabledId", GetSysTables)
		apiv1.GET("/gen/preview/:tableId", Preview)
		apiv1.GET("/menuTreeselect", system.GetMenuTreeelect)
		apiv1.GET("/dict/databytype/:dictType", dict.GetDictDataByDictType)
	}


	return r
}

func Dashboard(c *gin.Context)  {
	var user = make(map[string]interface{})
	user["login_name"] = "amdin"
	user["user_id"] = 1
	user["user_name"] = "管理员"
	user["dept_id"] = 1

	var cmenuList = make(map[string]interface{})
	cmenuList["children"] = nil
	cmenuList["parent_id"] = 1
	cmenuList["title"] = "用户管理"
	cmenuList["name"] = "Sysuser"
	cmenuList["icon"] = "user"
	cmenuList["order_num"] = 1
	cmenuList["id"] = 4
	cmenuList["path"] = "sysuser"
	cmenuList["component"] = "sysuser/index"

	var lista = make([]interface{}, 1)
	lista[0] = cmenuList

	var menuList = make(map[string]interface{})
	menuList["children"] = lista
	menuList["parent_id"] = 1
	menuList["name"] = "Upms"
	menuList["title"] = "权限管理"
	menuList["icon"] = "example"
	menuList["order_num"] = 1
	menuList["id"] = 4
	menuList["path"] = "/upms"
	menuList["component"] = "Layout"

	var list = make([]interface{}, 1)
	list[0] = menuList
	var data = make(map[string]interface{})
	data["user"] = user
	data["menuList"] = list

	var r = make(map[string]interface{})
	r["code"] = 200
	r["data"] = data
	c.JSON(200, r)
}