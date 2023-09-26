// Package requests 处理请求数据和表单验证
package sys

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" validate:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" validate:"email"`
}

// SignupRequest 用户注册
type SignupRequest struct {
	Name            string `json:"name" validate:"required"`
	Phone           string `json:"phone" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required" `
	PasswordConfirm string `json:"password_confirm" validate:"required" `
	VerifyCode      string `json:"verify_code" validate:"required"`
	RoleId          int64  `json:"role_id" validate:"required"`
	DeptId          int64  `json:"dept_id" validate:"required"`
}
