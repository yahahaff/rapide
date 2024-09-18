package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
)

// MenuController 菜单控制器
type MenuController struct {
	controllers.BaseAPIController
}

func (mc *MenuController) GetMenuList(c *gin.Context) {
	roleID := c.GetUint64("current_user_role_id")
	data, _ := service.Entrance.SysService.MenuService.GetMenuTreeByRoleID(roleID)
	response.OK(c, data)
	return

}
