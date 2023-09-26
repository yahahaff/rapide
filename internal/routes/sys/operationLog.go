package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/sys"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func OperationLogRouter(Router *gin.RouterGroup) {
	Router.Use(middlewares.AuthJWT(), middlewares.PermissionCheck())

	{
		// OperationLog路由组
		OperationLogGroup := Router.Group("/record")
		olc := new(sys.OperationLogController)
		OperationLogGroup.GET("getOperationLog", olc.GetOperationLog)

	}
}
