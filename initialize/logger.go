package initialize

import (
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		config.GetString("log.path", "storage/logs/rapide.log"),
		config.GetInt("log.max_size", 100),
		config.GetInt("log.max_backup", 10),
		config.GetInt("log.max_age", 30),
		config.GetBool("log.compress", true),
		config.GetString("log.lot_type", "json"),
		config.GetString("log.level", "info"),
	)
}
