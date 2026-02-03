// Package initialize 处理程序初始化逻辑
package initialize

import (
	"net/http"
	"strings"

	"rapide/internal/middlewares"
	"rapide/internal/routes"
	"rapide/pkg/console"
	"rapide/pkg/response"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {
	// 存储validator验证器对象到Context头中
	router.Use(func(c *gin.Context) {
		c.Set("validator", Validate)
		c.Next()
	})

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	routes.RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)

	// 打印log
	printLogo()
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.RecordOperation(),
		//middlewares.ForceUA(),
	)
}

func printLogo() {
	// ASCII艺术字
	const logo = "  _ __ __ _ _ __ (_) __| | ___ \n | '__/ _` | '_ \\| |/ _` |/ _ \\\n | | | (_| | |_) | | (_| |  __/\n |_|  \\__,_| .__/|_|\\__,_|\\___|\n           |_|                 "

	console.Error(logo)
	console.Success(":: Rpaide :: (v1.0.0 Release)")
	console.Success("Startup succeeded")
}
func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "404 not found")
		} else {
			// 默认返回 JSON
			response.Abort404(c, "404 not found")
		}
	})
}
