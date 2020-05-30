package middleware

import (
	"Web_Api/config"
	"Web_Api/models"
	"Web_Api/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// 实例化

var logger = logrus.New()

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	// 写入文件
	src, err := os.OpenFile(config.ApplicationConfig.LogPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err:", err)
	}

	// 设置输出
	logger.Out = src
	
	// 设置日志级别
	logger.SetFormatter(&logrus.TextFormatter{})
	
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		//执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUrl := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求id
		clientIP := c.ClientIP()

		// 日志格式
		logger.Infof("%3d  %13v  %15s  %s  %s\r\n",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUrl,
			)
		logger.Infof("-")
		fmt.Println(
			startTime.Format("\n2006-01-02 15:04:05.9999"),
			"[INFO]",
			reqMethod,
			reqUrl,
			statusCode,
			latencyTime,
			reqUrl,
			clientIP,
			)

		if c.Request.Method != "GET" && c.Request.Method != "OPTIONS" {
			menu := models.Menu{}
			menu.Path = reqUrl
			menu.Action = reqMethod
			menuList, _ := menu.Get()
			sysOperLog := models.SysOperLog{}
			sysOperLog.OperIp = clientIP
			sysOperLog.OperLocation = utils.GetLocation(clientIP)
			sysOperLog.Status = utils.IntToString(statusCode)
			sysOperLog.OperName = utils.GetRoleName(c)
			sysOperLog.RequestMethod = c.Request.Method
			sysOperLog.OperUrl = reqUrl
			if reqUrl == "/login" {
				sysOperLog.BusinessType = "10"
				sysOperLog.Title = "用户登录"
				sysOperLog.OperName = "-"
			}

		}

	}
}