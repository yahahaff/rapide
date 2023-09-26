package sys

type CasbinAddRequest struct {
	Type   string `json:"type,omitempty" validate:"required,oneof=p"`
	RoleID string `json:"role_id" validate:"required, max=3"`
	Uri    string `json:"uri" validate:"required"`
	Method string `json:"method" validate:"required"`
}
