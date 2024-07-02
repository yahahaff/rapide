package verifycode

import (
	"github.com/yahahaff/rapide/pkg/app"
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/redis"
	"time"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 实现 verifycode.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) bool {

	ExpireTime := time.Minute * time.Duration(config.GetInt64("VERIFYCODE_EXPIRE_TIME", 10)) // 默认10分钟
	// 本地环境方便调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("VERIFYCODE_DEBUG_EXPIRE_TIME", 30))
	}

	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

// Get 实现 verifycode.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// Verify 实现 verifycode.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
