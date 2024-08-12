package sys

type PaginationRequest struct {
	PerPage int    `form:"per_page" json:"per_page" binding:"required"`
	Page    int    `form:"page" json:"page" binding:"required"`
	Sort    string `form:"sort" json:"sort" binding:"omitempty"`
	Order   string `form:"order" json:"order" binding:"omitempty"`
}
