package sys

import "github.com/gin-gonic/gin"

func RouterGroup(router *gin.Engine) *gin.RouterGroup {
	internal := router.Group("")
	InternalRouter(internal)

	sys := router.Group("/api")
	AuthRouter(sys)
	AuthenticatorRouter(sys)
	CaptchaRouter(sys)
	CasbinRouter(sys)
	UserRouter(sys)
	OperationLogRouter(sys)
	MenuRouter(sys)
	return sys
}
