package routes

import (
	"rapide/internal/middlewares"
	"rapide/internal/routes/ssl"
	"rapide/internal/routes/sys"
	"rapide/internal/routes/traefik"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册分支路由
func RegisterAPIRoutes(Router *gin.Engine) {
	// ==================== 不需要 JWT 认证的路由 ====================

	// 1. 内部公开路由
	internalGroup := Router.Group("")
	{
		// 内部路由由sys包提供
		sys.InternalRouter(internalGroup)
	}

	// 3. 认证相关路由 (/api/auth)
	authGroup := Router.Group("/api/auth")
	{
		sys.LoginRouter(authGroup)  // 登录
		sys.SingupRouter(authGroup) // 注册
	}

	// ==================== 需要 JWT 认证的路由 ====================

	// 4. SSL相关路由 (/api/ssl)
	sslGroup := Router.Group("/api/ssl")
	sslGroup.Use(middlewares.AuthJWT()) // JWT认证
	{
		ssl.SSLCertRouter(sslGroup)
	}

	// 5. 系统管理路由 (/api/sys)
	sysGroup := Router.Group("/api/sys")
	sysGroup.Use(middlewares.AuthJWT()) // JWT认证
	{
		sys.MenuRouter(sysGroup)         // 菜单管理
		sys.CasbinRouter(sysGroup)       // 权限管理
		sys.UserRouter(sysGroup)         // 用户管理
		sys.OperationLogRouter(sysGroup) // 操作日志
		sys.RoleRouter(sysGroup)         // 角色管理
		sys.DeptRouter(sysGroup)         // 部门管理
	}

	// 6. Traefik相关路由 (/api/traefik)
	traefikGroup := Router.Group("/api/traefik")
	// traefikGroup.Use(middlewares.AuthJWT()) // JWT认证
	{
		traefik.TraefikRouter(traefikGroup)
	}
}
