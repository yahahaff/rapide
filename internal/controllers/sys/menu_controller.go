package sys

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	"rapide/internal/models/sys"
	sysReq "rapide/internal/requests/sys"
	"rapide/internal/requests/validators"
	"rapide/internal/service"
	"rapide/internal/utils"
	"rapide/pkg/database"
	"rapide/pkg/response"
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
	response.OK(c, menus)
}

// GetMenuList 获取菜单列表
func (mc *MenuController) GetMenuList(c *gin.Context) {
	request := sysReq.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 20 // 设置默认值
	}

	// 处理页码参数，确保页码大于0
	page := request.Page
	if page <= 0 {
		page = 1
	}

	// 查询一级菜单总数
	var total int64
	database.DB.Model(&sys.Menu{}).Where("parent_id IS NULL").Count(&total)

	// 一次性查询所有菜单，不分页
	var allMenus []*sys.Menu
	database.DB.Find(&allMenus)

	// 构建菜单树
	menuTree := utils.BuildMenuTree(allMenus)

	// 构造返回数据
	responseData := map[string]interface{}{
		"result": menuTree,
		"total":  total,
	}

	response.OK(c, responseData)
}

// CreateMenu 创建菜单
func (mc *MenuController) CreateMenu(c *gin.Context) {
	// 1. 解析和验证请求体
	request := sysReq.MenuCreateRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 处理菜单类型映射：vben-admin 的类型是字符串，数据库中是数字
	// vben-admin: catalog(目录)、menu(菜单)、button(按钮)、iframe(内嵌)、link(外链)
	// 数据库: 0(目录)、1(菜单)、2(按钮)
	var menuType int
	switch request.Type {
	case "catalog":
		menuType = 0
	case "menu":
		menuType = 1
	case "button":
		menuType = 2
	default:
		menuType = 1 // 默认菜单
	}

	// 3. 处理菜单状态：将int转换为bool
	menuStatus := request.Status == 1

	// 4. 创建菜单模型
	menu := &sys.Menu{
		Name:               request.Name,
		Title:              request.Title,
		Path:               request.Path,
		Component:          request.Component,
		Redirect:           request.Redirect,
		ParentID:           request.ParentID,
		Icon:               request.Icon,
		OrderNo:            request.OrderNo,
		Type:               menuType,
		Status:             menuStatus,
		KeepAlive:          request.KeepAlive,
		Hidden:             request.Hidden,
		HideBreadcrumb:     request.HideBreadcrumb,
		HideChildrenInMenu: request.HideChildrenInMenu,
		AffixTab:           request.AffixTab,
		NoBasicLayout:      request.NoBasicLayout,
		IgnoreKeepAlive:    request.IgnoreKeepAlive,
		CurrentActiveMenu:  request.CurrentActiveMenu,
	}

	// 5. 保存到数据库
	if err := database.DB.Create(menu).Error; err != nil {
		response.Abort500(c, "创建菜单失败")
		return
	}

	// 6. 返回响应
	response.OK(c, menu)
}

// DeleteMenu 删除菜单
func (mc *MenuController) DeleteMenu(c *gin.Context) {
	// 1. 从URL路径中获取菜单ID
	idStr := c.Param("id")
	var id uint64
	if _, err := fmt.Sscan(idStr, &id); err != nil {
		response.Abort500(c, "无效的菜单ID")
		return
	}

	// 2. 检查菜单是否存在
	var menu sys.Menu
	if err := database.DB.First(&menu, id).Error; err != nil {
		response.Abort500(c, "菜单不存在")
		return
	}

	// 3. 删除菜单
	if err := database.DB.Delete(&sys.Menu{}, id).Error; err != nil {
		response.Abort500(c, "删除菜单失败")
		return
	}

	// 4. 返回响应
	response.OK(c, gin.H{"id": id})
}
