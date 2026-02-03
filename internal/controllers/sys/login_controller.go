package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	"rapide/internal/middlewares"
	"rapide/internal/requests/sys"
	"rapide/internal/requests/validators"
	"rapide/internal/service"
	"rapide/pkg/jwt"
	"rapide/pkg/response"
)

// LoginController 用户控制器
type LoginController struct {
	controllers.BaseAPIController
}

// Login 合并的登录接口，支持手机登录和密码登录
func (lc *LoginController) Login(c *gin.Context) {
	// 1. 验证表单
	request := sys.LoginRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 根据请求参数判断登录方式
	if request.Phone != "" && request.VerifyCode != "" {
		// 手机验证码登录
		user, err := service.Entrance.SysService.AuthService.LoginByPhone(request.Phone)
		if err != nil {
			// 失败，显示错误提示
			response.Abort400(c, err.Error())
			return
		}

		if ok := validators.ValidateVerifyCode(request.Phone, request.VerifyCode); !ok {
			response.Abort401(c, "Verification code error")
			return
		}

		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.UserName)

		data := gin.H{
			"accessToken": token,
		}
		response.OK(c, data)
	} else if request.Username != "" && request.Password != "" {
		// 密码登录
		// 已使用外部三方验证码，移除验证码验证逻辑

		// 用户请求参数中不包含 captchaID 和 captchaAnswer
		user, err := service.Entrance.SysService.AuthService.Attempt(request.Username, request.Password)
		if err != nil {
			// 失败，显示错误提示
			middlewares.LoginFailureAdd(request.Username)
			response.Abort401(c, err.Error())
			return
		}

		if user.OtpEnabled {
			// 用户开启二步认证，直接返回 不生成JWT
			data := gin.H{"2FA": true}
			response.OK(c, data)
			return
		}

		// 正常登录流程，生成JWT
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.UserName)

		userData := gin.H{
			"accessToken": token,
		}
		response.OK(c, userData)
	} else {
		// 登录参数不完整
		response.Abort400(c, "Invalid login parameters")
		return
	}
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
		response.Abort400(c, err.Error())
		return

	} else {
		if ok := validators.ValidateVerifyCode(request.Phone, request.VerifyCode); !ok {
			response.Abort401(c, "验证码错误")
			return
		}
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.UserName)

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

	// 已使用外部三方验证码，移除验证码验证逻辑

	// 用户请求参数中不包含 captchaID 和 captchaAnswer
	user, err := service.Entrance.SysService.AuthService.Attempt(request.Username, request.Password)
	if err != nil {
		// 失败，显示错误提示
		middlewares.LoginFailureAdd(request.Username)
		response.Abort401(c, err.Error())
		return
	}

	if user.OtpEnabled {
		// 用户开启二步认证，直接返回 不生成JWT
		data := gin.H{"2FA": true}
		response.OK(c, data)
		return
	}

	// 正常登录流程，生成JWT
	token := jwt.NewJWT().IssueToken(user.GetStringID(), user.UserName)

	userData := gin.H{
		"accessToken": token,
	}

	response.OK(c, userData)
	return

}
