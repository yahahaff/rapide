package sys

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	"rapide/internal/requests/sys"
	"rapide/internal/requests/validators"
	"rapide/internal/service"
	"rapide/pkg/logger"
	"rapide/pkg/response"
)

// SignupController 注册控制器
type SignupController struct {
	controllers.BaseAPIController
}

func (sc *SignupController) SignupUsingUserName(c *gin.Context) {
	// 1. 验证表单
	request := sys.SignupRequest{} // 使用统一的请求参数
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// Debug logging
	logger.DebugString("Signup", "Request data", fmt.Sprintf("%+v", request))

	// 2. 验证密码确认
	if request.Password != request.PasswordConfirm {
		response.Abort400(c, "两次输入的密码不一致")
		return
	}

	// 3. 直接调用service层，传递请求参数
	data, errMsg := service.Entrance.SysService.SignupService.Signup(&request)
	if errMsg != "" {
		logger.ErrorString("Signup", "Signup error", errMsg)
		response.Abort500(c, errMsg)
		return
	}

	response.OK(c, data)
}
