package sys

// SignupRequest 用户注册
type SignupRequest struct {
	UserName        string `json:"username" validate:"required,max=15"`
	Phone           string `json:"phone" validate:"omitempty,phone,max=11"`
	Email           string `json:"email" validate:"omitempty,email,max=254"`
	Password        string `json:"password" validate:"required,max=255"`
	PasswordConfirm string `json:"password_confirm" validate:"required,max=255"`
	VerifyCode      string `json:"verify_code" validate:"omitempty,max=6"`
	RoleID          uint64 `json:"role_id" validate:"omitempty"`
	RealName        string `json:"real_name" validate:"omitempty,max=50"`
	Avatar          string `json:"avatar" validate:"omitempty,url"`
	DeptId          uint64 `json:"deptId" validate:"omitempty"`
}
