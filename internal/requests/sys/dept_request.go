package sys

// DeptCreateRequest 创建部门请求
type DeptCreateRequest struct {
	Pid    int    `json:"pid"`
	Name   string `json:"name" binding:"required,max=255"`
	Status int    `json:"status" binding:"omitempty,oneof=0 1"`
	Remark string `json:"remark" binding:"omitempty"`
}

// DeptUpdateRequest 更新部门请求
type DeptUpdateRequest struct {
	ID     int    `json:"id" binding:"required"`
	Pid    int    `json:"pid"`
	Name   string `json:"name"`
	Status int    `json:"status" binding:"omitempty,oneof=0 1"`
	Remark string `json:"remark"`
}

// DeptIDRequest 部门ID请求
type DeptIDRequest struct {
	ID int `json:"id" binding:"required" uri:"id"`
}
