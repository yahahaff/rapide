// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yahahaff/rapide/initialize"
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/console"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "Rapide",
	Short: "A simple forum project",
	Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

	// rootCmd 的所有子命令都会执行以下代码
	PersistentPreRun: func(command *cobra.Command, args []string) {

		// 配置文件初始化，依赖命令行 --config 参数
		config.InitConfig(ConfigCmd)

		// 初始化 Logger
		initialize.SetupLogger()

		// 初始化数据库
		initialize.SetupDB()

		// 初始化 Redis
		initialize.SetupRedis()

		// 初始化casbin
		initialize.SetupCasbinEnforcer()

		// 初始化Validator
		initialize.SetupValidators()

		// 初始化EtcdClient
		//initialize.SetupEtcd()

	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(CmdKey)
	rootCmd.AddCommand(CmdServe) //注册ServeCommand

	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")

	// 配置默认运行 Web 服务
	RegisterDefaultCmd(rootCmd, CmdServe)

	// 注册全局参数
	RegisterGlobalFlags(rootCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		console.Exit(fmt.Sprintf("Failed to run internal with %v: %s", os.Args, err.Error()))
	}
}
