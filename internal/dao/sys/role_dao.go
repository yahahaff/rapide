package sys

import (
	"rapide/internal/models/sys"
	"rapide/pkg/database"
)

// GetRoles 获取角色列表（带分页）
func GetRoles(page, pageSize int) (roles []map[string]interface{}, total int64) {
	// 查询总数
	database.DB.Model(&sys.Role{}).Count(&total)

	// 查询角色列表
	var roleModels []sys.Role
	offset := (page - 1) * pageSize
	database.DB.Offset(offset).Limit(pageSize).Find(&roleModels)

	// 转换为符合要求的数据格式
	roles = make([]map[string]interface{}, len(roleModels))
	for i, role := range roleModels {
		// 模拟权限数据，实际项目中应该从数据库查询
		permissions := []string{}
		if role.RoleCode == "admin" {
			permissions = []string{"*:*:*"}
		} else {
			permissions = []string{"user:*:*"}
		}

		roles[i] = map[string]interface{}{
			"id":         role.ID,
			"name":       role.RoleName,
			"code":       role.RoleCode,
			"sort":       role.Sort,
			"status":     role.Status,
			"createTime": role.CreatedAt.Format("2006-01-02T15:04:05Z"),
			"remark":     role.Remark,
			"permissions": permissions,
		}
	}

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
