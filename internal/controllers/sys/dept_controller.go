package sys

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	"rapide/internal/models/sys"
	sysReq "rapide/internal/requests/sys"
	"rapide/internal/requests/validators"
	"rapide/internal/service"
	"rapide/pkg/response"
)

// DeptController 部门控制器
type DeptController struct {
	controllers.BaseAPIController
}

// GetDeptTree 获取部门树形结构
func (dc *DeptController) GetDeptTree(c *gin.Context) {
	deptTree, err := service.Entrance.SysService.DeptService.GetDeptTree()
	if err != nil {
		response.Abort500(c, "获取部门树形结构失败")
		return
	}
	response.OK(c, deptTree)
}

// CreateDept 创建部门
func (dc *DeptController) CreateDept(c *gin.Context) {
	var request sysReq.DeptCreateRequest
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 将int类型的pid转换为uint64
	var pid uint64 = uint64(request.Pid)

	// 转换请求为部门模型
	dept := sys.Dept{
		Pid:        pid,
		Name:       request.Name,
		Status:     request.Status,
		Remark:     request.Remark,
		CreateTime: time.Now(),
	}

	// 调用服务创建部门
	err := service.Entrance.SysService.DeptService.CreateDept(dept)
	if err != nil {
		response.Abort500(c, "创建部门失败")
		return
	}

	response.OK(c, gin.H{"id": dept.ID})
}

// UpdateDept 更新部门
func (dc *DeptController) UpdateDept(c *gin.Context) {
	var request sysReq.DeptUpdateRequest

	// 1. 检查URL路径中是否有id参数，如果有则转换为int并设置为请求体的id
	if idStr := c.Param("id"); idStr != "" {
		var id int
		if _, err := fmt.Sscan(idStr, &id); err == nil {
			request.ID = id
		}
	}

	// 2. 解析和验证请求体
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 3. 转换ID为uint64
	id := uint64(request.ID)

	// 4. 获取要更新的部门
	dept, err := service.Entrance.SysService.DeptService.GetDeptByID(id)
	if err != nil {
		response.Abort500(c, "部门不存在")
		return
	}

	// 5. 只更新请求中包含的字段
	// 更新部门名称
	if request.Name != "" {
		dept.Name = request.Name
	}
	// 更新部门状态
	dept.Status = request.Status
	// 更新部门备注
	if request.Remark != "" {
		dept.Remark = request.Remark
	}
	// 更新部门父ID
	dept.Pid = uint64(request.Pid)

	// 6. 调用服务更新部门
	err = service.Entrance.SysService.DeptService.UpdateDept(dept)
	if err != nil {
		response.Abort500(c, "更新部门失败")
		return
	}

	response.OK(c, gin.H{"id": request.ID})
}

// DeleteDept 删除部门
func (dc *DeptController) DeleteDept(c *gin.Context) {
	// 直接从URL路径中获取id参数
	idStr := c.Param("id")
	var id int
	if _, err := fmt.Sscan(idStr, &id); err != nil {
		response.Abort500(c, "无效的部门ID")
		return
	}

	// 转换ID为uint64
	deptID := uint64(id)

	// 调用服务删除部门
	err := service.Entrance.SysService.DeptService.DeleteDept(deptID)
	if err != nil {
		response.Abort500(c, "删除部门失败")
		return
	}

	response.OK(c, gin.H{"id": id})
}

// GetDept 获取单个部门信息
func (dc *DeptController) GetDept(c *gin.Context) {
	// 直接从URL路径中获取id参数
	idStr := c.Param("id")
	var id int
	if _, err := fmt.Sscan(idStr, &id); err != nil {
		response.Abort500(c, "无效的部门ID")
		return
	}

	// 转换ID为uint64
	deptID := uint64(id)

	// 获取部门信息
	dept, err := service.Entrance.SysService.DeptService.GetDeptByID(deptID)
	if err != nil {
		response.Abort500(c, "部门不存在")
		return
	}

	response.OK(c, dept)
}
