package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/backend/internal/controllers/api"
	"github.com/yahahaff/rapide/backend/internal/middlewares"
	"github.com/yahahaff/rapide/backend/internal/requests/sys"
	"github.com/yahahaff/rapide/backend/internal/requests/validators"
	"github.com/yahahaff/rapide/backend/internal/response"
	"github.com/yahahaff/rapide/backend/internal/service"
	"github.com/yahahaff/rapide/backend/pkg/jwt"
)

// LoginController 用户控制器
type LoginController struct {
	api.BaseAPIController
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {

	// 1. 验证表单
	request := sys.LoginByPhoneRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 通过手机号获取用户名是否存在
	user, err := service.Entrance.SysService.AuthService.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, response.WithMessage("账号不存在"))
		return

	} else {
		if ok := validators.ValidateVerifyCode(request.Phone, request.VerifyCode); !ok {
			response.Abort401(c, "验证码错误")
			return
		}
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Username)

		data := gin.H{
			"token": token,
		}
		response.OK(c, data)
	}
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 1. 验证表单
	request := sys.LoginByPasswordRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	if request.CaptchaID != "" && request.CaptchaAnswer != "" {
		// 用户请求参数中包含 captchaID 和 captchaAnswer
		if ok := service.Entrance.SysService.AuthService.ValidateCaptcha(request.CaptchaID, request.CaptchaAnswer); !ok {
			response.Abort401(c)
			return
		}

	}
	// 用户请求参数中不包含 captchaID 和 captchaAnswer
	user, err := service.Entrance.SysService.AuthService.Attempt(request.Username, request.Password)
	if err != nil {
		// 失败，显示错误提示
		middlewares.LoginFailureAdd(request.Username)
		response.Abort401(c, "账号不存在或密码错误")
		return
	}

	if user.OtpEnabled {
		// 用户开启二步认证，直接返回 不生成JWT
		data := gin.H{"2FA": true}
		response.OK(c, data)
		return
	}

	// 正常登录流程，生成JWT
	token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Username)

	userData := gin.H{
		"accessToken":  token,
		"refreshToken": token,
	}

	response.OK(c, userData)
	return

}

func (lc *LoginController) GetRoleCodes(c *gin.Context) {

	var data []string
	data, _ = service.Entrance.SysService.AuthService.GetRolePermissions(1)
	response.OK(c, data)
	return

}
