package sys

// MenuService 菜单服务
type MenuService struct{}

// NewMenuService 创建菜单服务实例
func NewMenuService() *MenuService {
	return &MenuService{}
}

// GetUserMenus 根据角色ID获取用户菜单树
func (s *MenuService) GetUserMenus(roleIDs []uint64) (any, error) {

	return 1, nil
}
