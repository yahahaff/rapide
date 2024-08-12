package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/middlewares"
	"github.com/yahahaff/rapide/internal/routes/http/cloudflare"
)

func RouterGroup(router *gin.Engine) {
	httpRequest := router.Group("/api")
	httpRequest.Use(middlewares.AuthJWT(), middlewares.PermissionCheck(), middlewares.RecordOperation())
	EtcdRouter(httpRequest)
	cloudflare.DnsRouter(httpRequest)
	cloudflare.ZonesRouter(httpRequest)
	cloudflare.R2Router(httpRequest)

}
