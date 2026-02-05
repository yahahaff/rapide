package middlewares

import (
	"bytes"
	"io"
	"time"

	"rapide/internal/models/sys"
	"rapide/pkg/database"
	"rapide/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecordOperation() gin.HandlerFunc {
	recordChan := make(chan *sys.AuditLog, 100) // 缓冲区大小可以根据实际情况调整

	go func() {
		for log := range recordChan {
			database.DB.Create(log)
		}
	}()
	return func(c *gin.Context) {
		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		start := time.Now()

		// 记录所有请求方法
		if true {
			defer func() {
				requests := string(requestBody)
				end := time.Now()
				latency := end.Sub(start) //接口耗时
				path := c.Request.URL.Path
				method := c.Request.Method
				statusCode := c.Writer.Status()
				clientIP := c.ClientIP()

				// 获取操作人信息
				operator := c.GetString("current_user_name")
				// 记录请求的URL Query参数
				query := c.Request.URL.RawQuery
				if query != "" {
					requests = requests + "\nQuery: " + query
				}
				AuditLog := &sys.AuditLog{
					ClientIP: clientIP,
					Status:   statusCode,
					Method:   method,
					Path:     path,
					Latency:  latency.Milliseconds(),
					Requests: requests,
					Operator: operator,
				}
				select {
				case recordChan <- AuditLog:
				default:
					// 如果日志通道已满，则丢弃日志条目
					logger.Error("middlewares", zap.String("error", "record channel full, dropping log entry"))
				}
			}()
		}
		c.Next()
	}
}
