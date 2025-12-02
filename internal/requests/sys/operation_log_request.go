package sys

type OperationLogRequest struct {
	PaginationRequest
	ClientIP string `form:"client_ip" json:"client_ip" binding:"omitempty"`
	Method   string `form:"method" json:"method" binding:"omitempty"`
	Path     string `form:"path" json:"path" binding:"omitempty"`
	Status   int    `form:"status" json:"status" binding:"omitempty"`
}
