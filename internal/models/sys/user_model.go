// Package user 存放用户 Model 相关逻辑
package sys

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel
	Name         string `json:"name" gorm:"type:varchar(255);not null;index"`
	Email        string `json:"email" gorm:"type:varchar(255);index;default:null"`
	Phone        string `json:"phone" gorm:"type:varchar(20);index;default:null"`
	Password     string `json:"-" json:"-" gorm:"type:varchar(255)"`
	Avatar       string `json:"avatar,omitempty"`
	RoleID       int    `json:"roleID" gorm:"column:role_id"`
	Role         Role   // 关联的 Role 结构体
	DeptID       int    `json:"deptID" gorm:"column:dept_id"`
	Dept         Dept
	Status       int    `json:"status" gorm:"column:status"`
	OtpEnabled   bool   `json:"otpEnabled" gorm:"default:false;"`
	OtpVerified  bool   `json:"OtpVerified" gorm:"default:false;"`
	OtpSecret    string `json:"-" gorm:"type:varchar(255)"`
	OtpAuthUrl   string `json:"o" gorm:"type:varchar(255)"`
	Introduction string `json:"introduction,omitempty" gorm:"type:varchar(255)"`
	models.CommonTimestampsField
}

// TableName Set the table name
func (*User) TableName() string {
	return "sys_user"
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
