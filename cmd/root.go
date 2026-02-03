// Package cmd /*
package cmd

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"rapide/initialize"
	"rapide/pkg/config"
	"rapide/pkg/console"
	"rapide/pkg/logger"
)

func Execute() {
	// 1.初始化viper 以获取env环境变量
	config.InitConfig()

	// gin 实例
	gin.SetMode(config.GetString("APP_ENV", "debug")) // debug,test,release
	router := gin.New()

	// 初始化路由绑定
	initialize.SetupRoute(router)

	// 初始化 Logger
	initialize.SetupLogger()

	// 初始化数据库
	initialize.SetupDB()

	// 初始化Kubernetes客户端
	initialize.SetupKubernetes()

	// 初始化Validator
	initialize.SetupValidators()

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    ":" + config.GetString("APP_PORT", "8000"),
		Handler: router,
	}

	// 启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorString("gin", "serve", err.Error())
			console.Exit("Unable to start server, error:" + err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.InfoString("gin", "shutdown", "Shutting down server...")

	// 创建3秒的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.ErrorString("gin", "shutdown", "Server forced to shutdown: "+err.Error())
	}

	logger.InfoString("gin", "shutdown", "Server exiting")
}
