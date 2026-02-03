package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

// Apply the token bucket middleware with specified rate and capacity
// r.Use(TokenBucket(0.5, 10))

func TokenBucket(rate float64 /*令牌放入速率*/, capacity int64 /*桶的容量*/) gin.HandlerFunc {
	tokens := int64(0)      // 令牌放入速率
	lastToken := time.Now() //上一次存放令牌的时间
	m := sync.Mutex{}       // 互斥锁

	return func(c *gin.Context) {
		m.Lock()
		defer m.Unlock()
		now := time.Now()
		tokens = tokens + int64(rate*now.Sub(lastToken).Seconds())
		if tokens > capacity {
			tokens = capacity
		}
		if tokens >= 1 {
			tokens--
			lastToken = now
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			c.Abort()
		}
	}
}
