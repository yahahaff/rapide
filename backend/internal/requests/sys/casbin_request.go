package sys

type CasbinAddRequest struct {
	Type     string `json:"type" validate:"omitempty,oneof=p"`
	RoleName string `json:"role_name" validate:"required, max=15"`
	Uri      string `json:"uri" validate:"required"`
	Method   string `json:"method" validate:"required"`
}
