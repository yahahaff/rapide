// Package sys
package sys

import (
	"time"

	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/internal/models/sys"
)

// SSLCertService SSL证书服务
type SSLCertService struct{}

// GetSSLCertList 获取SSL证书列表
func (ss *SSLCertService) GetSSLCertList(page int, size int) (data interface{}, total int64, err error) {
	// 解析时间字符串
	createdAt1, _ := time.Parse(time.RFC3339, "2025-01-01T00:00:00Z")
	validityStart1, _ := time.Parse(time.RFC3339, "2025-01-01T00:00:00Z")
	validityEnd1, _ := time.Parse(time.RFC3339, "2026-01-01T00:00:00Z")
	
	createdAt2, _ := time.Parse(time.RFC3339, "2025-02-01T00:00:00Z")
	validityStart2, _ := time.Parse(time.RFC3339, "2025-02-01T00:00:00Z")
	validityEnd2, _ := time.Parse(time.RFC3339, "2026-02-01T00:00:00Z")
	
	createdAt3, _ := time.Parse(time.RFC3339, "2025-03-01T00:00:00Z")
	validityStart3, _ := time.Parse(time.RFC3339, "2025-03-01T00:00:00Z")
	validityEnd3, _ := time.Parse(time.RFC3339, "2026-03-01T00:00:00Z")
	
	// 模拟数据，实际项目中应该从数据库查询
	certList := []sys.SSLCert{
		{
			BaseModel: models.BaseModel{
				ID: 1,
			},
			Domain:           "example.com",
			CommonName:       "example.com",
			Organization:     "Example Company",
			OrganizationUnit: "IT Department",
			Country:          "CN",
			State:            "Beijing",
			City:             "Beijing",
			Email:            "admin@example.com",
			Type:             "DV",
			ValidityStart:    validityStart1,
			ValidityEnd:      validityEnd1,
			Status:           1,
			CommonTimestampsField: models.CommonTimestampsField{
				CreatedAt: createdAt1,
			},
		},
		{
			BaseModel: models.BaseModel{
				ID: 2,
			},
			Domain:           "test.example.com",
			CommonName:       "test.example.com",
			Organization:     "Example Company",
			OrganizationUnit: "Testing Department",
			Country:          "CN",
			State:            "Shanghai",
			City:             "Shanghai",
			Email:            "test@example.com",
			Type:             "OV",
			ValidityStart:    validityStart2,
			ValidityEnd:      validityEnd2,
			Status:           1,
			CommonTimestampsField: models.CommonTimestampsField{
				CreatedAt: createdAt2,
			},
		},
		{
			BaseModel: models.BaseModel{
				ID: 3,
			},
			Domain:           "api.example.com",
			CommonName:       "api.example.com",
			Organization:     "Example Company",
			OrganizationUnit: "API Department",
			Country:          "CN",
			State:            "Guangdong",
			City:             "Shenzhen",
			Email:            "api@example.com",
			Type:             "EV",
			ValidityStart:    validityStart3,
			ValidityEnd:      validityEnd3,
			Status:           1,
			CommonTimestampsField: models.CommonTimestampsField{
				CreatedAt: createdAt3,
			},
		},
	}

	// 计算总数
	total = int64(len(certList))

	// 模拟分页，实际项目中应该使用数据库的LIMIT和OFFSET
	return certList, total, nil
}
