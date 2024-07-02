package config

import (
	"github.com/spf13/viper"
)

var v *viper.Viper

// InitConfig initializes the configuration
func InitConfig() {
	v = viper.New()
	v.AutomaticEnv()
}

func GetString(envName, defaultValue string) string {
	if v.IsSet(envName) {
		value := v.GetString(envName)
		return value
	}
	return defaultValue
}

func GetInt(envName string, defaultValue int) int {
	if v.IsSet(envName) {
		value := v.GetInt(envName)
		return value
	}
	return defaultValue
}
func GetInt64(envName string, defaultValue int64) int64 {
	if v.IsSet(envName) {
		value := v.GetInt64(envName)
		return value
	}
	return defaultValue
}

func GetFloat64(envName string, defaultValue float64) float64 {
	if v.IsSet(envName) {
		value := v.GetFloat64(envName)
		return value
	}
	return defaultValue
}

func GetBool(envName string, defaultValue bool) bool {
	if v.IsSet(envName) {
		value := v.GetBool(envName)
		return value
	}
	return defaultValue
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
