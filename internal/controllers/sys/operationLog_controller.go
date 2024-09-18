package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
)

type OperationLogController struct {
	controllers.BaseAPIController
}

// GetOperationLog 分页获取操作记录
// @Summary 获取操作记录
// @Security Bearer
// @Schemes sys.PaginationRequest{}
// @Param sort query string false "sort"
// @Param order query string false "order"
// @Param per_page query int false "per_page"
// @Param page query int false "page"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/sys/record/getOperationLog [get]
func (*OperationLogController) GetOperationLog(c *gin.Context) {
	request := sys.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	data, pager, err := service.Entrance.SysService.OperationLogService.GetOperationLog(request.Page, request.PerPage, request.Sort, request.Order)
	// 如果错误存在，记录错误日志，并抛出异常
	if err != nil {
		response.Abort500(c, "获取操作列表失败")
		return
	}

	result := gin.H{
		"datalist": data,
		"pager":    pager,
	}
	response.OK(c, result)
}
