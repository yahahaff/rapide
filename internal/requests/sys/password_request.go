package sys

// 密码重置结构体
type ResetByEmailRequest struct {
	Email      string `json:"email,omitempty" validate:"required"`
	VerifyCode string `json:"verify_code,omitempty" validate:"required"`
	Password   string `validate:"password" json:"required"`
}
