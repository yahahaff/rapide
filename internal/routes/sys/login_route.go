package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers/sys"
)

func LoginRouter(Router *gin.RouterGroup) {

	//登录路由
	{
		loginGroup := Router.Group("")
		lgc := new(sys.LoginController)
		// 使用用户名密码登录 或 手机验证码登录（合并的登录接口）
		loginGroup.POST("/login", lgc.Login)

	}

}
