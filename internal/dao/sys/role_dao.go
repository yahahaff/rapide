package sys

import (
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
)

// GetRolesWithChildren 获取角色及其子角色
func GetRolesWithChildren(pageSize int) (roles []sys.Role) {
	db := database.DB.Where("parent_id IS NULL").
		Preload("Children").
		Preload("Children.Children")

	// 添加分页限制
	if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	// 查询数据
	db.Find(&roles)
	return
}

// GetRoleById 通过 用户ID 获取用户角色Model
func GetRoleById(roleID int64) (roleModel sys.Role) {
	database.DB.First(&roleModel, roleID)
	return
}

func RoleDeletelById(id int) {
	database.DB.Where("id=?", id).Delete(&sys.Role{})

}
