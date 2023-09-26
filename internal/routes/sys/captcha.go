package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/sys"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func CaptchaRouter(Router *gin.RouterGroup) {
	//验证码路由
	{
		captchaGroup := Router.Group("/captcha")
		// 发送验证码
		vcc := new(sys.VerifyCodeController)
		// 图片验证码，需要加限流
		captchaGroup.POST("/image", vcc.ShowCaptcha)
		captchaGroup.POST("/phone", vcc.SendUsingPhone)
		captchaGroup.POST("/email", vcc.SendUsingEmail)

	}

	// OPT路由
	{
		optGroup := Router.Group("/opt")
		opt := new(sys.OptController)
		optGroup.POST("", middlewares.AuthJWT(), opt.GenerateOTP)
		optGroup.POST("/disableOtp", middlewares.AuthJWT(), opt.DisableOTP)
		//绑定OPT时调用 写入密钥、更新状态等到数据库
		optGroup.POST("/verifyOtp", middlewares.AuthJWT(), opt.VerifyOTP)
		//验证码验证
		optGroup.POST("/Validate", opt.ValidateOTP)
	}
}
