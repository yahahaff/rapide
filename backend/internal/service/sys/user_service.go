package sys

import (
	"fmt"
	sysDao "github.com/yahahaff/rapide/backend/internal/dao/sys"
	"github.com/yahahaff/rapide/backend/internal/models/sys"
	"github.com/yahahaff/rapide/backend/pkg/database"
	"github.com/yahahaff/rapide/backend/pkg/paginator"
)

type UserService struct{}

var users []sys.User

// GetUserList 获取用户列表
func (us *UserService) GetUserList(page int, size int, sort, order string) (data interface{}, pager paginator.Paging, err error) {

	db := database.DB.Preload("Dept").Preload("Role")

	if sort != "" && order != "" {
		db = db.Order(fmt.Sprintf("%s %s", sort, order))
	}

	var users []sys.User
	db.Find(&users)

	// 分页数据
	data, pager = paginator.Paginate(users, page, size)
	return
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
