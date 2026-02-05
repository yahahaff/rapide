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

// GetCRDList 获取所有CRD列表，支持分页
func (ctrl *TraefikController) GetCRDList(c *gin.Context) {
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

	// 调用服务层获取CRD列表
	crds, total, err := service.Entrance.TraefikService.GetCRDList(page, pageSize)
	if err != nil {
		logger.Error("Failed to get CRD list: " + err.Error())
		response.Abort500(c, "获取CRD列表失败")
		return
	}

	// 构造符合要求的响应格式
	responseData := gin.H{
		"page":     page,
		"pageSize": pageSize,
		"result":   crds,
		"total":    total,
	}

	response.OK(c, responseData)
}

// GetRoutes 获取Traefik路由信息，支持分页
func (ctrl *TraefikController) GetRoutes(c *gin.Context) {
	// 获取分页参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	// 转换为整数
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	// 尝试获取Traefik IngressRoute资源
	resources, total, err := service.Entrance.TraefikService.GetCRResources("traefik.io", "v1alpha1", "IngressRoute", page, pageSize)
	if err != nil {
		// 如果Traefik CRD不存在，尝试获取Kubernetes HTTPRoute资源
		logger.Info("Traefik IngressRoute not found, trying Kubernetes HTTPRoute")
		resources, total, err = service.Entrance.TraefikService.GetCRResources("gateway.networking.k8s.io", "v1", "HTTPRoute", page, pageSize)
		if err != nil {
			logger.Error("Failed to get routes: " + err.Error())
			// 返回空列表而不是错误
			responseData := gin.H{
				"page":     page,
				"pageSize": pageSize,
				"result":   []interface{}{},
				"total":    0,
			}
			response.OK(c, responseData)
			return
		}
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
