// Package captcha 处理图片验证码逻辑
package captcha

import (
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/redis"
	"sync"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// once 确保 internalCaptcha 对象只初始化一次
var once sync.Once

// internalCaptcha 内部使用的 Captcha 对象
var internalCaptcha *Captcha

// NewCaptcha 单例模式获取
func NewCaptcha() *Captcha {
	once.Do(func() {
		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}

		// 使用全局 Redis 对象，并配置存储 Key 的前缀
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("APP_NAME", "Rapide") + ":captcha:",
		}

		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("CAPTCHA_HEIGHT", 80),       // 宽
			config.GetInt("CAPTCHA_WIDTH", 240),       // 高
			config.GetInt("CAPTCHA_lLENGTH", 6),       // 长度
			config.GetFloat64("CAPTCHA_MAXSKEW", 0.7), // 数字的最大倾斜角度
			config.GetInt("CAPTCHA_DOTCOUNT", 80),     // 图片背景里的混淆点数量
		)

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}

// GenerateCaptcha 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	id, b64s, _, err = c.Base64Captcha.Generate()
	return id, b64s, err
}

// VerifyCaptcha 验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {

	// 方便本地和 API 自动测试
	//if !internal.IsProduction() && id == config.GetString("captcha.testing_key") {
	//	return true
	//}
	// 第三个参数是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, answer, false)
}
