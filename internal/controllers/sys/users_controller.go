package sys

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/jwt"
	"github.com/yahahaff/rapide/pkg/logger"
	"github.com/yahahaff/rapide/pkg/response"
)

type UsersController struct {
	controllers.BaseAPIController
}

func (ctrl *UsersController) RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Abort400(c, "令牌刷新失败")
	} else {
		data := gin.H{"token": token}
		response.OK(c, data)
	}
}

func (ctrl *UsersController) GetUserInfo(c *gin.Context) {
	userModel := service.Entrance.SysService.AuthService.CurrentUser(c)
	// 转换角色为字符串数组

	data := gin.H{
		"id":       userModel.ID,
		"realName": userModel.RealName,
		"roles":    1,
		"userName": userModel.UserName,
	}
	response.OK(c, data)
}

func (ctrl *UsersController) GetUserList(c *gin.Context) {
	request := sys.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 20 // 设置默认值
	}

	// 处理页码参数，确保页码大于0
	page := request.Page
	if page <= 0 {
		page = 1
	}

	data, pager, err := service.Entrance.SysService.UserService.GetUserList(page, pageSize, request.Sort, request.Order)
	// 如果错误存在，记录错误日志，并抛出异常
	if err != nil {
		logger.ErrorString("user", "error", fmt.Sprintf(err.Error()))
		response.Abort500(c, "获取用户列表失败")
		return
	}

	result := gin.H{
		"datalist": data,
		"pager":    pager,
	}
	response.OK(c, result)
}
