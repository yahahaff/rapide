package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/backend/internal/routes/http"
	"github.com/yahahaff/rapide/backend/internal/routes/sys"
)

// RegisterAPIRoutes 注册分支路由
func RegisterAPIRoutes(Router *gin.Engine) {
	sys.RouterGroup(Router)
	http.RouterGroup(Router)

}
