package sys

import (
	"rapide/internal/models/sys"
)

// DeptService 部门服务
type DeptService struct{}

// CreateDept 创建部门
func (ds *DeptService) CreateDept(dept sys.Dept) error {
	return dept.Create()
}

// GetDeptByID 根据ID获取部门
func (ds *DeptService) GetDeptByID(id uint64) (sys.Dept, error) {
	return sys.GetDeptByID(id)
}

// GetDeptList 获取部门列表
func (ds *DeptService) GetDeptList() ([]sys.Dept, error) {
	return sys.GetDeptList()
}

// UpdateDept 更新部门
func (ds *DeptService) UpdateDept(dept sys.Dept) error {
	return dept.Update()
}

// DeleteDept 删除部门
func (ds *DeptService) DeleteDept(id uint64) error {
	return sys.DeleteDept(id)
}

// BuildDeptTree 构建部门树形结构
func (ds *DeptService) BuildDeptTree(deptList []sys.Dept) []sys.Dept {
	// 创建部门ID到部门的映射
	deptMap := make(map[uint64]*sys.Dept)

	// 首先，将所有部门添加到映射中，并初始化Children切片
	for i := range deptList {
		dept := &deptList[i]
		dept.Children = make([]*sys.Dept, 0)
		deptMap[dept.ID] = dept
	}

	// 然后，构建树形结构
	// 先将所有部门添加到父部门的Children中
	for i := range deptList {
		dept := &deptList[i]
		if dept.Pid != 0 {
			// 非根节点，添加到父部门的Children中
			if parent, exists := deptMap[dept.Pid]; exists {
				parent.Children = append(parent.Children, dept)
			}
		}
	}

	// 最后，收集所有根节点（Pid为0）作为树的顶部
	var tree []sys.Dept
	for i := range deptList {
		dept := deptList[i]
		if dept.Pid == 0 {
			// 根节点，直接添加到tree中
			tree = append(tree, dept)
		}
	}

	return tree
}

// GetDeptTree 获取部门树形结构
func (ds *DeptService) GetDeptTree() ([]sys.Dept, error) {
	deptList, err := sys.GetDeptList()
	if err != nil {
		return nil, err
	}
	return ds.BuildDeptTree(deptList), nil
}
