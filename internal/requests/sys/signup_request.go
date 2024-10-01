package sys

// SignupRequest 用户注册
type SignupRequest struct {
	Name            string `json:"name" validate:"required,max=15"`
	Phone           string `json:"phone" validate:"omitempty,phone,max=11"`
	Email           string `json:"email" validate:"omitempty,email,max=254"`
	Password        string `json:"password" validate:"required,max=255"`
	PasswordConfirm string `json:"password_confirm" validate:"required,max=255"`
	VerifyCode      string `json:"verify_code" validate:"omitempty,max=6"`
	RoleId          int    `json:"role_id" validate:"required,max=2"`
}
