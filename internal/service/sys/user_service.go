package sys

import (
	"fmt"

	sysDao "rapide/internal/dao/sys"
	"rapide/internal/models/sys"
	requests "rapide/internal/requests/sys"
	"rapide/pkg/database"
	"rapide/pkg/hash"
	"rapide/pkg/paginator"
)

// UserListResponse 用户列表响应结构体，包含角色和部门信息
type UserListResponse struct {
	sys.User
	RoleName string `json:"roleName"`
	DeptName string `json:"deptName"`
}

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

	// 执行分页查询，排除Depts关联以避免自动预加载
	var users []sys.User
	if err := db.Limit(size).Offset(offset).Omit("Depts").Find(&users).Error; err != nil {
		return nil, paginator.Paging{}, err
	}

	// 构建响应数据，包含角色和部门名称
	var responseUsers []UserListResponse
	for _, user := range users {
		// 获取用户角色名称
		roleName, err := us.getUserRoleName(user.ID)
		if err != nil {
			return nil, paginator.Paging{}, err
		}

		// 获取用户部门名称
		deptName, err := us.getUserDeptName(user.ID)
		if err != nil {
			return nil, paginator.Paging{}, err
		}

		// 构建响应结构体
		responseUser := UserListResponse{
			User:     user,
			RoleName: roleName,
			DeptName: deptName,
		}
		responseUsers = append(responseUsers, responseUser)
	}

	// 设置分页信息
	pager = paginator.Paging{
		CurrentPage: page,
		PerPage:     size,
		TotalCount:  totalCount,
		TotalPage:   totalPages,
	}

	return responseUsers, pager, nil
}

// getUserRoleName 获取用户的角色名称（单个）
func (us *UserService) getUserRoleName(userID uint64) (string, error) {
	var roleName string
	err := database.DB.Table("sys_role").
		Joins("JOIN sys_user_role ON sys_role.id = sys_user_role.role_id").
		Where("sys_user_role.user_id = ?", userID).
		Order("sys_role.id ASC").
		Limit(1).
		Pluck("sys_role.role_name", &roleName).Error
	return roleName, err
}

// getUserDeptName 获取用户的部门名称（单个）
func (us *UserService) getUserDeptName(userID uint64) (string, error) {
	var deptName string
	err := database.DB.Table("sys_dept").
		Joins("JOIN sys_user_dept ON sys_dept.id = sys_user_dept.dept_id").
		Where("sys_user_dept.user_id = ?", userID).
		Order("sys_dept.id ASC").
		Limit(1).
		Pluck("sys_dept.name", &deptName).Error
	return deptName, err
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

// DeleteUser 删除用户
func (us *UserService) DeleteUser(userID uint64) (err error) {
	var user sys.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		return result.Error
	}
	// 删除用户
	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}
	// 删除用户角色关联
	if err := database.DB.Where("user_id = ?", userID).Delete(&sys.UserRole{}).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID 获取单个用户详情
func (us *UserService) GetUserByID(userID uint64) (UserListResponse, error) {
	var user sys.User
	// 查询用户基本信息，排除Depts关联以避免自动预加载
	if err := database.DB.Omit("Depts").First(&user, userID).Error; err != nil {
		return UserListResponse{}, err
	}

	// 获取用户角色名称
	roleName, err := us.getUserRoleName(user.ID)
	if err != nil {
		return UserListResponse{}, err
	}

	// 获取用户部门名称
	deptName, err := us.getUserDeptName(user.ID)
	if err != nil {
		return UserListResponse{}, err
	}

	// 构建响应结构体
	responseUser := UserListResponse{
		User:     user,
		RoleName: roleName,
		DeptName: deptName,
	}

	return responseUser, nil
}

// UpdateUser 更新用户信息
func (us *UserService) UpdateUser(userID uint64, req *requests.UserUpdateRequest) error {
	// 开始事务
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 查询用户是否存在
	var user sys.User
	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 检查用户名是否已存在（如果提供了用户名，且排除当前用户）
	if req.UserName != "" {
		var count int64
		if err := tx.Model(&sys.User{}).Where("user_name = ? AND id != ?", req.UserName, userID).Count(&count).Error; err != nil {
			tx.Rollback()
			return err
		}
		if count > 0 {
			tx.Rollback()
			return fmt.Errorf("用户名已存在")
		}
		// 更新用户名
		user.UserName = req.UserName
	}

	// 检查邮箱是否已存在（如果提供了邮箱，且排除当前用户）
	if req.Email != "" {
		var count int64
		if err := tx.Model(&sys.User{}).Where("email = ? AND id != ?", req.Email, userID).Count(&count).Error; err != nil {
			tx.Rollback()
			return err
		}
		if count > 0 {
			tx.Rollback()
			return fmt.Errorf("邮箱已被注册")
		}
		// 更新邮箱
		user.Email = &req.Email
	} else if req.Email == "" {
		// 如果提供了空邮箱，将其设置为 nil
		user.Email = nil
	}

	// 检查手机号是否已存在（如果提供了手机号，且排除当前用户）
	if req.Phone != "" {
		var count int64
		if err := tx.Model(&sys.User{}).Where("phone = ? AND id != ?", req.Phone, userID).Count(&count).Error; err != nil {
			tx.Rollback()
			return err
		}
		if count > 0 {
			tx.Rollback()
			return fmt.Errorf("手机号已被注册")
		}
		// 更新手机号
		user.Phone = &req.Phone
	} else if req.Phone == "" {
		// 如果提供了空手机号，将其设置为 nil
		user.Phone = nil
	}

	// 更新真实姓名（如果提供了）
	if req.RealName != "" {
		user.RealName = req.RealName
	}

	// 更新昵称（如果提供了），映射到 RealName 字段
	if req.NickName != "" {
		user.RealName = req.NickName
	}

	// 更新状态（如果提供了）
	if req.Status != nil {
		user.Status = *req.Status
	}

	// 更新备注（如果提供了）
	if req.Remark != "" {
		user.Remark = req.Remark
	}

	// 如果提供了新密码，更新密码
	if req.Password != "" {
		if req.Password != req.PasswordConfirm {
			tx.Rollback()
			return fmt.Errorf("两次输入的密码不一致")
		}
		user.Password = hash.BcryptHash(req.Password)
	}

	// 更新用户信息
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新用户角色关联（如果提供了角色ID）
	if req.RoleID != nil {
		// 删除现有的用户角色关联
		if err := tx.Where("user_id = ?", userID).Delete(&sys.UserRole{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 创建新的用户角色关联
		userRole := sys.UserRole{
			UserID: userID,
			RoleID: *req.RoleID,
		}
		if err := tx.Create(&userRole).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 更新用户部门关联（如果提供了部门ID）
	if req.DeptId != nil {
		// 删除现有的用户部门关联
		if err := tx.Where("user_id = ?", userID).Delete(&sys.UserDept{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 创建新的用户部门关联
		userDept := sys.UserDept{
			UserID: userID,
			DeptID: *req.DeptId,
		}
		if err := tx.Create(&userDept).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
