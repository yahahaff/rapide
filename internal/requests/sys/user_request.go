package sys

import (
	"mime/multipart"
)

type UserUpdateProfileRequest struct {
	Name         string `json:"name" validate:"required,max=15" `
	Introduction string `json:"introduction" validate:"required" `
}

type UserUpdateEmailRequest struct {
	Email      string `json:"email" validate:"required,email,max=254"`
	VerifyCode string `json:"verify_code" validate:"required,max=6"`
}

type UserUpdatePhoneRequest struct {
	Phone      string `json:"phone" validate:"required,max=11"`
	VerifyCode string `json:"verify_code" validate:"required,max=6"`
}

type UserUpdatePasswordRequest struct {
	Password           string `json:"password" validate:"required"`
	NewPassword        string `json:"new_password," validate:"required" `
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required" `
}

// UserUpdateAvatarRequest 修改头像验证器
type UserUpdateAvatarRequest struct {
	Avatar *multipart.FileHeader `json:"avatar" validate:"required" `
}

// UserUpdateRequest 更新用户信息请求结构体
type UserUpdateRequest struct {
	UserName        string  `json:"username" validate:"omitempty,max=15"`
	Phone           string  `json:"phone" validate:"omitempty,phone,max=11"`
	Email           string  `json:"email" validate:"omitempty,email,max=254"`
	Password        string  `json:"password" validate:"omitempty,max=255"`
	PasswordConfirm string  `json:"password_confirm" validate:"omitempty,max=255"`
	RoleID          *uint64 `json:"role_id" validate:"omitempty"`
	RealName        string  `json:"real_name" validate:"omitempty,max=50"`
	NickName        string  `json:"nickname" validate:"omitempty,max=50"`
	Avatar          string  `json:"avatar" validate:"omitempty,url"`
	DeptId          *uint64 `json:"deptId" validate:"omitempty"`
	Status          *int    `json:"status" validate:"omitempty,oneof=0 1"`
	Remark          string  `json:"remark" validate:"omitempty,max=500"`
}
