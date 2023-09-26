// Package internal 应用信息
package app

import (
	"github.com/yahahaff/rapide/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.GetString("app.env") == "local"
}

func IsTesting() bool {
	return config.GetString("app.env") == "testing"
}

func IsProduction() bool {
	return config.GetString("app.env") == "production"
}

func IsOptop() bool {
	return config.GetString("app.env") == "true"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return config.GetString("app.url") + path
}

// V1URL 拼接带 v1 标示 URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}
