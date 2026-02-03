package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers/sys"
)

func SingupRouter(Router *gin.RouterGroup) {
	// 注册路由
	{
		signupGroup := Router.Group("")
		sgc := new(sys.SignupController)
		// 使用用户名注册
		signupGroup.POST("/signup", sgc.SignupUsingUserName)
	}
}
