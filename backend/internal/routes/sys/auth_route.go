package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/backend/internal/controllers/api/sys"
	"github.com/yahahaff/rapide/backend/internal/middlewares"
)

func AuthRouter(Router *gin.RouterGroup) {

	//登录路由
	{
		loginGroup := Router.Group("")
		lgc := new(sys.LoginController)
		// 使用用户名密码登录
		loginGroup.POST("/login", middlewares.LoginFailureCheck(), lgc.LoginByPassword)
		// 获取角色权限
		loginGroup.GET("/codes", lgc.GetRoleCodes)

	}

}
