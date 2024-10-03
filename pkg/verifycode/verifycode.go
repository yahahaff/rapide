// Package verifycode 用以发送手机验证码和邮箱验证码
package verifycode

import (
	"fmt"
	"github.com/yahahaff/rapide/pkg/app"
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/helpers"
	"github.com/yahahaff/rapide/pkg/logger"
	"github.com/yahahaff/rapide/pkg/mail"
	"github.com/yahahaff/rapide/pkg/redis"
	"go.uber.org/zap"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// NewVerifyCode 单例模式获取
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				// 增加前缀保持数据库整洁，出问题调试时也方便
				KeyPrefix: config.GetString("APP_NAME", "rapide") + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

// SendEmail 发送邮件验证码，调用示例：
//
//	verifycode.NewVerifyCode().SendEmail(request.Email)
func (vc *VerifyCode) SendEmail(email string) error {
	fmt.Printf("发送给:%s\r\n", email)
	// 生成验证码
	code := vc.generateVerifyCode(email)
	// 方便本地和 API 自动测试
	if app.IsLocal() && strings.HasSuffix(email, config.GetString("VERIFYCODE_DEBUG_EMAIL_SUFFIX", "")) {
		return nil
	}

	content := fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code)
	// 发送邮件
	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("MAIL_ADDRESS", ""),
			Name:    config.GetString("MAIL_NAME", ""),
		},
		To:      []string{email},
		Subject: "Email 验证码",
		HTML:    []byte(content),
	})

	return nil
}

// CheckAnswer 检查用户提交的验证码是否正确，key 可以是手机号或者 Email
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {

	logger.DebugString("verifyCode", "检查验证码", fmt.Sprintf("key: %s, answer: %s", key, answer))
	// 方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
	//if !internal.IsProduction() &&
	//	(strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) ||
	//		strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix"))) {
	//	return true
	//}

	return vc.Store.Verify(key, answer, false)
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {

	// 生成随机码
	code := helpers.RandomNumber(config.GetInt("VERIFYCODE_CODE_LENGTH", 6))

	// 为方便开发，本地环境使用固定验证码
	//if internal.IsLocal() {
	//	code = config.GetString("verifycode.debug_code")
	//}

	logger.Debug("验证码", zap.String(key, code))
	// 将验证码及 KEY（邮箱或手机号）存放到 Redis 中并设置过期时间
	vc.Store.Set(key, code)
	return code
}
