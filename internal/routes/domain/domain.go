package domain

import (
	"github.com/gin-gonic/gin"
)

func RouterGroup(router *gin.Engine) *gin.RouterGroup {
	domain := router.Group("")
	//cloudflare.DomainRouter(domain)
	// add other sys related routes to the group
	return domain
}
