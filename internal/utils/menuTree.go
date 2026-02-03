package utils

import (
	"path/filepath"
	"sort"
	"strings"

	"rapide/internal/models/sys"
)

// BuildMenuTree 构建符合 Vben Admin 5.x 要求的菜单树
// 注意：该函数不做额外过滤，仅构建菜单树结构，过滤逻辑应在调用前完成
func BuildMenuTree(menus []*sys.Menu) []map[string]interface{} {
	if len(menus) == 0 {
		return nil
	}

	// 1. 去重：只过滤nil值，保留所有菜单数据
	menuMap := make(map[uint64]*sys.Menu)
	for _, m := range menus {
		if m == nil {
			continue
		}
		menuMap[m.ID] = m
	}

	if len(menuMap) == 0 {
		return nil
	}

	// 2. 构建 Children 关系（在副本上操作，避免污染原始数据）
	validMenus := make([]*sys.Menu, 0, len(menuMap))
	for _, m := range menuMap {
		// 清空 Children（防止多次调用污染）
		m.Children = nil
		validMenus = append(validMenus, m)
	}

	// 建立父子关系
	for _, child := range validMenus {
		if child.ParentID != nil && *child.ParentID != 0 {
			if parent, ok := menuMap[*child.ParentID]; ok {
				parent.Children = append(parent.Children, child)
			}
		}
	}

	// 3. 找出根节点（ParentID == nil 或 0）
	var roots []*sys.Menu
	for _, m := range validMenus {
		isRoot := (m.ParentID == nil) || (*m.ParentID == 0)
		if isRoot {
			roots = append(roots, m)
		}
	}

	// 4. 排序并转换
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].OrderNo < roots[j].OrderNo
	})

	var result []map[string]interface{}
	for _, root := range roots {
		result = append(result, convertMenuToVben(root))
	}

	return result
}

// convertMenuToVben 递归转换单个菜单项
func convertMenuToVben(menu *sys.Menu) map[string]interface{} {
	item := map[string]interface{}{
		"id":     menu.ID,
		"name":   menu.Name,
		"path":   menu.Path,
		"status": menu.Status,
		"type":   menu.Type,
	}

	// 处理 redirect（非空才加）
	if menu.Redirect != "" {
		item["redirect"] = menu.Redirect
	}

	// 处理 component（关键：转换路径格式）
	if menu.Component != "" {
		// 示例: "views/dashboard/analytics/index.vue" → "dashboard/analytics/index"
		comp := strings.TrimPrefix(menu.Component, "views/")
		comp = strings.TrimSuffix(comp, ".vue")
		// 兼容 Windows 路径分隔符（虽然一般不会出现）
		comp = filepath.ToSlash(comp)
		item["component"] = comp
	}

	// 构建 meta（只包含非默认/有意义的字段）
	meta := make(map[string]interface{})

	// Title 必须有（前端用作显示）
	if menu.Title != "" {
		meta["title"] = menu.Title
	}
	// Icon 图标（不为空就返回）
	if menu.Icon != "" {
		meta["icon"] = menu.Icon
	}
	// Order（排序）
	if menu.OrderNo != 0 {
		meta["order"] = menu.OrderNo
	}

	// 布尔型字段：仅当为 true 时才加入（减少 payload）
	if menu.AffixTab {
		meta["affixTab"] = true
	}
	if menu.Hidden {
		meta["hidden"] = true
	}
	if menu.NoBasicLayout {
		meta["noBasicLayout"] = true
	}
	if menu.IgnoreKeepAlive {
		meta["ignoreKeepAlive"] = true
	}
	if menu.HideBreadcrumb {
		meta["hideBreadcrumb"] = true
	}
	if menu.HideChildrenInMenu {
		meta["hideChildrenInMenu"] = true
	}
	if menu.KeepAlive != true { // 默认是 true，所以只记录 false 的情况？但前端默认 keepAlive=true，通常只传 false
		meta["keepAlive"] = menu.KeepAlive
	}

	// CurrentActiveMenu（字符串，非空才加）
	if menu.CurrentActiveMenu != "" {
		meta["currentActiveMenu"] = menu.CurrentActiveMenu
	}

	// 只有 meta 非空才加入 item
	if len(meta) > 0 {
		item["meta"] = meta
	}

	// 递归处理子菜单
	if len(menu.Children) > 0 {
		sort.Slice(menu.Children, func(i, j int) bool {
			return menu.Children[i].OrderNo < menu.Children[j].OrderNo
		})
		children := make([]map[string]interface{}, 0, len(menu.Children))
		for _, child := range menu.Children {
			children = append(children, convertMenuToVben(child))
		}
		item["children"] = children
	}

	return item
}
