package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/sys"
)

func MenuRouter(Router *gin.RouterGroup) {
	// 菜单路由
	{
		menuGroup := Router.Group("menu")
		mc := new(sys.MenuController)
		menuGroup.GET("/all", mc.GetUserMenus)
		menuGroup.GET("/list", mc.GetMenuList)

	}
}
