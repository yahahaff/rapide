package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/sys"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func AuthRouter(Router *gin.RouterGroup) {

	//登录路由
	{
		loginGroup := Router.Group("/login")
		lgc := new(sys.LoginController)
		// 使用手机号，短信验证码进行登录
		loginGroup.POST("/using-phone", lgc.LoginByPhone)
		// 使用密码登录 支持手机号，Email 和 用户名
		loginGroup.POST("/using-password", middlewares.LoginFailureCheck(), lgc.LoginByPassword)

	}
	//Signup路由
	{
		signupGroup := Router.Group("/signup")
		suc := new(sys.SignupController)
		// 判断手机是否已注册
		signupGroup.POST("/phone/exist", suc.IsPhoneExist)
		// 判断邮箱是否已注册
		signupGroup.POST("/email/exist", suc.IsEmailExist)
		// 使用用户名注册
		signupGroup.POST("/using-username", suc.SignupUsingUserName)
		// 使用手机号注册
		signupGroup.POST("/using-phone", suc.SignupUsingPhone)
		// 使用邮箱注册
		signupGroup.POST("/using-email", suc.SignupUsingEmail)

	}

}
