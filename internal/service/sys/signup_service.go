package sys

import (
	sysMode "github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/internal/requests/sys"

	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/handleerror"
	"github.com/yahahaff/rapide/pkg/hash"
	"github.com/yahahaff/rapide/pkg/logger"
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

	// 4. 创建用户模型
	userModel := sysMode.User{
		UserName: req.UserName,
		Password: hash.BcryptHash(req.Password),
		RealName: req.RealName,
		Email:    req.Email,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
		Status:   1,            // 默认启用状态
		HomePath: "/dashboard", // 默认首页路径
	}

	// 5. 开始事务
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

	// 6. 创建用户
	if err := tx.Create(&userModel).Error; err != nil {
		tx.Rollback()
		logger.ErrorString("signup", "create_user_error", err.Error())
		errMsg = handleerror.GormError(err)
		return sysMode.User{}, errMsg
	}

	// 7. 处理用户角色关联
	if len(req.RoleIDs) > 0 {
		// 检查所有角色是否存在
		var roleCount int64
		if err := tx.Model(&sysMode.Role{}).Where("id IN ?", req.RoleIDs).Count(&roleCount).Error; err != nil {
			tx.Rollback()
			logger.ErrorString("signup", "check_roles_error", err.Error())
			return sysMode.User{}, "检查角色信息失败"
		}

		if int(roleCount) != len(req.RoleIDs) {
			tx.Rollback()
			return sysMode.User{}, "部分角色不存在"
		}

		// 创建用户-角色关联
		userRoles := make([]sysMode.UserRole, 0, len(req.RoleIDs))
		for _, roleID := range req.RoleIDs {
			userRoles = append(userRoles, sysMode.UserRole{
				UserID: userModel.ID,
				RoleID: roleID,
			})
		}

		if err := tx.Create(&userRoles).Error; err != nil {
			tx.Rollback()
			logger.ErrorString("signup", "create_user_roles_error", err.Error())
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

	// 8. 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		logger.ErrorString("signup", "commit_transaction_error", err.Error())
		return sysMode.User{}, "提交事务失败"
	}

	// 9. 返回用户数据（不包含密码）
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
		RoleIDs:  []uint64{}, // 空数组，会自动分配默认角色
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
