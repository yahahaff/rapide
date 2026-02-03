package sys

type LoginByPhoneRequest struct {
	Phone      string `json:"phone" validate:"required,phone,len=11"`
	VerifyCode string `json:"verify_code" validate:"required,len=6"`
}

type LoginByPasswordRequest struct {
	Username      string `json:"username" validate:"required" `
	Password      string `json:"password" validate:"required" `
	CaptchaID     string `json:"captcha_id" validate:"omitempty"`
	CaptchaAnswer string `json:"captcha_answer" validate:"omitempty"`
}

// LoginRequest 合并的登录请求结构体
type LoginRequest struct {
	// 手机登录字段
	Phone      string `json:"phone" validate:"omitempty,phone,len=11"`
	VerifyCode string `json:"verify_code" validate:"omitempty,len=6"`
	
	// 密码登录字段
	Username      string `json:"username" validate:"omitempty" `
	Password      string `json:"password" validate:"omitempty" `
	CaptchaID     string `json:"captcha_id" validate:"omitempty"`
	CaptchaAnswer string `json:"captcha_answer" validate:"omitempty"`
}
