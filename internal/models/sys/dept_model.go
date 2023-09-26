package sys

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/database"
)

type Dept struct {
	models.BaseModel
	Num      int    `json:"num" gorm:"comment:'排序'"`
	PID      int    `json:"pid" gorm:"comment:'父部门id'"`
	Pids     string `json:"pids" gorm:"type:varchar(10);comment:'父级ids'"`
	FullName string `json:"full_name" gorm:"type:varchar(10);comment:'全称'"`
	Tips     string `json:"tips" gorm:"type:varchar(10);comment:'提示'"`
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
