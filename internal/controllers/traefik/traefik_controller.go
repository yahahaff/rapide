package traefik

import (
	"strconv"

	"rapide/internal/controllers"
	"rapide/internal/service"
	"rapide/pkg/logger"
	"rapide/pkg/response"

	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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

// GetServices 获取traefik命名空间中的Services，支持分页
func (ctrl *TraefikController) GetServices(c *gin.Context) {
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

	// 调用服务层获取Services列表
	services, total, err := service.Entrance.TraefikService.GetServices(page, pageSize)
	if err != nil {
		logger.Error("Failed to get Services: " + err.Error())
		response.Abort500(c, "获取Services失败")
		return
	}

	// 构造符合要求的响应格式
	responseData := gin.H{
		"page":     page,
		"pageSize": pageSize,
		"result":   services,
		"total":    total,
	}

	response.OK(c, responseData)
}

// GetMiddlewares 获取traefik命名空间中的Middlewares，支持分页
func (ctrl *TraefikController) GetMiddlewares(c *gin.Context) {
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

	// 调用服务层获取Middlewares列表
	middlewares, total, err := service.Entrance.TraefikService.GetMiddlewares(page, pageSize)
	if err != nil {
		logger.Error("Failed to get Middlewares: " + err.Error())
		response.Abort500(c, "获取Middlewares失败")
		return
	}

	// 构造符合要求的响应格式
	responseData := gin.H{
		"page":     page,
		"pageSize": pageSize,
		"result":   middlewares,
		"total":    total,
	}

	response.OK(c, responseData)
}

// CreateHTTPRoute 创建HTTPRoute
func (ctrl *TraefikController) CreateHTTPRoute(c *gin.Context) {
	// 从请求体中获取HTTPRoute的定义
	var httpRoute map[string]interface{}
	if err := c.ShouldBindJSON(&httpRoute); err != nil {
		logger.Error("Failed to bind request body: " + err.Error())
		response.Abort400(c, "请求体格式错误: "+err.Error())
		return
	}

	// 自动补充 apiVersion 和 Kind 字段
	if _, ok := httpRoute["apiVersion"]; !ok {
		httpRoute["apiVersion"] = "gateway.networking.k8s.io/v1"
	}
	if _, ok := httpRoute["kind"]; !ok {
		httpRoute["kind"] = "HTTPRoute"
	}

	// 验证 metadata.name 不为空
	metadata, ok := httpRoute["metadata"].(map[string]interface{})
	if !ok {
		response.Abort400(c, "metadata 字段格式错误")
		return
	}

	name, ok := metadata["name"].(string)
	if !ok || name == "" {
		response.Abort400(c, "metadata.name 不能为空")
		return
	}

	// 获取 namespace，如果没有指定则使用默认值 "traefik"
	namespace, _ := metadata["namespace"].(string)
	if namespace == "" {
		namespace = "traefik"
	}

	// 创建unstructured.Unstructured对象
	unstructuredHTTPRoute := &unstructured.Unstructured{
		Object: httpRoute,
	}

	// 调用服务层创建HTTPRoute
	createdHTTPRoute, err := service.Entrance.TraefikService.CreateHTTPRoute(namespace, unstructuredHTTPRoute)
	if err != nil {
		logger.Error("Failed to create HTTPRoute: " + err.Error())
		response.Abort500(c, "创建HTTPRoute失败: "+err.Error())
		return
	}

	// 返回创建的HTTPRoute
	response.OK(c, createdHTTPRoute)
}

// CreateService 创建Service
func (ctrl *TraefikController) CreateService(c *gin.Context) {
	// 从请求体中获取Service的定义
	var serviceObj map[string]interface{}
	if err := c.ShouldBindJSON(&serviceObj); err != nil {
		logger.Error("Failed to bind request body: " + err.Error())
		response.Abort400(c, "请求体格式错误: " + err.Error())
		return
	}

	// 自动补充 apiVersion 和 Kind 字段
	if _, ok := serviceObj["apiVersion"]; !ok {
		serviceObj["apiVersion"] = "v1"
	}
	if _, ok := serviceObj["kind"]; !ok {
		serviceObj["kind"] = "Service"
	}

	// 验证 metadata.name 不为空
	metadata, ok := serviceObj["metadata"].(map[string]interface{})
	if !ok {
		response.Abort400(c, "metadata 字段格式错误")
		return
	}

	name, ok := metadata["name"].(string)
	if !ok || name == "" {
		response.Abort400(c, "metadata.name 不能为空")
		return
	}

	// 获取 namespace，如果没有指定则使用默认值 "traefik"
	namespace, _ := metadata["namespace"].(string)
	if namespace == "" {
		namespace = "traefik"
	}

	// 创建unstructured.Unstructured对象
	unstructuredService := &unstructured.Unstructured{
		Object: serviceObj,
	}

	// 调用服务层创建Service
	createdService, err := service.Entrance.TraefikService.CreateService(namespace, unstructuredService)
	if err != nil {
		logger.Error("Failed to create Service: " + err.Error())
		response.Abort500(c, "创建Service失败: " + err.Error())
		return
	}

	// 返回创建的Service
	response.OK(c, createdService)
}

// CreateMiddleware 创建Middleware
func (ctrl *TraefikController) CreateMiddleware(c *gin.Context) {
	// 从请求体中获取Middleware的定义
	var middlewareObj map[string]interface{}
	if err := c.ShouldBindJSON(&middlewareObj); err != nil {
		logger.Error("Failed to bind request body: " + err.Error())
		response.Abort400(c, "请求体格式错误: " + err.Error())
		return
	}

	// 自动补充 apiVersion 和 Kind 字段
	if _, ok := middlewareObj["apiVersion"]; !ok {
		middlewareObj["apiVersion"] = "traefik.io/v1alpha1"
	}
	if _, ok := middlewareObj["kind"]; !ok {
		middlewareObj["kind"] = "Middleware"
	}

	// 自动修复 spec 中的字段名称
	if spec, ok := middlewareObj["spec"].(map[string]interface{}); ok {
		// 创建一个映射，存储需要修复的字段名
		fieldMappings := map[string]string{
			"stripprefix":          "stripPrefix",
			"addprefix":            "addPrefix",
			"replacepath":          "replacePath",
			"replacepathregex":     "replacePathRegex",
			"chain":                "chain",
			"circuitbreaker":       "circuitBreaker",
			"compress":             "compress",
			"headers":              "headers",
			"ipwhitelist":          "ipWhiteList",
			"ratelimit":            "rateLimit",
			"redirectregex":        "redirectRegex",
			"retry":                "retry",
			"buffering":            "buffering",
			"errors":               "errors",
			"forwardauth":          "forwardAuth",
			"basicauth":            "basicAuth",
			"digestauth":           "digestAuth",
			"inflightreq":          "inFlightReq",
			"passtlsclientcert":    "passTLSClientCert",
			"plugin":               "plugin",
		}

		// 修复 spec 中的字段名
		for oldField, newField := range fieldMappings {
			if value, exists := spec[oldField]; exists {
				spec[newField] = value
				delete(spec, oldField)
			}
		}

		// 更新 middlewareObj 中的 spec
		middlewareObj["spec"] = spec
	}

	// 验证 metadata.name 不为空
	metadata, ok := middlewareObj["metadata"].(map[string]interface{})
	if !ok {
		response.Abort400(c, "metadata 字段格式错误")
		return
	}

	name, ok := metadata["name"].(string)
	if !ok || name == "" {
		response.Abort400(c, "metadata.name 不能为空")
		return
	}

	// 获取 namespace，如果没有指定则使用默认值 "traefik"
	namespace, _ := metadata["namespace"].(string)
	if namespace == "" {
		namespace = "traefik"
	}

	// 创建unstructured.Unstructured对象
	unstructuredMiddleware := &unstructured.Unstructured{
		Object: middlewareObj,
	}

	// 调用服务层创建Middleware
	createdMiddleware, err := service.Entrance.TraefikService.CreateMiddleware(namespace, unstructuredMiddleware)
	if err != nil {
		logger.Error("Failed to create Middleware: " + err.Error())
		response.Abort500(c, "创建Middleware失败: " + err.Error())
		return
	}

	// 返回创建的Middleware
	response.OK(c, createdMiddleware)
}
