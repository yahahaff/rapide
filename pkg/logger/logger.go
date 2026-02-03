// Package logger
package logger

import (
	"encoding/json"
	"fmt"
	"rapide/pkg/app"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 全局 Logger 对象
var Logger *zap.Logger

// 模块日志级别配置
var moduleLogLevels = map[string]zapcore.Level{
	"database": zapcore.InfoLevel,
	"http":     zapcore.DebugLevel,
	"ssl":      zapcore.DebugLevel,
	"auth":     zapcore.InfoLevel,
	"sys":      zapcore.InfoLevel,
}

// InitLogger 日志初始化
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType, level string) {
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)
	logLevel := parseLogLevel(level)

	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)
	Logger = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	zap.ReplaceGlobals(Logger)
}

// GetModuleLogger 获取指定模块的 logger 实例
func GetModuleLogger(moduleName string) *zap.Logger {
	if level, exists := moduleLogLevels[moduleName]; exists {
		// 为模块创建特定级别的 logger
		return Logger.WithOptions(zap.IncreaseLevel(level))
	}
	// 默认使用全局 logger
	return Logger
}

// SetModuleLogLevel 设置模块日志级别
func SetModuleLogLevel(moduleName string, level string) error {
	logLevel := parseLogLevel(level)
	moduleLogLevels[moduleName] = logLevel
	return nil
}

// parseLogLevel 解析日志级别
func parseLogLevel(level string) zapcore.Level {
	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改配置项")
	}
	return logLevel
}

// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if app.IsLocal() {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	// 生产环境使用 JSON 编码器，便于日志聚合和分析
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getLogWriter 获取日志写入介质
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	// 使用 lumberjack 进行日志轮转
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,  // 日志文件名
		MaxSize:    maxSize,   // 轮转之前日志文件的最大大小 (单位: MB)
		MaxBackups: maxBackup, // 保留的旧日志文件最大数量
		MaxAge:     maxAge,    // 保留的旧日志文件的最大天数
		Compress:   compress,  // 是否压缩旧日志文件
		LocalTime:  true,      // 使用本地时间进行日志切割
	}

	// 本地使用console输出，不写入日志文件
	if app.IsLocal() {
		return zapcore.AddSync(os.Stdout)
	}
	
	// 生产环境同时输出到控制台和文件，便于调试
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(lumberjackLogger),
	)
}

// Dump 调试专用，不会中断程序，会在终端打印出 warning 消息。
func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}

// LogIf 记录 error 等级的日志
func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred:", zap.Error(err))
	}
}

// LogWarnIf 记录 warning 等级的日志
func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

// LogInfoIf 记录 info 等级的日志
func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

// Debug 调试日志
func Debug(moduleName string, fields ...zap.Field) {
	Logger.Debug(moduleName, fields...)
}

// Info 告知类日志
func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

// Warn 警告类日志
func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

// Error 错误类日志
func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

// Fatal 记录日志后退出程序
func Fatal(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}

// DebugString 记录 debug 级别的字符串类型日志
func DebugString(moduleName, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

// InfoString 记录 info 级别的字符串类型日志
func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

// WarnString 记录 warn 级别的字符串类型日志
func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

// ErrorString 记录 error 级别的字符串类型日志
func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

// FatalString 记录 fatal 级别的字符串类型日志
func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录 debug 级别的 JSON 类型日志
func DebugJSON(moduleName, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

// InfoJSON 记录 info 级别的 JSON 类型日志
func InfoJSON(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

// WarnJSON 记录 warn 级别的 JSON 类型日志
func WarnJSON(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

// ErrorJSON 记录 error 级别的 JSON 类型日志
func ErrorJSON(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

// FatalJSON 记录 fatal 级别的 JSON 类型日志
func FatalJSON(moduleName, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

// jsonString 将对象编码为 JSON 字符串
func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
