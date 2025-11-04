package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/pkg/response"
)

// MenuController 菜单控制器
type MenuController struct {
	controllers.BaseAPIController
}

// GetUserMenus 获取当前用户的菜单列表
//func (mc *MenuController) GetUserMenus(c *gin.Context) {
//	// 获取当前用户角色ID
//	userOrleId := c.GetUint64("current_user_role_id")
//	// 获取用户的菜单树
//	menuTree, err := service.Entrance.SysService.MenuService.GetMenuTreeByRoleID(userOrleId)
//	if err != nil {
//		response.Abort500(c, "获取用户菜单失败")
//		return
//	}
//	response.OK(c, menuTree)
//}

// GetUserMenus 获取用户菜单
func (mc *MenuController) GetUserMenus(c *gin.Context) {

	b := "fddf"
	response.Success(c, b)
}
