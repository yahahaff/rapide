package sys

import (
	"errors"

	"github.com/gin-gonic/gin"
	sysDao "github.com/yahahaff/rapide/internal/dao/sys"
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/logger"
)

type AuthService struct{}

// Attempt 尝试登录
func (as *AuthService) Attempt(username, password string) (sys.User, error) {
	userModel := sysDao.GetByUsername(username)
	if userModel.ID == 0 {
		return sys.User{}, errors.New("账号不存在")
	}
	if !userModel.ComparePassword(password) {
		return sys.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// ValidateCaptcha 验证码
func (as *AuthService) ValidateCaptcha(captchaID, captchaAnswer string) bool {
	if ok := validators.ValidateCaptcha(captchaID, captchaAnswer); !ok {
		return false
	}
	return true
}

// LoginByPhone 登录指定用户
func (as *AuthService) LoginByPhone(phone string) (sys.User, error) {
	userModel := sysDao.GetByPhone(phone)
	if userModel.ID == 0 {
		return sys.User{}, errors.New("手机号未注册")
	}

	return sys.User{}, nil
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
