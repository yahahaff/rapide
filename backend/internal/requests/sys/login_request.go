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
