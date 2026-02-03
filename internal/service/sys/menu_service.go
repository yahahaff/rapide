package sys

import (
	"rapide/internal/models/sys"
	"rapide/internal/utils"
	"rapide/pkg/database"
)

// MenuService 菜单服务
type MenuService struct{}

// GetUserMenus 根据角色ID获取用户菜单树（适配 Vben Admin 5.x）
func (s *MenuService) GetUserMenus(roleID uint64) ([]map[string]interface{}, error) {
	if roleID == 0 {
		return []map[string]interface{}{}, nil
	}

	// 1. 获取该角色关联的所有菜单ID
	var menuIDs []uint64
	err := database.DB.Table("sys_role_menu").
		Where("role_id = ?", roleID).
		Pluck("menu_id", &menuIDs).Error

	if err != nil {
		return nil, err
	}

	if len(menuIDs) == 0 {
		return []map[string]interface{}{}, nil
	}
	// 2. 一次性查出所有启用的、类型为目录(0)或菜单(1)的菜单
	var menus []*sys.Menu
	err = database.DB.Where("id IN ? AND status = true AND type IN (0, 1)", menuIDs).
		Order("order_no ASC").
		Find(&menus).Error
	if err != nil {
		return nil, err
	}
	// 3. 构建的菜单树
	menuTree := utils.BuildMenuTree(menus)
	return menuTree, nil
}
