package sys

import (
	"github.com/yahahaff/rapide/backend/internal/models"
)

// Permission 权限表
type Permission struct {
	models.BaseModel
	Code        string `json:"code" gorm:"unique"`
	Description string `json:"description"`
	Roles       []Role `gorm:"many2many:sys_role_permissions;" json:"roles"`
	models.CommonTimestampsField
}

// TableName Set the table name
func (*Permission) TableName() string {
	return "sys_permission"
}
