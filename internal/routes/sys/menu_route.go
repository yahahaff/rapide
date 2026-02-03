package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers/sys"
)

func MenuRouter(Router *gin.RouterGroup) {
	// 菜单路由
	{
		menuGroup := Router.Group("menu")
		mc := new(sys.MenuController)
		menuGroup.GET("/all", mc.GetUserMenus)
		menuGroup.GET("/list", mc.GetMenuList)
		menuGroup.POST("/create", mc.CreateMenu)
		menuGroup.DELETE("/delete/:id", mc.DeleteMenu)

	}
}
