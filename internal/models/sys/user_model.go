// Package sys
package sys

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/hash"
)

type User struct {
	models.BaseModel
	Username    string `json:"username" `
	Password    string `json:"-" `
	RealName    string `json:"real_name" `
	RoleID      uint64 `json:"role_id"`                                      // RoleID 作为外键字段
	Role        Role   `gorm:"foreignKey:RoleID;references:ID;comment:用户角色"` // 关联 Role 表
	Roles       []Role `gorm:"many2many:sys_user_roles;" json:"roles"`
	OtpEnabled  bool   `json:"otpEnabled" gorm:"default:false;"`
	OtpVerified bool   `json:"otpVerified" gorm:"default:false;"`
	OtpSecret   string `json:"-"`
	OtpAuthUrl  string `json:"-"`
	Status      int    `json:"status"   ` // 1:enable 2:disable
	Note        string `json:"note"     ` //
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
