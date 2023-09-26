// Package validators 存放自定义规则和验证器
package validators

import (
	"github.com/yahahaff/rapide/pkg/captcha"
	"github.com/yahahaff/rapide/pkg/verifycode"
)

// ValidateCaptcha 自定义规则，验证『图片验证码』
func ValidateCaptcha(captchaID, captchaAnswer string) bool {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		return false
	}
	return true
}

// ValidatePasswordConfirm 自定义规则，检查两次密码是否正确
func ValidatePasswordConfirm(password, passwordConfirm string) bool {
	if password != passwordConfirm {
		return false
	}
	return true
}

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
func ValidateVerifyCode(key, answer string) bool {
	if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
		return false
	}
	return true
}
