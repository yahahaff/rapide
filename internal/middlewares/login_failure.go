// Package middlewares Gin 中间件
// 登录失败2次 强制使用验证码
package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/pkg/redis"
	"io"
	"strconv"
	"time"
)

func LoginFailureAdd(loginId string) {
	key := "rapide:" + "login_fail_count:" + loginId
	err := redis.Redis.Incr(key)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	//设置TTL
	redis.Redis.Expire(key, time.Minute*10)
}

func LoginFailureCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用 Redis 统计密码输错次数
		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}

		// 字符串json化
		m := map[string]string{}
		json.Unmarshal(data, &m)
		key := "rapide:login_fail_count:" + m["login_id"]
		count := redis.Redis.Get(key)

		// 密码输错两次后强制使用验证码
		intNum, _ := strconv.Atoi(count)
		if intNum >= 2 && m["captcha_id"] == "" {
			response.Error(c, response.WithCode(1401), response.WithMessage("Verification code is required"))
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // 复写Body,不然c.Next后面的流程 获取不到参数 关键点
		c.Next()
	}
}
