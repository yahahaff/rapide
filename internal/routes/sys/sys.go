package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func RouterGroup(router *gin.Engine) {
	// 内部公开路由
	internal := router.Group("")
	InternalRouter(internal)

	// 公开路由
	auth := router.Group("/api/auth")
	LoginRouter(auth)

	// 注册路由
	signup := router.Group("/api/auth")
	SingupRouter(signup)

	captcha := router.Group("/api/captcha")
	CaptchaRouter(captcha)

	// sys相关需要jwt认证路由
	sys := router.Group("/api")
	//JWT认证 接口权限校验 && 日志记录
	sys.Use(middlewares.AuthJWT(), middlewares.RecordOperation())
	MenuRouter(sys)
	CasbinRouter(sys)
	UserRouter(sys)
	OperationLogRouter(sys)
	SSLCertRouter(sys)
}
