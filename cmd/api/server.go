package api

import (
	"Web_Api/database"
	"Web_Api/global/orm"
	"Web_Api/router"
	"Web_Api/tools"
	config2 "Web_Api/tools/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	config string
	port string
	mode string
	StartCmd = &cobra.Command{
		Use: "server",
		Short: "start api server",
		Example: "Web_Api server config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func usage()  {
	usageStr := `starting api server`
	log.Printf("%s\n", usageStr)
}

func setup()  {
	//读取配置
	config2.ConfigSetup(config)
	//读取日志
	tools.InitLogger()
	//初始化数据库连接
	database.Setup()
}

func run() error {
	if mode != "" {
		config2.SetConfig(config, "settings.application.mode", mode)
	}
	if viper.GetString("settings.application.mode") == string(tools.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.InitRouter()

	defer orm.Eloquent.Close()
	if port != "" {
		config2.SetConfig(config, "settings.application.port", port)
	}
	srv := &http.Server{
		Addr:    config2.ApplicationConfig.Host + ":" + config2.ApplicationConfig.Port,
		Handler: r,
	}
	go func() {
		// 服务连接
		if config2.ApplicationConfig.IsHttps {
			if err := srv.ListenAndServeTLS(config2.SslConfig.Pem, config2.SslConfig.KeyStr); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}
	}()
	content, _ := ioutil.ReadFile("./static/api.txt")
	fmt.Println(string(content))
	//fmt.Printf("%s Server Run http://127.0.0.1:%s/ \r\n", tools.GetCurrntTimeStr(), config2.ApplicationConfig.Port)
	//fmt.Printf("%s Swagger URL http://127.0.0.1:%s/swagger/index.html \r\n", tools.GetCurrntTimeStr(), config2.ApplicationConfig.Port)
	//fmt.Printf("%s Enter Control + C Shutdown Server \r\n", tools.GetCurrntTimeStr())
	//// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//fmt.Printf("%s Shutdown Server ... \r\n", tools.GetCurrntTimeStr())
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//log.Println("Server exiting")
	return nil
}