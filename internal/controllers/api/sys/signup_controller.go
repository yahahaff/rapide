// Package sys 处理用户身份认证相关逻辑
package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	sysMod "github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
)

// SignupController 注册控制器
type SignupController struct {
	api.BaseAPIController
}

// 定义接口
type SignupServiceInterface interface {
	Signup(request *sys.SignupRequest)
}

// IsPhoneExist 检查手机号是否已注册
// @Summary 检查手机号是否已注册
// @Schemes sys.SignupPhoneExistRequest{}
// @Description
// @Tags 登录注册
// @Param data body sys.SignupPhoneExistRequest{} true "body"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/signup/phone/exist [post]
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := sys.SignupPhoneExistRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	//  检查数据库并返回响应
	data, err := service.Entrance.SysService.SignupService.IsPhoneExist(request.Phone)
	if err != nil {
		fmt.Println(err.Error())
	}
	result := gin.H{"exist": data}
	response.OK(c, result)
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := sys.SignupEmailExistRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	data, err := service.Entrance.SysService.SignupService.IsEmailExist(request.Email)
	if err != nil {
		fmt.Println(err.Error())
	}
	result := gin.H{"exist": data}
	response.OK(c, result)
}

// SignupUsingUserName 使用用户名进行注册
// @Summary 使用用户名进行注册
// @Schemes sys.SignupRequest{}
// @Description
// @Tags 登录注册
// @Param data body sys.SignupRequest{} true "body"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/signup/using-username [post]
func (sc *SignupController) SignupUsingUserName(c *gin.Context) {

	// 1. 验证表单
	request := sys.SignupRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	// 2. 组装数据
	userModel := &sysMod.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Email:    request.Email,
		Password: request.Password,
		RoleID:   int(request.RoleId),
		DeptID:   int(request.DeptId),
	}
	// 3. 调用service层
	data, err := service.Entrance.SysService.SignupService.Signup(*userModel)
	if err != nil {
		response.Abort500(c, "创建失败")
		return

	}
	response.OK(c, data)
}

// SignupUsingPhone 使用手机和验证码进行注册
// @Summary 使用手机和验证码进行注册
// @Schemes sys.SignupRequest{}
// @Description
// @Tags 登录注册
// @Param data body sys.SignupRequest{} true "body"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/signup/using-phone [post]
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := sys.SignupRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 调用service层
	userModel := &sysMod.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Email:    request.Email,
		Password: request.Password,
		RoleID:   int(request.RoleId),
		DeptID:   int(request.DeptId),
	}

	data, err := service.Entrance.SysService.SignupService.Signup(*userModel)
	if err != nil {
		response.Abort500(c, "创建失败")
		return
	}
	response.OK(c, data)
}

// SignupUsingEmail 使用 Email + 验证码进行注册
// @Summary 使用邮箱进行注册
// @Schemes sys.SignupRequest{}
// @Description
// @Tags 登录注册
// @Param data body sys.SignupRequest{} true "body"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/signup/using-email [post]
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := sys.SignupRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	// 2. 调用service层
	userModel := &sysMod.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Email:    request.Email,
		Password: request.Password,
		RoleID:   int(request.RoleId),
		DeptID:   int(request.DeptId),
	}

	data, err := service.Entrance.SysService.SignupService.Signup(*userModel)
	if err != nil {
		response.Abort500(c, "创建失败")
		return
	}
	response.OK(c, data)
}
