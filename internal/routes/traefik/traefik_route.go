package traefik

import (
	traefikCtl "rapide/internal/controllers/traefik"

	"github.com/gin-gonic/gin"
)

// TraefikRouter 注册Traefik相关路由
func TraefikRouter(routerGroup *gin.RouterGroup) {
	traefikController := traefikCtl.TraefikController{}

	// 获取HTTPRoutes列表路由
	// /api/traefik/http-routes?page=1&pageSize=20
	routerGroup.GET("/http-routes", traefikController.GetHTTPRoutes)

	// 创建HTTPRoute路由
	// /api/traefik/http-routes
	routerGroup.POST("/http-routes", traefikController.CreateHTTPRoute)

	// 获取Services列表路由
	// /api/traefik/services?page=1&pageSize=10
	routerGroup.GET("/services", traefikController.GetServices)

	// 创建Service路由
	// /api/traefik/services
	routerGroup.POST("/services", traefikController.CreateService)

	// 获取Middlewares列表路由
	// /api/traefik/middlewares?page=1&pageSize=10
	routerGroup.GET("/middlewares", traefikController.GetMiddlewares)

	// 创建Middleware路由
	// /api/traefik/middlewares
	routerGroup.POST("/middlewares", traefikController.CreateMiddleware)

}
