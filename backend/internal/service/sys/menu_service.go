package sys

import (
	"github.com/yahahaff/rapide/backend/internal/models/sys"
	"github.com/yahahaff/rapide/backend/pkg/database"
)

type MenuService struct{}

type MenuResponse struct {
	Component string         `json:"component"`
	Meta      sys.MenuMeta   `json:"meta"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Redirect  string         `json:"redirect"`
	Children  []MenuResponse `json:"children,omitempty"`
}

// GetMenuList 获取菜单列表
func (m *MenuService) GetMenuList() (data []sys.Menu, err error) {
	var menus []sys.Menu
	db := database.DB.Where("parent_id IS NULL").
		Preload("Children").
		Preload("Children.Children")

	// 查询数据
	if err = db.Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}
