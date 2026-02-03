package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	"rapide/internal/requests/sys"
	"rapide/internal/requests/validators"
	"rapide/internal/service"
	"rapide/pkg/response"
)

type OperationLogController struct {
	controllers.BaseAPIController
}

// GetOperationLog 分页获取操作记录
// @Summary 获取操作记录
// @Security Bearer
// @Schemes sys.OperationLogRequest{}
// @Param sort query string false "sort"
// @Param order query string false "order"
// @Param page query int false "page"
// @Param client_ip query string false "client_ip"
// @Param method query string false "method"
// @Param path query string false "path"
// @Param status query int false "status"
// @Param page_size query int false "page_size"
// @Description
// @Tags 系统管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/record/getOperationLog [get]
func (*OperationLogController) GetOperationLog(c *gin.Context) {
	request := sys.OperationLogRequest{}
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

	data, pager, err := service.Entrance.SysService.OperationLogService.GetOperationLog(page, pageSize, request.Sort, request.Order, request.ClientIP, request.Method, request.Path, request.Status, request.Operator, request.StartTime, request.EndTime)
	// 如果错误存在，记录错误日志，并抛出异常
	if err != nil {
		response.Abort500(c, "获取操作列表失败")
		return
	}

	result := gin.H{
		"page":     page,
		"pageSize": pageSize,
		"result":   data,
		"total":    pager.TotalCount,
	}
	response.OK(c, result)
}
