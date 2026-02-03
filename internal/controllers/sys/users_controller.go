package sys

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	"rapide/internal/requests/sys"
	"rapide/internal/requests/validators"
	"rapide/internal/service"
	"rapide/pkg/jwt"
	"rapide/pkg/logger"
	"rapide/pkg/response"
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
		logger.ErrorString("user", "GetUserList_error", fmt.Sprintf("获取用户列表失败: %v", err))
		response.Abort500(c, "获取用户列表失败")
		return
	}

	result := gin.H{
		"result": data,
		"pager":  pager,
	}
	response.OK(c, result)
}

// DeleteUser 删除用户
func (ctrl *UsersController) DeleteUser(c *gin.Context) {
	// 从URL获取用户ID
	userID := c.Param("id")
	// 转换为uint64
	var id uint64
	_, err := fmt.Sscanf(userID, "%d", &id)
	if err != nil {
		response.Abort400(c, "无效的用户ID")
		return
	}

	// 调用服务层删除用户
	err = service.Entrance.SysService.UserService.DeleteUser(id)
	if err != nil {
		logger.ErrorString("user", "error", fmt.Sprintf(err.Error()))
		response.Abort500(c, "删除用户失败")
		return
	}

	response.OK(c, gin.H{"message": "用户删除成功"})
}

// GetUserByID 获取用户详情
func (ctrl *UsersController) GetUserByID(c *gin.Context) {
	// 从URL获取用户ID
	userID := c.Param("id")
	// 转换为uint64
	var id uint64
	_, err := fmt.Sscanf(userID, "%d", &id)
	if err != nil {
		response.Abort400(c, "无效的用户ID")
		return
	}

	// 调用服务层获取用户详情
	user, err := service.Entrance.SysService.UserService.GetUserByID(id)
	if err != nil {
		logger.ErrorString("user", "error", fmt.Sprintf("获取用户详情失败: %v", err))
		response.Abort500(c, "获取用户详情失败")
		return
	}

	response.OK(c, user)
}

// UpdateUser 更新用户信息
func (ctrl *UsersController) UpdateUser(c *gin.Context) {
	// 从URL获取用户ID
	userID := c.Param("id")
	// 转换为uint64
	var id uint64
	_, err := fmt.Sscanf(userID, "%d", &id)
	if err != nil {
		response.Abort400(c, "无效的用户ID")
		return
	}

	// 验证请求数据
	request := sys.UserUpdateRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 调用服务层更新用户信息
	err = service.Entrance.SysService.UserService.UpdateUser(id, &request)
	if err != nil {
		logger.ErrorString("user", "error", fmt.Sprintf("更新用户失败: %v", err))
		response.Abort500(c, fmt.Sprintf("更新用户失败: %v", err))
		return
	}

	response.OK(c, gin.H{"message": "更新用户成功"})
}
