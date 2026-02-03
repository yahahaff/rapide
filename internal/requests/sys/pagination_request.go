package sys

type PaginationRequest struct {
	Page     int    `form:"page" json:"page" binding:"omitempty"`
	PageSize int    `form:"page_size" json:"page_size" binding:"omitempty"`
	Sort     string `form:"sort" json:"sort" binding:"omitempty"`
	Order    string `form:"order" json:"order" binding:"omitempty"`
}
