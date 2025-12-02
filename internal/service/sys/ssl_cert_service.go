// Package sys
package sys

import (
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
)

// SSLCertService SSL证书服务
type SSLCertService struct{}

// GetSSLCertList 获取SSL证书列表
func (ss *SSLCertService) GetSSLCertList(page int, size int) (data interface{}, total int64, err error) {
	// 参数验证和默认值处理
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20 // 默认每页20条，最大100条
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 构建查询
	db := database.DB.Model(&sys.SSLCert{})

	// 获取总记录数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 执行分页查询
	var certList []sys.SSLCert
	if err := db.Order("id desc").Limit(size).Offset(offset).Find(&certList).Error; err != nil {
		return nil, 0, err
	}

	return certList, total, nil
}
