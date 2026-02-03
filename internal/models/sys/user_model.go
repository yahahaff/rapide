// Package sys
package sys

import (
	"rapide/internal/models"
	"rapide/pkg/database"
	"rapide/pkg/hash"
)

// User 用户模型 - 适配 vue-vben-admin 5.x
type User struct {
	models.BaseModel
	UserName string  `json:"UserName" gorm:"type:varchar(255);uniqueIndex;not null;comment:'用户名'"`
	RealName string  `json:"realName" gorm:"type:varchar(255);comment:'真实姓名'"`
	Password string  `json:"-" gorm:"type:varchar(255);not null;comment:'密码'"`
	Salt     string  `json:"-" gorm:"type:varchar(255);comment:'密码盐'"`
	Email    *string `json:"email" gorm:"type:varchar(255);uniqueIndex;default:null;comment:'邮箱'"`
	Phone    *string `json:"phone" gorm:"type:varchar(20);uniqueIndex;default:null;comment:'手机号'"`
	Avatar   string  `json:"avatar" gorm:"type:varchar(255);comment:'头像'"`

	// 状态信息
	Status   int    `json:"status" gorm:"default:1;comment:'状态 0:禁用 1:启用'"`
	Remark   string `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
	HomePath string `json:"homePath" gorm:"type:varchar(255);comment:'首页路径'"`

	// 安全认证
	OtpEnabled  bool   `json:"otpEnabled" gorm:"default:false;comment:'是否启用OTP'"`
	OtpVerified bool   `json:"otpVerified" gorm:"default:false;comment:'OTP是否已验证'"`
	OtpSecret   string `json:"-" gorm:"type:varchar(255);comment:'OTP密钥'"`
	OtpAuthUrl  string `json:"-" gorm:"type:varchar(255);comment:'OTP认证URL'"`

	// 其他信息
	LastLoginIp   string `json:"lastLoginIp" gorm:"type:varchar(50);comment:'最后登录IP'"`
	LastLoginTime string `json:"lastLoginTime" gorm:"type:varchar(50);comment:'最后登录时间'"`

	models.CommonTimestampsField

	// 关联关系
	Depts []Dept `json:"depts" gorm:"many2many:sys_user_dept;"`
	Roles []Role `json:"roles" gorm:"many2many:sys_user_role;"`
}

// TableName Set the table name
func (*User) TableName() string {
	return "sys_user"
}

// UserRole 中间表模型
type UserRole struct {
	UserID uint64 `gorm:"primaryKey;autoIncrement:false"`
	RoleID uint64 `gorm:"primaryKey;autoIncrement:false"`
}

func (UserRole) TableName() string {
	return "sys_user_role"
}

// UserDept 用户部门中间表模型
type UserDept struct {
	UserID uint64 `gorm:"primaryKey;autoIncrement:false"`
	DeptID uint64 `gorm:"primaryKey;autoIncrement:false"`
}

func (UserDept) TableName() string {
	return "sys_user_dept"
}

// Create 创建用户
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// Update 更新用户
func (userModel *User) Update() {
	database.DB.Save(&userModel)
}

// Delete 删除用户
func (userModel *User) Delete() {
	database.DB.Delete(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

// GetUserRoleIDs 获取用户的角色ID列表
func (userModel *User) GetUserRoleIDs() ([]uint64, error) {
	var roleIDs []uint64
	err := database.DB.Table("sys_user_role").
		Where("user_id = ?", userModel.ID).
		Pluck("role_id", &roleIDs).Error
	if err != nil {
		return nil, err
	}
	return roleIDs, nil
}

// GetUserRoles 获取用户的完整角色信息
func (userModel *User) GetUserRoles() ([]Role, error) {
	var roles []Role
	err := database.DB.Table("sys_role").
		Joins("JOIN sys_user_role ON sys_role.id = sys_user_role.role_id").
		Where("sys_user_role.user_id = ?", userModel.ID).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// Save 保存用户
func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
