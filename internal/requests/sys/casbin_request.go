package sys

type CasbinAddRequest struct {
	Type   string `json:"type" validate:"omitempty,oneof=p"`
	RoleID string `json:"role_id" validate:"required, max=3"`
	Uri    string `json:"uri" validate:"required"`
	Method string `json:"method" validate:"required"`
}
