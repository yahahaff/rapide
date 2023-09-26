package sys

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/database"
)

// Menu 菜单表
type Menu struct {
	models.BaseModel
	Code     string  `json:"code" gorm:"type:varchar(50);comment:'菜单编号'"`
	PCode    string  `json:"p_code" gorm:"type:varchar(50);comment:'菜单父编号'"`
	PCodes   string  `json:"p_codes" gorm:"type:text;comment:'当前菜单的所有父菜单编号'"`
	Name     string  `json:"name" gorm:"type:varchar(100);comment:'菜单名称'"`
	Icon     string  `json:"icon" gorm:"type:varchar(100);comment:'菜单图标'"`
	URL      string  `json:"url" gorm:"type:varchar(255);default:'#';comment:'URL地址'"`
	Method   string  `json:"method" gorm:"type:varchar(10);comment:'Method'"`
	Sort     int64   `json:"sort" gorm:"comment:'菜单排序'"`
	Children []*Menu `json:"children" gorm:"foreignkey:PCode"`
	Levels   int64   `json:"levels" gorm:"comment:'菜单层级'"`
	IsMenu   bool    `json:"is_menu" gorm:"comment:'是否是菜单: 1:是 0:不是'"`
	Tips     string  `json:"tips" gorm:"type:varchar(255);comment:'备注'"`
	Status   bool    `json:"status" gorm:"comment:'菜单状态: 1:启用 0:不启用'"`
	Roles    []Role  `json:"roles" gorm:"many2many:sys_role_menus;"`
	models.CommonTimestampsField
}

// TableName Set the table name
func (menuModel *Menu) TableName() string {
	return "sys_menu"
}

// Create 创建菜单，通过 User.ID 来判断是否创建成功
func (menuModel *Menu) Create() {
	database.DB.Create(&menuModel)
}

// Delete 删除菜单，通过 User.ID 来判断是否创建成功
func (menuModel *Menu) Delete() {
	database.DB.Delete(&menuModel)
}

// TableOptions Set the table options
func (OperationLog) TableOptions() string {
	return "ENGINE=InnoDB AUTO_INCREMENT=1215187098755629063 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单表'"
}
