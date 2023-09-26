package cloudflare

type CreateZoneRequest struct {
	AccountID string `json:"id" validate:"required,max=32"`
	Name      string `json:"name" validate:"required,max=253"`
	Type      string `json:"type" validate:"omitempty,oneof=full partial"`
}

type ZoneIDRequest struct {
	ZoneID string `json:"zone_id" validate:"required,max=32"`
}

type EditZoneRequest struct {
	ZoneID string `json:"zone_id" validate:"required,max=32"`
	Paused bool   `json:"paused" validate:"required"`
}
