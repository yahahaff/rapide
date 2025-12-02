package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	sysDao "github.com/yahahaff/rapide/internal/dao/sys"
	"github.com/yahahaff/rapide/internal/models/sys"
	sysReq "github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/pkg/response"
)

type RoleController struct {
	controllers.BaseAPIController
}

func (rc *RoleController) GetRole(c *gin.Context) {
	// 1. 验证表单
	request := sysReq.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 10 // 设置默认值
	}

	// 处理页码参数
	page := request.Page
	if page <= 0 {
		page = 1
	}

	// 获取角色列表和总数
	roles, total := sysDao.GetRoles(page, pageSize)

	// 构造返回数据
	responseData := map[string]interface{}{
		"result":   roles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}

	response.OK(c, responseData)
}

func (rc *RoleController) AddRole(c *gin.Context) {

	// 1. 验证表单
	request := sysReq.RoleAddRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 验证成功，创建数据
	RoleModel := sys.Role{

		RoleName: request.Name,
	}

	RoleModel.Create()
	//response.Success(c)
	if RoleModel.ID > 0 {
		response.OK(c, RoleModel)
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}

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
