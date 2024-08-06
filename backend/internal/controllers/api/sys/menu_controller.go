package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/backend/internal/controllers/api"
	"github.com/yahahaff/rapide/backend/internal/response"
	"github.com/yahahaff/rapide/backend/internal/service"
)

// MenuController 菜单控制器
type MenuController struct {
	api.BaseAPIController
}

func (mc *MenuController) GetMenuList(c *gin.Context) {
	data, _ := service.Entrance.SysService.MenuService.GetMenuList()
	response.OK(c, data)
	return

}
