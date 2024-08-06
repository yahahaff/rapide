package sys

type MenuRequest struct {
	ID     int    `json:"id" validate:"required"`
	Code   string `json:"code" validate:"required"`
	PCode  string `json:"p_code" validate:"required"`
	PCodes string `json:"p_codes" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Icon   string `json:"icon" validate:"required"`
	URL    string `json:"url" validate:"required"`
	Sort   int    `json:"sort" validate:"required"`
	Levels int    `json:"levels" validate:"required"`
	IsMenu int    `json:"is_menu" validate:"required"`
	Status bool   `json:"status" validate:"required"`
	Tips   string `json:"tips" validate:"required"`
}
