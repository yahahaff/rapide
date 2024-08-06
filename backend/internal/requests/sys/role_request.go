package sys

import "encoding/json"

type RoleAddRequest struct {
	Name       string          `json:"name"  validate:"required"`
	Value      string          `json:"sort" validate:"required"`
	Desc       string          `json:"tips"  validate:"required"`
	Permission json.RawMessage `json:"permission" `
	Status     int             `json:"status"`
}

type RoleDeleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type RoleMenuRequest struct {
	RoleID  int   `json:"role_id" validate:"required"`
	MenuIDs []int `json:"menu_ids" validate:"required"`
}
