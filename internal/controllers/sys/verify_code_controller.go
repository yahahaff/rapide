package sys

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/pkg/captcha"
	"github.com/yahahaff/rapide/pkg/logger"
	"github.com/yahahaff/rapide/pkg/response"
	"github.com/yahahaff/rapide/pkg/verifycode"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	controllers.BaseAPIController
}

type ErrorResponse struct {
	Error string `json:"error"`
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

	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 如果错误存在，记录错误日志
	if err != nil {
		logger.ErrorString("verify-codes", "error", fmt.Sprintf(err.Error()))
		response.Abort400(c, ("获取验证码失败"))
		return
	}

	// 构造要返回的JSON数据
	data := gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	}
	response.OK(c, data)
}

func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := sys.VerifyCodeEmailRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	// 2. 发送邮件
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "发送 Email 验证码失败~")
	} else {
		response.Success(c)
	}
}
