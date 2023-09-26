package sys

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	"github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/handleerror"
	"github.com/yahahaff/rapide/pkg/logger"
)

// MenuController 菜单控制器
type MenuController struct {
	api.BaseAPIController
}

// GetMenuList 获取所有菜单
// @BasePath
// @Summary 获取所有菜单
// @Security Bearer
// @Schemes sys.PaginationRequest{}
// @Param per_page query int false "per_page"
// @Param page query int false "page"
// @Param sort query string false "sort"
// @Param order query string false "order"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/menu/getMenuList [get]
func (mc *MenuController) GetMenuList(c *gin.Context) {
	request := sys.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	data, pager, err := service.Entrance.SysService.MenuService.GetMenuList(request.Page, request.PerPage, request.Sort, request.Order)
	// 如果错误存在，记录错误日志，并抛出异常
	if err != nil {
		switch {
		case errors.Is(err, handleerror.ErrNotFound):
			logger.ErrorString("user", "error", fmt.Sprintf(err.Error()))
			response.Abort404(c, "未找到数据")
		default:
			logger.ErrorString("user", "error", fmt.Sprintf(err.Error()))
			response.Abort500(c)
		}
		return
	}
	result := gin.H{
		"datalist": data,
		"pager":    pager,
	}
	response.OK(c, result)
}

// GetMenu 获取用户菜单
// @BasePath
// @Summary 获取用户菜单
// @Security Bearer
// @Schemes
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/menu/getMenu [get]
func (mc *MenuController) GetMenu(c *gin.Context) {

	roleID := c.GetInt("current_user_role_id")
	data, err := service.Entrance.SysService.MenuService.GetMenu(roleID)
	if err != nil {
		fmt.Println(err.Error())
	}
	response.OK(c, data)
}

// AddMenu 添加菜单
// @BasePath
// @Summary 添加菜单
// @Security Bearer
// @Schemes requests.sys.MenuRequest{}
// @Param data body sys.MenuRequest{} true "body"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/menu/addMenu [post]
func (mc *MenuController) AddMenu(c *gin.Context) {

	// 1. 验证表单
	request := sys.MenuRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	data, err := service.Entrance.SysService.MenuService.AddMenu(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	response.OK(c, data)

}

// UpdateMenuByID 更新菜单
// @BasePath
// @Summary 更新菜单
// @Security Bearer
// @Schemes sys.MenuRequest{}
// @Param data body sys.MenuRequest{} true "body"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/menu/updateMenu [put]
func (mc *MenuController) UpdateMenuByID(c *gin.Context) {

	// 1. 验证表单
	request := sys.MenuRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	data, err := service.Entrance.SysService.MenuService.AddMenu(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	response.OK(c, data)

}

// DelMenuByID 删除菜单
// @BasePath
// @Summary 删除菜单
// @Security Bearer
// @Schemes requests.sys.MenuRequest{}
// @Param data body sys.MenuRequest{} true "body"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/menu/deleteMenu [delete]
func (mc *MenuController) DelMenuByID(c *gin.Context) {

	// 1. 验证表单
	request := sys.MenuRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	err := service.Entrance.SysService.MenuService.DelMenu(uint(request.ID))
	if err != nil {
		response.Abort500(c)
		return
	}
	response.Success(c)

}
