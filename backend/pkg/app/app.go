package app

import (
	"github.com/yahahaff/rapide/backend/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.GetString("APP_ENV", "debug") == "debug"
}
func IsRelease() bool {
	return config.GetString("APP_ENV", "debug") == "release"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("APP.TIMEZONE", "Asia/Shanghai"))
	return time.Now().In(chinaTimezone)
}
