// Package middlewares 存放系统中间件
package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"rapide/pkg/helpers"
	"rapide/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成请求 ID
		requestID := helpers.GenerateRequestID()
		// 设置到请求头和上下文
		c.Header("X-Request-ID", requestID)
		c.Set("request_id", requestID)

		// 获取 response 内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始时间
		start := time.Now()
		c.Next()

		// 开始记录日志的逻辑
		cost := time.Since(start)
		responStatus := c.Writer.Status()

		// 获取用户 ID（从上下文中获取，由 JWT 中间件设置）
		userID, exists := c.Get("current_user_id")
		userIDStr := ""
		if exists {
			userIDStr = fmt.Sprintf("%v", userID)
		}

		logFields := []zap.Field{
			zap.String("request_id", requestID),
			zap.String("user_id", userIDStr),
			zap.Int("status", responStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			// 请求的内容
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))

			// 响应的内容 - 跳过二进制响应（如文件下载）
			contentType := c.Writer.Header().Get("Content-Type")
			if contentType != "application/zip" && contentType != "application/octet-stream" {
				logFields = append(logFields, zap.String("Response Body", w.body.String()))
			}
		}

		if responStatus > 400 && responStatus <= 499 {
			// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
			logger.Warn("HTTP Warning "+cast.ToString(responStatus), logFields...)
		} else if responStatus >= 500 && responStatus <= 599 {
			// 除了内部错误，记录 error
			logger.Error("HTTP Error "+cast.ToString(responStatus), logFields...)
		} else {
			logger.Debug("HTTP AccessLog", logFields...)
		}
	}
}
