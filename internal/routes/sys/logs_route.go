package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/sys"
)

func OperationLogRouter(Router *gin.RouterGroup) {

	{
		// OperationLog路由组
		OperationLogGroup := Router.Group("/record")
		olc := new(sys.OperationLogController)
		OperationLogGroup.GET("getOperationLog", olc.GetOperationLog)

	}
}
