package cloudflare

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func RouterGroup(router *gin.Engine) *gin.RouterGroup {
	cloudflare := router.Group("/api/cloudflare")
	cloudflare.Use(middlewares.AuthJWT(), middlewares.PermissionCheck())
	ZonesRouter(cloudflare)
	DnsRouter(cloudflare)
	return cloudflare
}
