package sys

import (
	"fmt"

	sysDao "github.com/yahahaff/rapide/internal/dao/sys"
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/paginator"
)

type UserService struct{}

var users []sys.User

// GetUserRoleID 获取用户的角色ID
func (us *UserService) GetUserRoleID(userID uint64) (uint64, error) {
	var user sys.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		return 0, result.Error
	}
	roleID := 1
	return uint64(roleID), nil
}

// GetUserList 获取用户列表
func (us *UserService) GetUserList(page int, size int, sort, order string) (data interface{}, pager paginator.Paging, err error) {
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
			"id":         true,
			"user_name":  true,
			"created_at": true,
			"updated_at": true,
		}
		if !allowedSorts[sort] {
			sort = "id" // 默认按ID排序
		}

		// 排序方向验证
		if order != "asc" && order != "desc" {
			order = "asc" // 默认升序
		}
	} else {
		sort = "id"
		order = "asc"
	}

	// 构建查询
	db := database.DB.Model(&sys.User{})

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
	var users []sys.User
	if err := db.Limit(size).Offset(offset).Find(&users).Error; err != nil {
		return nil, paginator.Paging{}, err
	}

	// 设置分页信息
	pager = paginator.Paging{
		CurrentPage: page,
		PerPage:     size,
		TotalCount:  totalCount,
		TotalPage:   totalPages,
	}

	return users, pager, nil
}

// ResetByEmail 重置密码
func (us *UserService) ResetByEmail(Email, Password string) (err error) {
	userModel := sysDao.GetByEmail(Email)
	if userModel.ID == 0 {
		fmt.Println("邮箱不存在")
	} else {
		userModel.Password = Password
		userModel.Save()
		fmt.Println("更新成功")
	}
	return
}
