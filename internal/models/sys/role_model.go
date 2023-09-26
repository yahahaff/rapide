package sys

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/database"
)

// Role 角色
type Role struct {
	models.BaseModel
	Num     int64  `json:"num" gorm:"column:num;type:int(11);comment:'序号'"`
	Pid     int64  `json:"pid" gorm:"column:pid;type:int(11);comment:'父角色id'"`
	Name    string `json:"name" gorm:"column:name;type:varchar(255);comment:'角色名称'"`
	Menus   []Menu `json:"menus" gorm:"many2many:sys_role_menus;"`
	Tips    string `json:"tips" gorm:"column:tips;type:varchar(255);comment:'提示'"`
	Version int64  `json:"version" gorm:"column:version;type:int(11);comment:'保留字段(暂时没用）'"`
}

// TableName Set the table name
func (*Role) TableName() string {
	return "sys_role"
}

// Create 创建部门，通过 ID 来判断是否创建成功
func (roleModel *Role) Create() {
	database.DB.Create(&roleModel)
}
