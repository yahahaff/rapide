package ssl

type PaginationRequest struct {
	Page        int    `form:"page" json:"page" binding:"omitempty"`
	PageSize    int    `form:"pageSize" json:"pageSize" binding:"omitempty"`
	Sort        string `form:"sort" json:"sort" binding:"omitempty"`
	Order       string `form:"order" json:"order" binding:"omitempty"`
	Domain      string `form:"domain" json:"domain" binding:"omitempty"`
	ApplyStatus string `form:"applyStatus" json:"applyStatus" binding:"omitempty"`
}