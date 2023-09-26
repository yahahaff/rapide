package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/sys"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func MenuRouter(Router *gin.RouterGroup) {
	menuGroup := Router.Group("/menu")
	menuGroup.Use(middlewares.AuthJWT(), middlewares.PermissionCheck())

	//菜单路由
	{
		smc := new(sys.MenuController)
		menuGroup.GET("getMenu", smc.GetMenu)
		menuGroup.GET("getMenuList", smc.GetMenuList)
		menuGroup.POST("addMenu", smc.AddMenu)
		menuGroup.DELETE("deleteMenu", smc.DelMenuByID)
		menuGroup.PUT("updateMenu", smc.UpdateMenuByID)

	}
}
