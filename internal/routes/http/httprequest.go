package http

import (
	"github.com/gin-gonic/gin"
)

func RouterGroup(router *gin.Engine) *gin.RouterGroup {
	httprequest := router.Group("/api")
	//httprequest.Use(middlewares.AuthJWT(), middlewares.PermissionCheck())
	EtcdRouter(httprequest)
	DnsRouter(httprequest)
	ZonesRouter(httprequest)
	return httprequest
}
