package sys

import (
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
)

// RoleAll 获取所有数据
func RoleAll() (role []sys.Role) {
	database.DB.Find(&role)
	return
}

func RoleDeletelById(id int) {
	database.DB.Where("id=?", id).Delete(&sys.Role{})

}
