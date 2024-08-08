package sys

import (
	"github.com/yahahaff/rapide/backend/internal/models/sys"
	"github.com/yahahaff/rapide/backend/pkg/database"
)

type MenuService struct{}

func (s *MenuService) GetMenuTreeByRoleID(roleID uint64) ([]*sys.Menu, error) {
	var role sys.Role
	err := database.DB.Preload("Menus.Children").First(&role, roleID).Error
	if err != nil {
		return nil, err
	}

	// 假设 role.Menus 是 []sys.Menu 类型，需要转换为 []*sys.Menu
	var menus []*sys.Menu
	for i := range role.Menus {
		menu := role.Menus[i]
		menus = append(menus, &menu)
	}

	// 构建 ID 到菜单指针的映射
	menuMap := make(map[uint64]*sys.Menu)
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}

	// 构建菜单树
	var menuTree []*sys.Menu
	for _, menu := range menus {
		if menu.ParentID == nil {
			menuTree = append(menuTree, menu) // 添加菜单项
		} else {
			parentID := *menu.ParentID
			if parentMenu, exists := menuMap[parentID]; exists {
				// 确保子菜单只添加一次
				if !contains(parentMenu.Children, menu) {
					parentMenu.Children = append(parentMenu.Children, menu)
				}
			}
		}
	}

	return menuTree, nil
}

// 检查切片中是否包含某个菜单项
func contains(slice []*sys.Menu, item *sys.Menu) bool {
	for _, menu := range slice {
		if menu.ID == item.ID {
			return true
		}
	}
	return false
}
