package validators

import (
	"github.com/go-playground/validator/v10"
	"rapide/pkg/database"
	"net"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 注册自定义验证规则
func RegisterCustomValidation(v *validator.Validate) {
	// 注册自定义规则 not_exists
	v.RegisterValidation("not_exists", validateNotExists)
	// 注册自定义规则 max_cn
	v.RegisterValidation("max_cn", validateMaxCN)
	// 注册自定义规则 min_cn
	v.RegisterValidation("min_cn", validateMinCN)
	// 注册自定义规则 exists
	v.RegisterValidation("exists", validateExists)
	// 注册自定义规则 ipv4
	v.RegisterValidation("ipv4", validateIPv4)
	// 注册自定义规则 phone
	v.RegisterValidation("phone", validatePhone)
}

// 自定义规则 not_exists 验证请求数据必须不存在于数据库中
func validateNotExists(fl validator.FieldLevel) bool {
	rng := strings.Split(strings.TrimPrefix(fl.Param(), "not_exists:"), ",")
	tableName := rng[0]
	dbField := rng[1]
	var exceptID string
	if len(rng) > 2 {
		exceptID = rng[2]
	}

	requestValue := fl.Field().String()

	query := database.DB.Table(tableName).Where(dbField+" = ?", requestValue)
	if len(exceptID) > 0 {
		query = query.Where("id != ?", exceptID)
	}

	var count int64
	query.Count(&count)

	return count == 0
}

// 自定义规则 max_cn 中文长度不超过指定长度
func validateMaxCN(fl validator.FieldLevel) bool {
	valLength := utf8.RuneCountInString(fl.Field().String())
	l, _ := strconv.Atoi(strings.TrimPrefix(fl.Param(), "max_cn:"))
	return valLength <= l
}

// 自定义规则 min_cn 中文长度不小于指定长度
func validateMinCN(fl validator.FieldLevel) bool {
	valLength := utf8.RuneCountInString(fl.Field().String())
	l, _ := strconv.Atoi(strings.TrimPrefix(fl.Param(), "min_cn:"))
	return valLength >= l
}

// 自定义规则 exists 确保数据库存在某条数据
func validateExists(fl validator.FieldLevel) bool {
	rng := strings.Split(strings.TrimPrefix(fl.Param(), "exists:"), ",")
	tableName := rng[0]
	dbField := rng[1]

	requestValue := fl.Field().String()

	var count int64
	database.DB.Table(tableName).Where(dbField+" = ?", requestValue).Count(&count)

	return count > 0
}

// 自定义规则 ipv4 验证 IPv4 地址格式是否正确
func validateIPv4(fl validator.FieldLevel) bool {
	ip := net.ParseIP(fl.Field().String())
	return ip != nil && strings.Contains(fl.Field().String(), ".")
}

// 自定义规则 phone 验证手机号码格式是否正确
func validatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// 使用正则表达式验证手机号码格式
	pattern := `^1[3456789]\d{9}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phone)
}
