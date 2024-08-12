package sys

import (
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
)

type MenuService struct{}

// GetRoleIDByUserID 根据用户ID查询角色ID列表
func (s *MenuService) GetRoleIDByUserID(userID string) ([]uint64, error) {
	var user sys.User
	err := database.DB.Preload("Roles").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	var roleIDs []uint64
	for _, role := range user.Roles {
		roleIDs = append(roleIDs, role.ID)
	}

	return roleIDs, nil
}

func (s *MenuService) GetMenuTreeByRoleID(roleID uint64) ([]*sys.Menu, error) {
	var role sys.Role
	err := database.DB.Preload("Menus.Children").First(&role, roleID).Error
	if err != nil {
		return nil, err
	}
	// 构建菜单树
	menuTree := buildMenuTree(role.Menus, nil)
	return menuTree, nil
}

// buildMenuTree 构建菜单树
func buildMenuTree(menus []sys.Menu, parentID *uint64) []*sys.Menu {
	var menuTree []*sys.Menu
	for _, menu := range menus {
		// 处理根菜单的情况（ParentID 为 nil）
		if parentID == nil {
			if menu.ParentID == nil {
				// 创建 menu 的副本
				menuCopy := menu
				children := buildMenuTree(menus, &menu.ID)
				menuCopy.Children = children
				menuTree = append(menuTree, &menuCopy)
			}
		} else {
			// 处理子菜单的情况
			if menu.ParentID != nil && *menu.ParentID == *parentID {
				// 创建 menu 的副本
				menuCopy := menu
				children := buildMenuTree(menus, &menu.ID)
				menuCopy.Children = children
				menuTree = append(menuTree, &menuCopy)
			}
		}
	}
	return menuTree
}
