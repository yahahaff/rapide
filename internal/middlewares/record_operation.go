package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/logger"
	"go.uber.org/zap"
	"io"
	"time"
)

func RecordOperation() gin.HandlerFunc {
	recordChan := make(chan *sys.OperationLog, 100) // 缓冲区大小可以根据实际情况调整

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

		// 获取 response 内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		start := time.Now()

		notRecordMethods := map[string]bool{
			"GET": false,
		}

		specificURL := "/api/cloudflare/r2/file/upload"

		// 判断那些请求方法不被记录
		if notRecordMethods[c.Request.Method] && c.Request.URL.Path == specificURL {
			defer func() {
				requests := string(requestBody)
				end := time.Now()
				latency := end.Sub(start) //接口耗时
				path := c.Request.URL.Path
				method := c.Request.Method
				statusCode := c.Writer.Status()
				clientIP := c.ClientIP()
				response := w.body.String()

				OperationLog := &sys.OperationLog{
					ClientIP: clientIP,
					Status:   statusCode,
					Method:   method,
					Path:     path,
					Latency:  latency.Milliseconds(),
					Requests: requests,
					Response: response,
				}
				select {
				case recordChan <- OperationLog:
				default:
					// 如果日志通道已满，则丢弃日志条目
					logger.Error("middlewares", zap.String("error", "record channel full, dropping log entry"))
				}
			}()
		}
		c.Next()
	}
}
