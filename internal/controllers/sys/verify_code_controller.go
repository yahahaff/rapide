package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	"rapide/pkg/response"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	controllers.BaseAPIController
}

// ShowCaptcha 获取图片验证码
// @Summary 获取图片验证码
// @Description
// @Tags 验证码
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/captcha/image [post]
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 已使用外部三方验证码服务，此接口不再提供功能
	response.Abort500(c, "验证码服务已迁移到外部")
}

func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	// 已使用外部三方验证码服务，此接口不再提供功能
	response.Abort500(c, "验证码服务已迁移到外部")
}
