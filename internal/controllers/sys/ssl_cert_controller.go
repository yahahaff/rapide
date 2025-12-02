// Package sys
package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/response"
)

// SSLCertController SSL证书控制器
type SSLCertController struct {
	controllers.BaseAPIController
}

// GetSSLCertList 获取SSL证书列表
func (ctrl *SSLCertController) GetSSLCertList(c *gin.Context) {
	request := sys.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 10 // 设置默认值
	}

	// 处理页码参数，确保页码大于0
	page := request.Page
	if page <= 0 {
		page = 1
	}

	data, total, err := service.Entrance.SysService.SSLCertService.GetSSLCertList(page, pageSize)
	if err != nil {
		response.Abort500(c, "获取SSL证书列表失败")
		return
	}

	result := gin.H{
		"list":     data,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}
	response.OK(c, result)
}
