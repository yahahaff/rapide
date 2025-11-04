package sys

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/database"
)

// Dept 定义
type Dept struct {
	models.BaseModel
	PCode    int     `json:"p_code" gorm:"comment:'父部门id'"`
	PCodes   string  `json:"p_codes" gorm:"type:varchar(255);comment:'父级ids'"`
	Name     string  `json:"name" gorm:"type:varchar(255);unique;comment:'部门名称'"`
	Sort     int     `json:"sort" gorm:"comment:'排序'"`
	Level    int     `json:"level" gorm:"comment:'部门级别'"`
	Children []*Dept `json:"children" gorm:"foreignKey:PCode"`
	Tips     string  `json:"tips" gorm:"type:varchar(255);comment:'提示'"`
	models.CommonTimestampsField
}

// TableName Set the table name
func (*Dept) TableName() string {
	return "sys_dept"
}

// Create 创建部门，通过 ID 来判断是否创建成功
func (deptModel *Dept) Create() {
	database.DB.Create(&deptModel)
}
