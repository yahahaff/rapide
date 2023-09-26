package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/paginator"
)

type MenuService struct {
	menuModel sys.Menu
}

// GetMenuList 获取菜单列表
func (m *MenuService) GetMenuList(page, size int, sort, order string) (data interface{}, pager paginator.Paging, err error) {

	var menus []sys.Menu
	db := database.DB.Where("p_code = ?", 0).
		Preload("Children").
		Preload("Children.Children").
		Preload("Children.Children.Children")

	// 判断是否需要排序
	if sort != "" && order != "" {
		db = db.Order(fmt.Sprintf("%s %s", sort, order))
	}
	// 查询数据
	db.Find(&menus)

	// 分页数据
	data, pager = paginator.Paginate(menus, page, size)

	return
}

// GetMenu 获取个人菜单
func (m *MenuService) GetMenu(roleID int) ([]sys.Menu, error) {
	var role sys.Role
	// 支持到三级菜单
	if err := database.DB.Preload("Menus", "p_code = ? AND status = ?", 0, 1).
		Preload("Menus.Children").
		Preload("Menus.Children.Children").
		First(&role, roleID).Error; err != nil {
		return nil, err
	}
	return role.Menus, nil
}

// AddMenu 新增菜单
func (m *MenuService) AddMenu(c *gin.Context) (data interface{}, err error) {
	menuModel := sys.Menu{
		Name: c.Request.Host,
	}
	menuModel.Create()
	return
}

// DelMenu 删除菜单
func (m *MenuService) DelMenu(id uint) (err error) {
	database.DB.Where("id = ?", id).Delete(&m.menuModel)
	return
}
