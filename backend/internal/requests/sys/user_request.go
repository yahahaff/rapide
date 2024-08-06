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
