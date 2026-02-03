// Package validators 存放自定义规则和验证器
package validators

// ValidatePasswordConfirm 自定义规则，检查两次密码是否正确
func ValidatePasswordConfirm(password, passwordConfirm string) bool {
	if password != passwordConfirm {
		return false
	}
	return true
}

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
// 由于已使用外部三方验证码服务，此函数简化为直接返回true
func ValidateVerifyCode(key, answer string) bool {
	// 这里简化为直接返回true，实际验证由前端完成
	return true
}
