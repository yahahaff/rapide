package middlewares

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/dao/sys"
	"github.com/yahahaff/rapide/pkg/jwt"
	"github.com/yahahaff/rapide/pkg/logger"
	"github.com/yahahaff/rapide/pkg/response"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		logger.DebugString("jwt", "Header", fmt.Sprintf("%s", c.Request.Header))
		claims, err := jwt.NewJWT().ParserToken(c)
		logger.DebugString("jwt", "claims", fmt.Sprintf("%v", claims))
		// JWT 解析失败，有错误发生
		if err != nil {
			response.Abort401(c, "您需要进行身份验证才能访问此资源")
			return
		}

		// 将 claims.UserID 从字符串转换为整数
		userID, err := strconv.Atoi(claims.UserID)
		if err != nil {
			logger.DebugString("jwt", "claims", fmt.Sprintf("%v", err.Error()))
			response.Abort401(c, "无效的用户ID")
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := sys.GetById(userID)
		if userModel.ID == 0 {
			response.Abort401(c, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.UserName)
		c.Set("current_user", userModel)
		c.Set("current_user_role_id", 1)
		c.Next()
	}
}
