package sys

import (
	"fmt"
	"rapide/internal/models/sys"
	"rapide/pkg/database"
	"rapide/pkg/paginator"
)

type OperationLogService struct{}

// GetOperationLog 分页获取操作记录
func (*OperationLogService) GetOperationLog(page int, size int, sort, order string, clientIP, method, path string, status int, startTime, endTime string) (data interface{}, pager paginator.Paging, err error) {
	// 参数验证和默认值处理
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20 // 默认每页20条，最大100条
	}
	
	// 排序参数验证
	if sort != "" {
		// 只允许特定字段排序，防止SQL注入
		allowedSorts := map[string]bool{
			"id":        true,
			"client_ip": true,
			"method":    true,
			"path":      true,
			"status":    true,
			"created_at": true,
		}
		if !allowedSorts[sort] {
			sort = "created_at" // 默认按创建时间排序
		}
		
		// 排序方向验证
		if order != "asc" && order != "desc" {
			order = "desc" // 默认降序
		}
	} else {
		sort = "created_at" // 默认按创建时间排序
		order = "desc"      // 默认降序
	}

	// 构建查询
	db := database.DB.Model(&sys.OperationLog{})

	// 应用查询条件
	if clientIP != "" {
		db = db.Where("client_ip LIKE ?", "%"+clientIP+"%")
	}
	if method != "" {
		db = db.Where("method = ?", method)
	}
	if path != "" {
		db = db.Where("path LIKE ?", "%"+path+"%")
	}
	if status > 0 {
		db = db.Where("status = ?", status)
	}
	if startTime != "" {
		db = db.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		db = db.Where("created_at <= ?", endTime)
	}

	// 添加排序
	db = db.Order(fmt.Sprintf("%s %s", sort, order))

	// 获取总记录数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, paginator.Paging{}, err
	}

	// 计算总页数
	totalPages := int(totalCount) / size
	if int(totalCount)%size != 0 {
		totalPages++
	}

	// 限制当前页不超过总页数
	if page > totalPages && totalPages > 0 {
		page = totalPages
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 执行分页查询
	var operationLog []sys.OperationLog
	if err := db.Limit(size).Offset(offset).Find(&operationLog).Error; err != nil {
		return nil, paginator.Paging{}, err
	}

	// 设置分页信息
	pager = paginator.Paging{
		CurrentPage: page,
		PerPage:     size,
		TotalCount:  totalCount,
		TotalPage:   totalPages,
	}

	return operationLog, pager, nil
}
