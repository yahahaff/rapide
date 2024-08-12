package initialize

import (
	"fmt"
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/redis"
)

// SetupRedis 初始化 Redis
func SetupRedis() {

	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("REDIS_HOST", "localhost"), config.GetString("REDIS_PORT", "6379")),
		config.GetString("REDIS_USERNAME", ""),
		config.GetString("REDIS_PASSWORD", ""),
		config.GetInt("REDIS_DATABASE", 0),
	)
}
