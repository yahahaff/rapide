package config

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/yahahaff/rapide/pkg/console"
	"github.com/yahahaff/rapide/pkg/helpers"
	"path/filepath"
)

var config *viper.Viper

func InitConfig(ConfigCmd string) {
	config = viper.New()
	config.SetConfigType("yaml")

	// 获取外部参数指定的配置文件路径
	configPath := ConfigCmd
	if configPath != "" {
		// 获取目录路径和文件名
		dir := filepath.Dir(configPath)
		fileName := filepath.Base(configPath)
		// 设置配置路径和配置名称
		config.AddConfigPath(dir)
		config.SetConfigName(fileName)
	} else {
		config.AddConfigPath(".")
		config.SetConfigName("rapide")
	}

	err := config.ReadInConfig()
	if err != nil {
		console.Exit(fmt.Sprintf("Failed to read config file: %v", err.Error()))
	}
}

func internalGet(key string, defaultValue ...interface{}) interface{} {
	// config变量不存在的情况
	if !config.IsSet(key) || helpers.Empty(config.Get(key)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	return config.Get(key)
}

// GetString 获取字符串类型的配置项
func GetString(key string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(key, defaultValue...))
}

// GetInt 获取整数类型的配置项
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))

}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
