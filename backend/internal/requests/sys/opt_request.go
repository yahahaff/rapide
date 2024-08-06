package sys

type VerifyActivateOtpRequest struct {
	Token string `json:"token" validate:"required"`
}

type GenerateVerifyRequest struct {
	LoginId string `json:"login_id" validate:"required" `
	Token   string `json:"token" validate:"required"`
}
