package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
	"github.com/yahahaff/rapide/internal/controllers/api"
	"github.com/yahahaff/rapide/internal/dao/sys"
	sysReq "github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/jwt"
	"net/http"
	"strconv"
)

type AuthenticatorController struct {
	api.BaseAPIController
}

// Generate generates a 2FA code using the provided secret
// GenerateCode 生成2FA密钥与二维码
// @Summary 生成2FA密钥与二维码
// @Security Bearer
// @Description
// @Tags 2FA
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/authenticator/generate [post]
func (ac *AuthenticatorController) Generate(c *gin.Context) {

	//获取当前用户名
	userModel := service.Entrance.SysService.AuthService.CurrentUser(c)

	//生成密钥
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Rapide",
		AccountName: userModel.Name,
		SecretSize:  15,
	})
	if err != nil {
		panic(err)
	}

	// 2.存储密钥 and 二维码URL
	sys.UpdateOpt(userModel.Name, key.Secret(), key.URL())

	// 3.返回数组 组装
	response := gin.H{
		"base32":       key.Secret(),
		"otp-auth-url": key.URL(),
	}

	c.JSON(http.StatusOK, response)

}

// Verify 验证2FA
// @Summary 验证2FA 用户token和数据库密钥
// @Description
// @Security Bearer
// @Schemes sys.VerifyActivateOtpRequest{}
// @Param data body sys.VerifyActivateOtpRequest{} true "body"
// @Tags 2FA
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/authenticator/verify [post]
func (ac *AuthenticatorController) Verify(c *gin.Context) {

	// 1. 验证表单
	request := sysReq.VerifyActivateOtpRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	//获取当前用户名
	userModel := service.Entrance.SysService.AuthService.CurrentUser(c)

	// 2.获取用户密钥
	secret := sys.GetOtpSecret(userModel.Name)

	// 3.验证用户token 和数据库密钥
	valid := totp.Validate(request.Token, secret.OtpSecret)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "message"})
		return
	}

	// 4.设置用户2FA状态
	sys.SetOptStatus(userModel.Name)

	// 5.生成JWT
	token := jwt.NewJWT().IssueToken(strconv.FormatUint(userModel.ID, 10), userModel.Name)
	// 验证成功 数据组装
	c.JSON(http.StatusOK, gin.H{"otp_verified": true, "token": token})

}

// Validate 验证2FA
// @Summary 验证2FA
// @Description
// @Schemes sys.GenerateVerifyRequest{}
// @Param data body sys.GenerateVerifyRequest{} true "body"
// @Tags 2FA
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/authenticator/validate [post]
func (ac *AuthenticatorController) Validate(c *gin.Context) {

	// 1. 验证表单
	request := sysReq.GenerateVerifyRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	//获取用户
	userModel := sys.GetByMulti(request.LoginId)
	// 验证
	valid := totp.Validate(request.Token, userModel.OtpSecret)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "message"})
		return
	}

	// 5.验证成功 生成JWT
	token := jwt.NewJWT().IssueToken(strconv.FormatUint(userModel.ID, 10), userModel.Name)

	c.JSON(http.StatusOK, gin.H{"otp_verified": true, "token": token})

}

// Disable 关闭2FA
// @Summary 关闭2FA
// @Description
// @Security Bearer
// @Schemes
// @Tags 2FA
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/authenticator/disable [post]
func (ac *AuthenticatorController) Disable(c *gin.Context) {

	//获取当前用户名
	userModel := service.Entrance.SysService.AuthService.CurrentUser(c)

	// 数据库操作
	sys.DisableOpt(userModel.Name)
	// 验证成功 数据组装

	c.JSON(http.StatusOK, gin.H{"otp_disabled": true})

}
