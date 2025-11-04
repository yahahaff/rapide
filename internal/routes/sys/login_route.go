package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/sys"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func LoginRouter(Router *gin.RouterGroup) {

	//登录路由
	{
		loginGroup := Router.Group("")
		lgc := new(sys.LoginController)
		// 使用用户名密码登录
		loginGroup.POST("/login", middlewares.LoginFailureCheck(), lgc.LoginByPassword)

	}

}
