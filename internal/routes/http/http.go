package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func RouterGroup(router *gin.Engine) *gin.RouterGroup {
	httpRequest := router.Group("/api")
	httpRequest.Use(middlewares.AuthJWT(), middlewares.PermissionCheck())
	EtcdRouter(httpRequest)
	DnsRouter(httpRequest)
	ZonesRouter(httpRequest)
	return httpRequest
}
