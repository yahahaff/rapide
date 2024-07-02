// Package sys
package sys

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel
	Name         string `json:"name" gorm:"type:varchar(255);uniqueIndex;not null"`
	Email        string `json:"email" gorm:"type:varchar(255);uniqueIndex;default:null"`
	Phone        string `json:"phone" gorm:"type:varchar(20);uniqueIndex;default:null"`
	Password     string `json:"-"`
	Avatar       string `json:"avatar,omitempty"`
	RoleID       int    `json:"roleID" gorm:"column:role_id"`
	Role         Role   `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeptID       int    `json:"deptID" gorm:"column:dept_id"`
	Dept         Dept   `json:"dept" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status       int    `json:"status"`
	OtpEnabled   bool   `json:"otpEnabled" gorm:"default:false;"`
	OtpVerified  bool   `json:"otpVerified" gorm:"default:false;"`
	OtpSecret    string `json:"-"`
	OtpAuthUrl   string `json:"-"`
	Introduction string `json:"introduction,omitempty"`
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
