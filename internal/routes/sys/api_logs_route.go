package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers/sys"
)

func OperationLogRouter(Router *gin.RouterGroup) {
	// api-logs 路由
	olc := new(sys.OperationLogController)
	Router.GET("/api-logs", olc.GetOperationLog)
}
