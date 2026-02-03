package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers/sys"
)

func CaptchaRouter(Router *gin.RouterGroup) {
	//验证码路由
	{
		captchaGroup := Router.Group("")
		// 发送验证码
		vcc := new(sys.VerifyCodeController)
		// 图片验证码，需要加限流
		captchaGroup.POST("/image", vcc.ShowCaptcha)
		captchaGroup.POST("/email", vcc.SendUsingEmail)

	}

}
