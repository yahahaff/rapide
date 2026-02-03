package traefik

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	"rapide/internal/service"
	"rapide/pkg/logger"
	"rapide/pkg/response"
)

// TraefikController 处理Traefik相关请求

type TraefikController struct {
	controllers.BaseAPIController
}

// GetCRResources 获取特定CRD下的所有自定义资源，支持分页
func (ctrl *TraefikController) GetCRResources(c *gin.Context) {
	// 获取查询参数
	group := c.Query("group")
	version := c.Query("version")
	kind := c.Query("kind")

	// 验证参数
	if group == "" || version == "" || kind == "" {
		response.Abort400(c, "group、version和kind参数不能为空")
		return
	}

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

	// 调用服务层获取资源列表
	resources, total, err := service.Entrance.TraefikService.GetCRResources(group, version, kind, page, pageSize)
	if err != nil {
		logger.Error("Failed to get CR resources: " + err.Error())
		response.Abort500(c, "获取自定义资源失败")
		return
	}

	// 构造符合要求的响应格式
	responseData := gin.H{
		"page":     page,
		"pageSize": pageSize,
		"result":   resources,
		"total":    total,
	}

	response.OK(c, responseData)
}
