package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/pkg/casbin"
	"net/url"
	"strconv"
)

func PermissionCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
		userRoleId := c.GetInt("current_user_role_id")

		// 解析URL
		parsedURL, err := url.Parse(c.Request.RequestURI)
		if err != nil {
			fmt.Println("URL解析失败:", err)
			return
		}
		// 获取不带查询参数的URL
		path := parsedURL.Path

		// Casbin校验规则
		if has, err := casbin.Enforcer.Enforce(strconv.Itoa(userRoleId), path, c.Request.Method); err != nil || !has {
			response.Abort403(c)
		} else {
			c.Next()
		}
	}
}
