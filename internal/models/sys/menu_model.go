package sys

import (
	"github.com/yahahaff/rapide/internal/models"
)

// Menu 表示菜单项
type Menu struct {
	models.BaseModel
	ParentID  *uint64                                    `json:"parent_id" gorm:"index;default:NULL" ` // 父菜单项ID
	Name      string                                     `json:"name" gorm:"size:255;not null"`
	Path      string                                     `json:"path" gorm:"size:255;not null"`
	Component string                                     `json:"component" gorm:"size:255"`
	Redirect  string                                     `json:"redirect" gorm:"size:255"`
	Meta      `json:"meta" gorm:"embedded;comment:附加属性"` // 附加属性
	Children  []*Menu                                    `json:"children" gorm:"foreignkey:ParentID;association_foreignkey:ID"`
	models.CommonTimestampsField
}

// Meta 表示菜单元数据
type Meta struct {
	Order                    int    `json:"order" gorm:"default:0 "` // 用于菜单排序
	Title                    string `json:"title" gorm:"size:255"`
	Icon                     string `json:"icon" gorm:"size:255"`
	KeepAlive                bool   `json:"keep_alive" gorm:"default:false"` // 是否缓存
	AffixTab                 bool   `json:"affixTab" gorm:"default:false"`
	Authority                string `json:"authority" gorm:"size:255"`
	MenuVisibleWithForbidden bool   `json:"menu_visible_with_forbidden" gorm:"default:false"`
}

func (Menu) TableName() string {
	return "sys_menu"
}
