package sys

import "github.com/gin-gonic/gin"

func RouterGroup(router *gin.Engine) *gin.RouterGroup {
	sys := router.Group("/api")
	internal := router.Group("")
	InternalRouter(internal)
	AuthRouter(sys)
	UserRouter(sys)
	CaptchaRouter(sys)
	CasbinRouter(sys)
	OperationLogRouter(sys)
	MenuRouter(sys)
	return sys
}
