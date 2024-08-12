package sys

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/database"
)

// Role 角色
type Role struct {
	models.BaseModel
	RoleName    string       `json:"role_name" gorm:"unique"`
	Users       []User       `json:"-" gorm:"many2many:sys_user_roles;" `
	Status      int          `json:"status"     `
	Permissions []Permission `json:"permissions" gorm:"many2many:sys_role_permissions;" `
	Menus       []Menu       `json:"menus" gorm:"many2many:sys_role_menus;"` // 添加菜单关联
	models.CommonTimestampsField
}

// TableName Set the table name
func (*Role) TableName() string {
	return "sys_role"
}

// Create 创建部门，通过 ID 来判断是否创建成功
func (roleModel *Role) Create() {
	database.DB.Create(&roleModel)
}
