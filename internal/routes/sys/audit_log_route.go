package sys

import (
	"rapide/internal/controllers/sys"

	"github.com/gin-gonic/gin"
)

func OperationLogRouter(Router *gin.RouterGroup) {
	// api-logs 路由
	olc := new(sys.OperationLogController)
	Router.GET("/auditLog", olc.GetOperationLog)
}
