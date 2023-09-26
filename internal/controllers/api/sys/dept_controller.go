package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	sysDao "github.com/yahahaff/rapide/internal/dao/sys"
	"github.com/yahahaff/rapide/internal/models/sys"
	sys2 "github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
)

type DeptController struct {
	api.BaseAPIController
}

// GetDept 获取部门列表
// @Summary 获取部门列表
// @Security Bearer
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/dept/getDept [get]
func (dept *DeptController) GetDept(c *gin.Context) {
	data := sysDao.DeptAll()
	response.OK(c, data)

}

// AddDept 新增部门
// @Summary 新增部门
// @Security Bearer
// @Schemes sys.DeptAddRequest{}
// @Param data body sys.DeptAddRequest{} true "body"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/dept/addDept [post]
func (dept *DeptController) AddDept(c *gin.Context) {

	// 1. 验证表单
	request := sys2.DeptAddRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 验证成功，创建数据
	deptModel := sys.Dept{
		Num:      request.Num,
		PID:      request.PID,
		Pids:     request.Pids,
		FullName: request.FullName,
		Tips:     request.Tips,
	}
	deptModel.Create()
	if deptModel.ID > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

// DeleteDeptById 删除部门
// @Summary 删除部门
// @Security Bearer
// @Schemes sys.DeptDeleteRequest{}
// @Param data body sys.DeptDeleteRequest{} true "body"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/dept/deleteDept [delete]
func (dept *DeptController) DeleteDeptById(c *gin.Context) {

	// 1. 验证表单
	request := sys2.DeptDeleteRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 验证成功，删除数据
	sysDao.DeptDeletelById(request.Id)
	response.Success(c)

}
