package sys

type VerifyCodePhoneRequest struct {
	Phone string `json:"phone" validate:"required,phone"`
}

type VerifyCodeEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
