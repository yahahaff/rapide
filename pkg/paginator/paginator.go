// Package paginator 处理分页逻辑
package paginator

import (
	"reflect"
)

// Paging 分页数据
type Paging struct {
	CurrentPage int   // 当前页
	PerPage     int   // 每页条数
	TotalPage   int   // 总页数
	TotalCount  int64 // 总条数
}

func Paginate(data interface{}, page int, size int) (interface{}, Paging) {
	var result interface{}
	var paging Paging

	// 计算总记录数
	total := reflect.ValueOf(data).Len()

	// 计算总页数
	totalPages := total / size
	if total%size != 0 {
		totalPages += 1
	}

	// 计算当前页数
	currentPage := page
	if currentPage > totalPages {
		currentPage = totalPages
	}

	// 计算起始位置和结束位置
	start := (currentPage - 1) * size
	end := currentPage * size
	if end > total {
		end = total
	}

	// 根据起始位置和结束位置获取数据
	if start < end {
		slice := reflect.ValueOf(data).Slice(start, end)
		result = slice.Interface()
	} else {
		result = []interface{}{}
	}

	// 设置分页信息
	paging = Paging{
		CurrentPage: currentPage,
		PerPage:     size,
		TotalCount:  int64(total),
		TotalPage:   totalPages,
	}

	return result, paging
}
