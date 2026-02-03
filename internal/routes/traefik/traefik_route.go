package traefik

import (
	"github.com/gin-gonic/gin"
	traefikCtl "rapide/internal/controllers/traefik"
)

// TraefikRouter 注册Traefik相关路由
func TraefikRouter(routerGroup *gin.RouterGroup) {
	traefikController := traefikCtl.TraefikController{}

	// 获取特定CRD下的资源列表路由
	// /api/traefik/crds-resources?group=gateway.networking.k8s.io&version=v1&kind=HTTPRoute
	routerGroup.GET("/crds-resources", traefikController.GetCRResources)
}
