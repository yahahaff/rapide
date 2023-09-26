package sys

type RoleAddRequest struct {
	Num     int64  `json:"num" validate:"required"`
	Pid     int64  `json:"pid" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Deptid  int64  `json:"deptid" validate:"required"`
	Tips    string `json:"tips" validate:"required"`
	Version int64  `json:"version" validate:"required"`
}

type RoleDeleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type RoleMenuRequest struct {
	RoleID  int64   `json:"role_id" validate:"required"`
	MenuIDs []int64 `json:"menu_ids" validate:"required"`
}
