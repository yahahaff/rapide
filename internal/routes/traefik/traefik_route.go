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

}
