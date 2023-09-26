package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	sysDao "github.com/yahahaff/rapide/internal/dao/sys"
	"github.com/yahahaff/rapide/internal/models/sys"
	sysReq "github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
)

type RoleController struct {
	api.BaseAPIController
}

// GetRole 获取角色列表
// @Summary 获取角色列表
// @Security Bearer
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/role/getRole [get]
func (rc *RoleController) GetRole(c *gin.Context) {
	data := sysDao.RoleAll()
	response.OK(c, data)

}

// AddRole 新增角色
// @Summary 新增角色
// @Security Bearer
// @Schemes sys.RoleAddRequest{}
// @Param data body sys.RoleAddRequest{} true "body"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/role/addRole [post]
func (rc *RoleController) AddRole(c *gin.Context) {

	// 1. 验证表单
	request := sysReq.RoleAddRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 验证成功，创建数据
	RoleModel := sys.Role{
		Num:     request.Num,
		Pid:     request.Pid,
		Name:    request.Name,
		Tips:    request.Tips,
		Version: request.Version,
	}
	RoleModel.Create()
	//response.Success(c)
	if RoleModel.ID > 0 {
		response.OK(c, RoleModel)
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}

// DeleteRoleById 删除角色
// @Summary 删除角色
// @Security Bearer
// @Schemes sys.RoleDeleteRequest{}
// @Param data body sys.RoleDeleteRequest{} true "body"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/role/deleteRole [delete]
func (rc *RoleController) DeleteRoleById(c *gin.Context) {

	// 1. 验证表单
	request := sysReq.RoleDeleteRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 验证成功，删除数据
	sysDao.RoleDeletelById(request.Id)
	response.Success(c)

}

// AssignRoleMenu 处理分配角色菜单的请求
// @Summary 处理分配角色菜单的请求
// @Security Bearer
// @Schemes sys.RoleMenuRequest{}
// @Param data body sys.RoleMenuRequest{} true "body"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/role/assignRoleMenu [POST]
func (rc *RoleController) AssignRoleMenu(c *gin.Context) {
	// 1. 验证表单
	request := sysReq.RoleMenuRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 验证角色和菜单并更新角色菜单关联
	if err := service.Entrance.SysService.RoleService.AssignRoleMenuByID(request.RoleID, request.MenuIDs); err != nil {
		response.Error(c, response.WithMessage(err.Error()))
		return
	}
	response.Success(c)
}
