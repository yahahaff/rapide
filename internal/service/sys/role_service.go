package sys

import (
	"errors"
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
)

type RoleService struct {
	roleModel sys.Role
	menuModel sys.Menu
}

// AssignRoleMenuByID 验证角色和菜单并更新角色菜单关联
func (rs *RoleService) AssignRoleMenuByID(roleID int64, menuIDs []int64) error {
	// 验证角色是否存在
	if err := database.DB.First(&rs.roleModel, roleID).Error; err != nil {
		return errors.New("角色不存在")
	}

	// 验证菜单是否存在
	//var menus []sys.Menu
	if err := database.DB.Find(&rs.menuModel, menuIDs).Error; err != nil {
		return errors.New("菜单不存在")
	}

	// 更新角色菜单关联
	if err := database.DB.Model(&rs.roleModel).Association("Menus").Replace(&rs.menuModel); err != nil {
		return errors.New("更新角色菜单关联失败")
	}

	return nil
}
