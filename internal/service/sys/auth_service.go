package sys

import (
	"errors"

	sysDao "rapide/internal/dao/sys"
	"rapide/internal/models/sys"
	"rapide/pkg/database"
	"rapide/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AuthService struct{}

// Attempt 尝试登录
func (as *AuthService) Attempt(username, password string) (sys.User, error) {
	userModel := sysDao.GetByUsername(username)
	if userModel.ID == 0 {
		return sys.User{}, errors.New("账号不存在")
	}
	// 检查用户状态，0:禁用 1:启用
	if userModel.Status == 0 {
		return sys.User{}, errors.New("账号已被封禁")
	}
	if !userModel.ComparePassword(password) {
		return sys.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// LoginByPhone 登录指定用户
func (as *AuthService) LoginByPhone(phone string) (sys.User, error) {
	userModel := sysDao.GetByPhone(phone)
	if userModel.ID == 0 {
		return sys.User{}, errors.New("手机号未注册")
	}
	// 检查用户状态，0:禁用 1:启用
	if userModel.Status == 0 {
		return sys.User{}, errors.New("账号已被封禁")
	}

	return userModel, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func (as *AuthService) CurrentUser(c *gin.Context) sys.User {
	userModel, ok := c.MustGet("current_user").(sys.User)

	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return sys.User{}
	}
	// 预加载角色关联
	database.DB.Preload("Roles").Find(&userModel)
	// db is now a *DB value
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func (as *AuthService) CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}

// GetRolePermissions 获取角色权限
