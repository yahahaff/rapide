package initialize

import (
	"github.com/go-playground/validator/v10"
	"rapide/internal/requests/validators"
)

var Validate *validator.Validate

// SetupValidators 初始化 Validators
func SetupValidators() {
	// 初始化验证器
	Validate = validator.New()
	// 注册自定义验证规则
	validators.RegisterCustomValidation(Validate)
}
