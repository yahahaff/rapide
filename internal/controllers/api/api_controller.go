// Package api Package v1 处理业务逻辑,  控制器 v1
package api

import "github.com/yahahaff/rapide/internal/service"

// BaseAPIController 基础控制器
type BaseAPIController struct{}

// ServiceGroup 初始化Service统一入口 方便API调用
var ServiceGroup = service.ServiceGroup{}
