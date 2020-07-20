package api

import (
	config2 "Web_Api/tools/config"
	"github.com/spf13/cobra"
	"log"
)

var (
	config string
	post string
	mode string
	StartCmd = &cobra.Command{
		Use: "server",
		Short: "start api server",
		Example: "Web_Api server config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage
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
	tools.in
	//初始化数据库连接

}