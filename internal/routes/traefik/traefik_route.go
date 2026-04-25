package traefik

import (
	traefikCtl "rapide/internal/controllers/traefik"

	"github.com/gin-gonic/gin"
)

// TraefikRouter 注册Traefik相关路由
func TraefikRouter(routerGroup *gin.RouterGroup) {
	traefikController := traefikCtl.TraefikController{}

	// HTTPRoute 相关路由
	// 获取HTTPRoutes列表路由
	// /api/traefik/http-routes?page=1&pageSize=20
	routerGroup.GET("/http-routes", traefikController.GetHTTPRoutes)
	// 获取单个HTTPRoute详情
	// /api/traefik/http-routes/:name?namespace=traefik
	routerGroup.GET("/http-routes/:name", traefikController.GetHTTPRoute)
	// 创建HTTPRoute路由
	// /api/traefik/http-routes
	routerGroup.POST("/http-routes", traefikController.CreateHTTPRoute)
	// 更新HTTPRoute路由
	// /api/traefik/http-routes/:name?namespace=traefik
	routerGroup.PUT("/http-routes/:name", traefikController.UpdateHTTPRoute)
	// 删除HTTPRoute路由
	// /api/traefik/http-routes/:name?namespace=traefik
	routerGroup.DELETE("/http-routes/:name", traefikController.DeleteHTTPRoute)

	// Service 相关路由
	// 获取Services列表路由
	// /api/traefik/services?page=1&pageSize=10
	routerGroup.GET("/services", traefikController.GetServices)
	// 获取单个Service详情
	// /api/traefik/services/:name?namespace=traefik
	routerGroup.GET("/services/:name", traefikController.GetService)
	// 创建Service路由
	// /api/traefik/services
	routerGroup.POST("/services", traefikController.CreateService)
	// 更新Service路由
	// /api/traefik/services/:name?namespace=traefik
	routerGroup.PUT("/services/:name", traefikController.UpdateService)
	// 删除Service路由
	// /api/traefik/services/:name?namespace=traefik
	routerGroup.DELETE("/services/:name", traefikController.DeleteService)

	// Middleware 相关路由
	// 获取Middlewares列表路由
	// /api/traefik/middlewares?page=1&pageSize=10
	routerGroup.GET("/middlewares", traefikController.GetMiddlewares)
	// 获取单个Middleware详情
	// /api/traefik/middlewares/:name?namespace=traefik
	routerGroup.GET("/middlewares/:name", traefikController.GetMiddleware)
	// 创建Middleware路由
	// /api/traefik/middlewares
	routerGroup.POST("/middlewares", traefikController.CreateMiddleware)
	// 更新Middleware路由
	// /api/traefik/middlewares/:name?namespace=traefik
	routerGroup.PUT("/middlewares/:name", traefikController.UpdateMiddleware)
	// 删除Middleware路由
	// /api/traefik/middlewares/:name?namespace=traefik
	routerGroup.DELETE("/middlewares/:name", traefikController.DeleteMiddleware)

}
