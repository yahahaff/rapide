package sys

import (
	"fmt"
	"github.com/yahahaff/rapide/backend/internal/models/sys"
	"github.com/yahahaff/rapide/backend/pkg/database"
	"github.com/yahahaff/rapide/backend/pkg/paginator"
)

type OperationLogService struct{}

// GetOperationLog 分页获取操作记录
func (*OperationLogService) GetOperationLog(page int, size int, sort, order string) (data interface{}, pager paginator.Paging, err error) {
	var operationLog []sys.OperationLog
	db := database.DB.Order("id DESC")

	if sort != "" && order != "" {
		db = db.Order(fmt.Sprintf("%s %s", sort, order))
	}
	db.Find(&operationLog)

	// 分页数据
	data, pager = paginator.Paginate(operationLog, page, size)
	return
}
