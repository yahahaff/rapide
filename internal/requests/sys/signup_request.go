// Package requests 处理请求数据和表单验证
package sys

type SignupPhoneExistRequest struct {
	Phone string `json:"phone" validate:"omitempty,phone,max=11"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email" validate:"omitempty,email,max=254"`
}

// SignupRequest 用户注册
type SignupRequest struct {
	Name            string `json:"name" validate:"required,max=15"`
	Phone           string `json:"phone" validate:"omitempty,phone,max=11"`
	Email           string `json:"email" validate:"omitempty,email,max=254"`
	Password        string `json:"password" validate:"required,max=255"`
	PasswordConfirm string `json:"password_confirm" validate:"required,max=255"`
	VerifyCode      string `json:"verify_code" validate:"omitempty,max=6"`
	RoleId          int64  `json:"role_id" validate:"required,max=2"`
	DeptId          int64  `json:"dept_id" validate:"required,max=2"`
}
