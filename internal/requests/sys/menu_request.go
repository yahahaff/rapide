package sys

type MenuRequest struct {
	ID     int64  `json:"id" validate:"required"`
	Code   string `json:"code" validate:"required"`
	PCode  string `json:"p_code" validate:"required"`
	PCodes string `json:"p_codes" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Icon   string `json:"icon" validate:"required"`
	URL    string `json:"url" validate:"required"`
	Sort   int64  `json:"sort" validate:"required"`
	Levels int64  `json:"levels" validate:"required"`
	IsMenu int    `json:"is_menu" validate:"required"`
	Status int64  `json:"status" validate:"required"`
	Tips   string `json:"tips" validate:"required"`
}
