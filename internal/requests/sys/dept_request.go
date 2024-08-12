package sys

type DeptAddRequest struct {
	PCode  int    `json:"p_code"  validate:"required"`
	PCodes string `json:"p_codes"  validate:"required"`
	Name   string `json:"name"  validate:"required"`
	Sort   int    `json:"sort" validate:"required"`
	Tips   string `json:"tips"  validate:"required"`
}

type DeptDeleteRequest struct {
	Id int `json:"id" valid:"required"`
}
