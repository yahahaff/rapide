package sys

import (
	"rapide/internal/models"
)

type Menu struct {
	models.BaseModel

	// 基本信息
	Name      string `json:"name" gorm:"type:varchar(100);comment:'菜单名称'"`
	Path      string `json:"path" gorm:"type:varchar(255);comment:'路由路径'"`
	Component string `json:"component" gorm:"type:varchar(255);comment:'组件路径'"`
	Redirect  string `json:"redirect" gorm:"type:varchar(255);comment:'重定向路径'"`

	// 父子关系 - 不使用外键约束
	ParentID *uint64 `json:"parentId" gorm:"index:idx_parent_id;comment:'父菜单ID'"`
	Children []*Menu `json:"children" gorm:"-"`

	// 菜单元数据
	Title         string `json:"title" gorm:"type:varchar(100);comment:'菜单标题'"`
	Icon          string `json:"icon" gorm:"type:varchar(100);comment:'菜单图标'"`
	OrderNo       int    `json:"orderNo" gorm:"default:0;comment:'菜单排序'"`
	KeepAlive     bool   `json:"keepAlive" gorm:"default:true;comment:'是否缓存'"`
	Hidden        bool   `json:"hidden" gorm:"default:false;comment:'是否隐藏'"`
	AffixTab      bool   `json:"affixTab" gorm:"default:false;comment:'是否固定标签页'"`
	NoBasicLayout bool   `json:"noBasicLayout" gorm:"default:false;comment:'是否不使用基础布局'"`

	// 权限相关
	//Permission string `json:"permission" gorm:"type:varchar(255);comment:'权限标识'"`
	Type   int  `json:"type" gorm:"comment:'菜单类型 (0:目录 1:菜单 2:按钮)'"`
	Status bool `json:"status" gorm:"default:true;comment:'菜单状态: true:启用 false:禁用'"`

	// 扩展字段
	IgnoreKeepAlive    bool   `json:"ignoreKeepAlive" gorm:"default:false;comment:'是否忽略KeepAlive缓存'"`
	HideBreadcrumb     bool   `json:"hideBreadcrumb" gorm:"default:false;comment:'隐藏面包屑'"`
	HideChildrenInMenu bool   `json:"hideChildrenInMenu" gorm:"default:false;comment:'隐藏所有子菜单'"`
	CurrentActiveMenu  string `json:"currentActiveMenu" gorm:"type:varchar(255);comment:'当前激活的菜单'"`

	models.CommonTimestampsField
}

func (Menu) TableName() string {
	return "sys_menu"
}

// RoleMenu 中间表模型
type RoleMenu struct {
	RoleID uint64 `gorm:"primaryKey;autoIncrement:false"`
	MenuID uint64 `gorm:"primaryKey;autoIncrement:false"`
}

func (RoleMenu) TableName() string {
	return "sys_role_menu"
}
