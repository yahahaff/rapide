package sys

import (
	"github.com/yahahaff/rapide/backend/internal/models"
)

// Menu 表示菜单项
type Menu struct {
	models.BaseModel
	ParentID  uint     `json:"parent_id" gorm:"index"` // 父菜单项ID
	Name      string   `json:"name" gorm:"size:255;not null"`
	Path      string   `json:"path" gorm:"size:255;not null"`
	Component string   `json:"component" gorm:"size:255"`
	Redirect  string   `json:"redirect" gorm:"size:255"`
	Order     int      `json:"order" gorm:"default:0"` // 用于菜单排序
	Meta      MenuMeta `json:"meta" gorm:"foreignKey:ID"`
	Children  []*Menu  `json:"children" gorm:"foreignKey:ParentID"`
	models.CommonTimestampsField
}

// MenuMeta 表示菜单元数据
type MenuMeta struct {
	models.BaseModel
	Title                    string `json:"title" gorm:"size:255"`
	Icon                     string `json:"icon" gorm:"size:255"`
	KeepAlive                bool   `json:"keep_alive" gorm:"default:false"`
	AffixTab                 bool   `json:"affixTab" gorm:"default:false"`
	Authority                string `json:"authority" gorm:"size:255"`
	MenuVisibleWithForbidden bool   `json:"menu_visible_with_forbidden" gorm:"default:false"`
	models.CommonTimestampsField
}

func (Menu) TableName() string {
	return "sys_menu"
}

func (MenuMeta) TableName() string {
	return "sys_menu_meta"
}
