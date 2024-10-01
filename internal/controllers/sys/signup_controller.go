package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	sysMod "github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
)

// SignupController 注册控制器
type SignupController struct {
	controllers.BaseAPIController
}

func (sc *SignupController) SignupUsingUserName(c *gin.Context) {

	// 1. 验证表单
	request := sys.SignupRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	// 2. 组装数据
	userModel := &sysMod.User{
		Username: request.Name,
		Password: request.Password,
		RoleID:   uint64(request.RoleId),
	}
	// 3. 调用service层
	data, errMsg := service.Entrance.SysService.SignupService.Signup(*userModel)
	if errMsg != "" {
		response.Abort500(c, errMsg)
		return
	}
	response.OK(c, data)
	return
}
