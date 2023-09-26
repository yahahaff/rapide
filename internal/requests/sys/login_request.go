package sys

type LoginByPhoneRequest struct {
	Phone      string `json:"phone" validate:"required,phone"`
	VerifyCode string `json:"verify_code" validate:"required,len=6"`
}

type LoginByPasswordRequest struct {
	LoginID       string `json:"login_id" validate:"required" `
	Password      string `json:"password" validate:"required" `
	CaptchaID     string `json:"captcha_id" validate:"omitempty"`
	CaptchaAnswer string `json:"captcha_answer" validate:"omitempty"`
}
