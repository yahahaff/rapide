package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/pkg/casbin"
	"net/url"
)

func PermissionCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
		userRoleName := c.GetString("current_user_role_name")

		// 解析URL
		parsedURL, err := url.Parse(c.Request.RequestURI)
		if err != nil {
			fmt.Println("URL解析失败:", err)
			return
		}
		// 获取不带查询参数的URL
		path := parsedURL.Path

		// Casbin校验规则
		if has, err := casbin.Enforcer.Enforce(userRoleName, path, c.Request.Method); err != nil || !has {
			response.Abort403(c)
		} else {
			c.Next()
		}
	}
}
