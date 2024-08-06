package sys

// 密码重置结构体
type ResetByEmailRequest struct {
	Email      string `json:"email" validate:"required,omitempty"`
	VerifyCode string `json:"verify_code" validate:"required,omitempty"`
	Password   string `validate:"password" json:"required"`
}
