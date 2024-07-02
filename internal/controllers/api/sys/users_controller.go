package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	"github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/file"
	"github.com/yahahaff/rapide/pkg/jwt"
	"github.com/yahahaff/rapide/pkg/logger"
)

type UsersController struct {
	api.BaseAPIController
}

// RefreshToken 刷新 Access Token
// @Summary 刷新 Access Token
// @Security Bearer
// @Schemes sys.LoginByPhoneRequest{}
// @Description
// @Tags 用户中心
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/refresh-token [post]
func (ctrl *UsersController) RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, response.WithMessage("令牌刷新失败"))
	} else {
		data := gin.H{"token": token}
		response.OK(c, data)
	}
}

// CurrentUser 当前登录用户信息
// @Summary 当前登录用户信息
// @Security Bearer
// @Description
// @Tags 用户中心
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/getUserInfo [get]
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := service.Entrance.SysService.AuthService.CurrentUser(c)
	response.OK(c, userModel)
}

// GetUserList 所有用户
// @Summary 用户列表信息
// @Security Bearer
// @Schemes sys.PaginationRequest{}
// @Param sort query string false "sort"
// @Param order query string false "order"
// @Param per_page query int false "per_page"
// @Param page query int false "page"
// @Description
// @Tags 用户中心
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/getUserList [get]
func (ctrl *UsersController) GetUserList(c *gin.Context) {
	request := sys.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	data, pager, err := service.Entrance.SysService.UserService.GetUserList(request.Page, request.PerPage, request.Sort, request.Order)
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

// UpdateProfile 编辑个人资料
// @Summary 编辑个人资料
// @Security Bearer
// @Schemes requests.sys.UserUpdateProfileRequest{}
// @Param data body sys.UserUpdateProfileRequest{} true "body"
// @Description
// @Tags 用户中心
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/updateProfile [PUT]
func (ctrl *UsersController) UpdateProfile(c *gin.Context) {

	request := sys.UserUpdateProfileRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	currentUser := service.Entrance.SysService.AuthService.CurrentUser(c)
	currentUser.Name = request.Name
	currentUser.Introduction = request.Introduction
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.OK(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// UpdateEmail 修改邮箱
// @Summary 修改邮箱
// @Security Bearer
// @Param data body sys.UserUpdateEmailRequest{} true "body"
// @Description
// @Tags 用户中心
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/updateEmail [PUT]
func (ctrl *UsersController) UpdateEmail(c *gin.Context) {

	request := sys.UserUpdateEmailRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	currentUser := service.Entrance.SysService.AuthService.CurrentUser(c)
	currentUser.Email = request.Email
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		// 失败，显示错误提示
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// UpdatePhone 修改手机
// @Summary 修改手机
// @Security Bearer
// @Param data body sys.UserUpdatePhoneRequest{} true "body"
// @Description
// @Tags 用户中心
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/updatePhone [PUT]
func (ctrl *UsersController) UpdatePhone(c *gin.Context) {

	request := sys.UserUpdatePhoneRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	currentUser := service.Entrance.SysService.AuthService.CurrentUser(c)
	currentUser.Phone = request.Phone
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// ResetByEmail 使用Email和验证码重置密码
// @Summary 重置密码接口
// @Schemes sys.ResetByEmailRequest{}
// @Param data body sys.ResetByEmailRequest{} true "body"
// @Description
// @Tags 用户中心
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/password-reset/using-email [post]
func (ctrl *UsersController) ResetByEmail(c *gin.Context) {
	// 1. 验证表单
	request := sys.ResetByEmailRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 调用service 重置密码
	err := service.Entrance.SysService.UserService.ResetByEmail(request.Email, request.Password)

	// 如果错误存在，记录错误日志，并抛出异常
	if err != nil {
		logger.Error("xxx")
		logger.ErrorString("xxx", "fdfdfd", "fdfdfdfd")
		response.Abort500(c, "获取验证码失败")
		return
	}

	response.Success(c)
}

// UpdatePassword 修改密码
// @Summary 修改密码
// @Security Bearer
// @Param data body sys.UserUpdatePasswordRequest{} true "body"
// @Description
// @Tags 用户中心
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/updatePassword [PUT]
func (ctrl *UsersController) UpdatePassword(c *gin.Context) {

	request := sys.UserUpdatePasswordRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	// 验证密码是否一致
	if ok := validators.ValidatePasswordConfirm(request.NewPassword, request.NewPasswordConfirm); !ok {
		response.Abort401(c, "两次输入密码不匹配")
		return
	}
	currentUser := service.Entrance.SysService.AuthService.CurrentUser(c)
	// 验证原始密码是否正确
	_, err := service.Entrance.SysService.AuthService.Attempt(currentUser.Name, request.Password)

	if err != nil {
		// 失败，显示错误提示
		response.Abort401(c, "原密码不正确")
	} else {
		// 更新密码为新密码
		currentUser.Password = request.NewPassword
		currentUser.Save()

		response.Success(c)
	}
}

// UpdateAvatar 修改头像
// @Summary 修改头像
// @Security Bearer
// @Schemes sys.UserUpdateAvatarRequest{}
// @Param file formData file true "file"
// @Description
// @Tags 用户中心
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/updateAvatar [PUT]
func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {

	request := sys.UserUpdateAvatarRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	avatar, err := file.SaveUploadAvatar(c, request.Avatar)
	if err != nil {
		response.Abort500(c, "上传头像失败，请稍后尝试~")
		return
	}

	currentUser := service.Entrance.SysService.AuthService.CurrentUser(c)
	currentUser.Avatar = avatar
	currentUser.Save()

	response.OK(c, currentUser)
}
