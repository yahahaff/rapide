package sys

type DeptAddRequest struct {
	Num      int    `json:"num" validate:"required"`
	PID      int    `json:"pid"  validate:"required"`
	Pids     string `json:"pids"  validate:"required"`
	FullName string `json:"fullname"  validate:"required"`
	Tips     string `json:"tips"  validate:"required"`
}

type DeptDeleteRequest struct {
	Id int `json:"id" valid:"required"`
}
