package sys

import (
	"fmt"

	sysMode "rapide/internal/models/sys"
	"rapide/internal/requests/sys"

	"rapide/pkg/database"
	"rapide/pkg/handleerror"
	"rapide/pkg/hash"
	"rapide/pkg/logger"
)

type SignupService struct{}

// Signup 注册方法 - 使用统一的请求参数
func (ss *SignupService) Signup(req *sys.SignupRequest) (userData sysMode.User, errMsg string) {
	// 1. 检查用户名是否已存在
	if ss.CheckUsernameExists(req.UserName) {
		return sysMode.User{}, "用户名已存在"
	}

	// 2. 检查邮箱是否已存在（如果提供了邮箱）
	if req.Email != "" && ss.CheckEmailExists(req.Email) {
		return sysMode.User{}, "邮箱已被注册"
	}

	// 3. 检查手机号是否已存在（如果提供了手机号）
	if req.Phone != "" && ss.CheckPhoneExists(req.Phone) {
		return sysMode.User{}, "手机号已被注册"
	}

	// 4. 检查部门是否存在（如果提供了部门ID）
	if req.DeptId > 0 {
		var deptCount int64
		if err := database.DB.Model(&sysMode.Dept{}).Where("id = ?", req.DeptId).Count(&deptCount).Error; err != nil {
			logger.ErrorString("signup", "check_dept_error", err.Error())
			return sysMode.User{}, "检查部门信息失败"
		}
		if deptCount == 0 {
			return sysMode.User{}, "指定的部门不存在"
		}
	}

	// 5. 创建用户模型
	// 处理Email字段，转换为指针类型
	var email *string
	if req.Email != "" {
		email = &req.Email
	}

	// 处理Phone字段，转换为指针类型
	var phone *string
	if req.Phone != "" {
		phone = &req.Phone
	}

	userModel := sysMode.User{
		UserName: req.UserName,
		Password: hash.BcryptHash(req.Password),
		RealName: req.RealName,
		Email:    email,
		Phone:    phone,
		Avatar:   req.Avatar,
		Status:   1,            // 默认启用状态
		HomePath: "/dashboard", // 默认首页路径
	}

	// 6. 开始事务
	tx := database.DB.Begin()
	if tx.Error != nil {
		logger.ErrorString("signup", "begin_transaction_error", tx.Error.Error())
		return sysMode.User{}, "系统内部错误"
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			logger.ErrorString("signup", "transaction_panic", "事务发生panic")
			errMsg = "系统内部错误"
		}
	}()

	// 7. 创建用户
	if err := tx.Create(&userModel).Error; err != nil {
		tx.Rollback()
		logger.ErrorString("signup", "create_user_error", err.Error())
		errMsg = handleerror.GormError(err)
		return sysMode.User{}, errMsg
	}

	// 8. 处理用户角色关联
	if req.RoleID > 0 {
		// 检查角色是否存在
		var roleCount int64
		if err := tx.Model(&sysMode.Role{}).Where("id = ?", req.RoleID).Count(&roleCount).Error; err != nil {
			tx.Rollback()
			logger.ErrorString("signup", "check_role_error", err.Error())
			return sysMode.User{}, "检查角色信息失败"
		}

		if roleCount == 0 {
			tx.Rollback()
			return sysMode.User{}, "指定的角色不存在"
		}

		// 创建用户-角色关联
		userRole := sysMode.UserRole{
			UserID: userModel.ID,
			RoleID: req.RoleID,
		}
		if err := tx.Create(&userRole).Error; err != nil {
			tx.Rollback()
			logger.ErrorString("signup", "create_user_role_error", err.Error())
			errMsg = handleerror.GormError(err)
			return sysMode.User{}, errMsg
		}
	} else {
		// 如果没有指定角色，分配默认角色
		defaultRoleID, err := ss.getDefaultRoleID()
		if err != nil {
			tx.Rollback()
			logger.ErrorString("signup", "get_default_role_error", err.Error())
			return sysMode.User{}, "获取默认角色失败"
		}

		if defaultRoleID > 0 {
			userRole := sysMode.UserRole{
				UserID: userModel.ID,
				RoleID: defaultRoleID,
			}
			if err := tx.Create(&userRole).Error; err != nil {
				tx.Rollback()
				logger.ErrorString("signup", "create_default_user_role_error", err.Error())
				return sysMode.User{}, "分配默认角色失败"
			}
		}
	}

	// 9. 处理用户部门关联
	if req.DeptId > 0 {
		// 创建用户-部门关联
		userDept := sysMode.UserDept{
			UserID: userModel.ID,
			DeptID: req.DeptId,
		}
		if err := tx.Create(&userDept).Error; err != nil {
			tx.Rollback()
			logger.ErrorString("signup", "create_user_dept_error", fmt.Sprintf("创建用户部门关联失败: %v", err))
			return sysMode.User{}, "关联部门失败"
		}
	}

	// 10. 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		logger.ErrorString("signup", "commit_transaction_error", err.Error())
		return sysMode.User{}, "提交事务失败"
	}

	// 11. 返回用户数据（不包含密码）
	userModel.Password = ""
	return userModel, ""
}

// getDefaultRoleID 获取默认角色ID
func (ss *SignupService) getDefaultRoleID() (uint64, error) {
	var defaultRole sysMode.Role
	// 尝试查找名称为 "user" 或 "普通用户" 的角色作为默认角色
	err := database.DB.Where("name = ?", "user").Or("name = ?", "普通用户").First(&defaultRole).Error
	if err != nil {
		// 如果找不到默认角色，尝试获取第一个角色
		err = database.DB.First(&defaultRole).Error
		if err != nil {
			return 0, err
		}
	}
	return defaultRole.ID, nil
}

// SignupWithDefaultRole 使用默认角色注册（简化版）
func (ss *SignupService) SignupWithDefaultRole(userName, password, realName string) (data sysMode.User, errMsg string) {
	req := &sys.SignupRequest{
		UserName: userName,
		Password: password,
		RealName: realName,
		RoleID:   0, // 0表示使用默认角色
	}
	return ss.Signup(req)
}

// CheckUsernameExists 检查用户名是否存在
func (ss *SignupService) CheckUsernameExists(username string) bool {
	var count int64
	database.DB.Model(&sysMode.User{}).Where("user_name = ?", username).Count(&count)
	return count > 0
}

// CheckEmailExists 检查邮箱是否存在
func (ss *SignupService) CheckEmailExists(email string) bool {
	var count int64
	database.DB.Model(&sysMode.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// CheckPhoneExists 检查手机号是否存在
func (ss *SignupService) CheckPhoneExists(phone string) bool {
	var count int64
	database.DB.Model(&sysMode.User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
