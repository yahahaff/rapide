package sys

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/response"
)

// MenuController 菜单控制器
type MenuController struct {
	controllers.BaseAPIController
}

// GetUserMenus 获取用户菜单
func (mc *MenuController) GetUserMenus(c *gin.Context) {
	// 获取当前用户角色ID
	userRoleId := c.GetUint64("current_user_role_id")
	menus, err := service.Entrance.SysService.MenuService.GetUserMenus(userRoleId)
	if err != nil {
		response.Abort500(c, "获取用户菜单失败")
		return
	}
	fmt.Println(menus)
	response.OK(c, menus)
}
