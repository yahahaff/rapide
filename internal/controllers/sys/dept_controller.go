package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/models/sys"
	sysReq "github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/response"
)

// DeptController 部门控制器
type DeptController struct {
	controllers.BaseAPIController
}

// GetDept 获取部门列表
func (*DeptController) GetDept(c *gin.Context) {
	request := sysReq.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 20 // 设置默认值
	}

	// 处理页码参数，确保页码大于0
	page := request.Page
	if page <= 0 {
		page = 1
	}

	// 查询部门数据
	var depts []*sys.Dept
	db := database.DB.Where("p_code IS NULL").Preload("Children").Preload("Children.Children")

	// 添加分页限制
	if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	// 查询数据
	db.Find(&depts)

	response.OK(c, depts)
}

// AddDept 添加部门
func (*DeptController) AddDept(c *gin.Context) {
	request := sysReq.DeptAddRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 创建部门
	dept := sys.Dept{
		PCode:  request.PCode,
		PCodes: request.PCodes,
		Name:   request.Name,
		Sort:   request.Sort,
		Tips:   request.Tips,
	}

	// 保存到数据库
	dept.Create()

	response.OK(c, dept)
}

// DeleteDept 删除部门
func (*DeptController) DeleteDept(c *gin.Context) {
	request := sysReq.DeptDeleteRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 删除部门
	database.DB.Where("id=?", request.Id).Delete(&sys.Dept{})

	response.OK(c, nil)
}
