package sys

import (
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
)

// DeptAll 获取所有数据
func DeptAll() (depots []sys.Dept) {
	database.DB.Find(&depots)
	return
}

func DeptDeletelById(id int) {
	database.DB.Where("id=?", id).Delete(&sys.Dept{})

}
