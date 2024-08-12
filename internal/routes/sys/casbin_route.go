package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/sys"
)

func CasbinRouter(Router *gin.RouterGroup) {
	{
		//RBAC
		{
			// role
			deptGroup := Router.Group("/role")
			roc := new(sys.RoleController)
			deptGroup.GET("getRole", roc.GetRole)
			deptGroup.POST("addRole", roc.AddRole)
			deptGroup.DELETE("deleteRole", roc.DeleteRoleById)

		}
	}

}
