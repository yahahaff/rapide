// Package sys
package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	requestsSys "github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	sysModel "github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/response"
)

// SSLCertController SSL证书控制器
type SSLCertController struct {
	controllers.BaseAPIController
}

// GetSSLCertList 获取SSL证书列表
func (ctrl *SSLCertController) GetSSLCertList(c *gin.Context) {
	request := requestsSys.PaginationRequest{}
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

// CreateSSLCert 创建SSL证书
func (ctrl *SSLCertController) CreateSSLCert(c *gin.Context) {
	request := requestsSys.SSLCertCreateRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 转换请求数据为证书模型
	cert := sysModel.SSLCert{
		Domain:           request.Domain,
		CommonName:       request.CommonName,
		Organization:     request.Organization,
		OrganizationUnit: request.OrganizationUnit,
		Country:          request.Country,
		State:            request.State,
		City:             request.City,
		Email:            request.Email,
		Type:             "DV", // Let's Encrypt 只提供 DV 证书
		Provider:         request.Provider,
		ChallengeType:    request.ChallengeType,
		ApplyStatus:      "pending",
		AutoRenew:        request.AutoRenew,
		RenewStatus:      "idle",
		Status:           1, // 默认为启用状态
	}

	// 调用服务层创建证书
	err := service.Entrance.SysService.SSLCertService.CreateSSLCert(cert)
	if err != nil {
		response.Abort500(c, "创建SSL证书失败")
		return
	}

	response.OK(c, gin.H{"message": "SSL证书创建成功，正在申请中"})
}
