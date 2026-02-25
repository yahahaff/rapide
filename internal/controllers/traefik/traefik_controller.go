package traefik

import (
	"strconv"

	"rapide/internal/controllers"
	"rapide/internal/service"
	"rapide/pkg/logger"
	"rapide/pkg/response"

	"github.com/gin-gonic/gin"
)

// TraefikController 处理Traefik相关请求

type TraefikController struct {
	controllers.BaseAPIController
}

// GetHTTPRoutes 获取traefik命名空间中的HTTPRoutes，支持分页
func (ctrl *TraefikController) GetHTTPRoutes(c *gin.Context) {
	// 获取分页参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "20")

	// 转换为整数
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 20
	}

	// 调用服务层获取HTTPRoutes列表
	httpRoutes, total, err := service.Entrance.TraefikService.GetHTTPRoutes(page, pageSize)
	if err != nil {
		logger.Error("Failed to get HTTPRoutes: " + err.Error())
		response.Abort500(c, "获取HTTPRoutes失败")
		return
	}

	// 构造符合要求的响应格式
	responseData := gin.H{
		"page":     page,
		"pageSize": pageSize,
		"result":   httpRoutes,
		"total":    total,
	}

	response.OK(c, responseData)
}
