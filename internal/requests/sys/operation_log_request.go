package sys

type OperationLogRequest struct {
	PaginationRequest
	ClientIP  string `form:"clientIp" json:"clientIp" binding:"omitempty"`
	Method    string `form:"method" json:"method" binding:"omitempty"`
	Path      string `form:"path" json:"path" binding:"omitempty"`
	Status    int    `form:"status" json:"status" binding:"omitempty"`
	StartTime string `form:"startTime" json:"startTime" binding:"omitempty"`
	EndTime   string `form:"endTime" json:"endTime" binding:"omitempty"`
}
