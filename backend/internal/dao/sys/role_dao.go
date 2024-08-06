package sys

import (
	"github.com/yahahaff/rapide/backend/internal/models/sys"
	"github.com/yahahaff/rapide/backend/pkg/database"
)

// GetRolesWithChildren 获取角色及其子角色
func GetRolesWithChildren() (roles []sys.Role) {
	db := database.DB.Where("p_codes IS NULL").
		Preload("Children").
		Preload("Children.Children")

	// 查询数据
	db.Find(&roles)
	return
}

func RoleDeletelById(id int) {
	database.DB.Where("id=?", id).Delete(&sys.Role{})

}
