// Package sys
package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers/sys"
)

// DeptRouter 部门路由
func DeptRouter(Router *gin.RouterGroup) {
	deptGroup := Router.Group("/dept")
	{
		deptController := new(sys.DeptController)
		// 获取部门树形结构
		deptGroup.GET("/tree", deptController.GetDeptTree)
		// 获取部门列表（兼容旧接口）
		deptGroup.GET("/list", deptController.GetDeptTree)
		// 创建部门
		deptGroup.POST("/create", deptController.CreateDept)
		// 更新部门
		deptGroup.PUT("/update", deptController.UpdateDept)
		// 支持从URL路径获取id的更新部门接口
		deptGroup.PUT("/update/:id", deptController.UpdateDept)
		// 删除部门
		deptGroup.DELETE("/delete/:id", deptController.DeleteDept)
		// 获取单个部门
		deptGroup.GET("/:id", deptController.GetDept)
	}
}
