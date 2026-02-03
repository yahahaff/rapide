package sys

import (
	"time"

	"rapide/internal/models"
	"rapide/pkg/database"
)

// Dept 部门模型
type Dept struct {
	models.BaseModel
	Pid        uint64     `json:"pid" gorm:"comment:'父部门ID';default:0"`
	ID         uint64     `json:"id" gorm:"primaryKey;autoIncrement;comment:'部门ID'"`
	Name       string     `json:"name" gorm:"type:varchar(255);comment:'部门名称'"`
	Status     int        `json:"status" gorm:"default:1;comment:'状态 0:禁用 1:启用'"`
	Remark     string     `json:"remark" gorm:"type:text;comment:'备注'"`
	CreateTime time.Time  `json:"createTime" gorm:"type:datetime;comment:'创建时间'"`
	Children   []*Dept    `json:"children" gorm:"-"`
	models.CommonTimestampsField
}

// TableName 设置表名
func (*Dept) TableName() string {
	return "sys_dept"
}

// Create 创建部门
func (dept *Dept) Create() error {
	return database.DB.Create(dept).Error
}

// GetDeptByID 根据ID获取部门
func GetDeptByID(id uint64) (Dept, error) {
	var dept Dept
	err := database.DB.Where("id = ?", id).First(&dept).Error
	return dept, err
}

// GetDeptList 获取所有部门列表
func GetDeptList() ([]Dept, error) {
	var deptList []Dept
	err := database.DB.Find(&deptList).Error
	return deptList, err
}

// Update 更新部门
func (dept *Dept) Update() error {
	return database.DB.Save(dept).Error
}

// Delete 删除部门
func DeleteDept(id uint64) error {
	return database.DB.Delete(&Dept{}, "id = ?", id).Error
}
