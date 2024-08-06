package initialize

import (
	"github.com/yahahaff/rapide/backend/pkg/config"
	"github.com/yahahaff/rapide/backend/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		config.GetString("LOG_PATH", "../rapide.log"),
		config.GetInt("LOG_MAX_SIZE", 100),
		config.GetInt("LOG_MAX_BACKUP", 10),
		config.GetInt("LOG_MAX_AGE", 30),
		config.GetBool("LOG_COMPRESS", true),
		config.GetString("LOG_TYPE", "json"),
		config.GetString("LOG_LEVEL", "debug"),
	)
}
