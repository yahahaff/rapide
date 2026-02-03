package sys

type MenuRequest struct {
	ID     int    `json:"id" validate:"required"`
	Code   string `json:"code" validate:"required"`
	PCode  string `json:"p_code" validate:"required"`
	PCodes string `json:"p_codes" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Icon   string `json:"icon" validate:"required"`
	URL    string `json:"url" validate:"required"`
	Sort   int    `json:"sort" validate:"required"`
	Levels int    `json:"levels" validate:"required"`
	IsMenu int    `json:"is_menu" validate:"required"`
	Status bool   `json:"status" validate:"required"`
	Tips   string `json:"tips" validate:"required"`
}

// MenuCreateRequest 创建菜单请求
type MenuCreateRequest struct {
	Name              string `json:"name" validate:"required"`
	Title             string `json:"title" validate:"required"`
	Path              string `json:"path" validate:"required"`
	Type              string `json:"type" validate:"required"`
	Visible           int    `json:"visible"`
	Icon              string `json:"icon"`
	Component         string `json:"component"`
	Redirect          string `json:"redirect"`
	ParentID          *uint64 `json:"parentId"`
	OrderNo           int    `json:"orderNo"`
	Status            int    `json:"status"`
	KeepAlive         bool   `json:"keepAlive"`
	Hidden            bool   `json:"hidden"`
	HideBreadcrumb    bool   `json:"hideBreadcrumb"`
	HideChildrenInMenu bool   `json:"hideChildrenInMenu"`
	AffixTab          bool   `json:"affixTab"`
	NoBasicLayout     bool   `json:"noBasicLayout"`
	IgnoreKeepAlive   bool   `json:"ignoreKeepAlive"`
	CurrentActiveMenu string `json:"currentActiveMenu"`
}
