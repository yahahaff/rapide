package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/sys"
)

func MenuRouter(Router *gin.RouterGroup) {

	//登录路由
	{
		authGroup := Router.Group("menu")
		smc := new(sys.MenuController)

		authGroup.GET("/all", smc.GetMenuList)

	}

}
