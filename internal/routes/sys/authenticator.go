package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/sys"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func AuthenticatorRouter(Router *gin.RouterGroup) {

	// 2FA路由  authenticator
	{
		authenticatorGroup := Router.Group("/authenticator")
		ag := new(sys.AuthenticatorController)
		authenticatorGroup.POST("/generate", middlewares.AuthJWT(), ag.Generate)
		authenticatorGroup.POST("/disable", middlewares.AuthJWT(), ag.Disable)
		//绑定2FA时调用 写入密钥、更新状态等到数据库
		authenticatorGroup.POST("/verify", middlewares.AuthJWT(), ag.Verify)
		//验证码验证
		authenticatorGroup.POST("/validate", ag.Validate)
	}

}
