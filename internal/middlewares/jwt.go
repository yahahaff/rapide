package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/dao/sys"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/pkg/jwt"
	"github.com/yahahaff/rapide/pkg/logger"
	"go.uber.org/zap"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		logger.Debug("Header", zap.String("Header", fmt.Sprintf("%s", c.Request.Header)))
		claims, err := jwt.NewJWT().ParserToken(c)
		logger.Debug("claims", zap.String("claims", fmt.Sprintf("%v", claims)))
		// JWT 解析失败，有错误发生
		if err != nil {
			response.Abort401(c, "您需要进行身份验证才能访问此资源")
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := sys.GetById(claims.UserID)
		if userModel.ID == 0 {
			response.Abort401(c, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user_role_id", userModel.RoleID)
		c.Set("current_user", userModel)
		c.Next()
	}
}
