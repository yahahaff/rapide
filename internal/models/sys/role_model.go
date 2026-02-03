package sys

import (
	"rapide/internal/models"
	"rapide/pkg/database"
)

// Role 角色 - 适配 vue-vben-admin 5.x
type Role struct {
	models.BaseModel

	// 基本信息
	RoleName  string `json:"roleName" gorm:"type:varchar(255);unique;comment:'角色名称'"`
	RoleValue string `json:"roleValue" gorm:"type:varchar(255);unique;comment:'角色值'"`
	RoleCode  string `json:"roleCode" gorm:"type:varchar(255);unique;comment:'角色编码'"`

	// 父子关系
	ParentID *uint64 `json:"parentId" gorm:"comment:'父角色ID'"`
	Children []*Role `json:"children" gorm:"foreignKey:ParentID"`

	// 菜单关联
	Menus []*Menu `json:"menus" gorm:"many2many:sys_role_menu;"`

	// 其他信息
	Sort   int    `json:"sort" gorm:"comment:'排序'"`
	Status int    `json:"status" gorm:"default:1;comment:'状态 0:禁用 1:启用'"`
	Remark string `json:"remark" gorm:"type:varchar(500);comment:'备注'"`

	models.CommonTimestampsField
}

// TableName Set the table name
func (*Role) TableName() string {
	return "sys_role"
}

// Create 创建角色
func (roleModel *Role) Create() {
	database.DB.Create(&roleModel)
}

// Update 更新角色
func (roleModel *Role) Update() error {
	return database.DB.Save(&roleModel).Error
}

// Delete 删除角色
func (roleModel *Role) Delete() {
	database.DB.Delete(&roleModel)
}

// AssignMenus 为角色分配菜单
func (roleModel *Role) AssignMenus(menuIDs []uint64) error {
	// 清除现有菜单关联
	err := database.DB.Model(roleModel).Association("Menus").Clear()
	if err != nil {
		return err
	}

	// 添加新的菜单关联
	var menus []*Menu
	err = database.DB.Where("id IN ?", menuIDs).Find(&menus).Error
	if err != nil {
		return err
	}

	return database.DB.Model(roleModel).Association("Menus").Append(menus)
}
