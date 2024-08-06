package sys

import (
	"github.com/yahahaff/rapide/backend/internal/models/sys"
	"github.com/yahahaff/rapide/backend/pkg/database"
)

type MenuService struct{}

// GetMenuList 获取菜单列表
func (m *MenuService) GetMenuList() (data []sys.Menu, err error) {
	var menus []sys.Menu
	db := database.DB.
		Preload("Meta").
		Preload("Children.Meta").
		Preload("Children.Children.Meta")

	// 查询数据
	if err = db.Where("parent_id IS NULL").Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}
