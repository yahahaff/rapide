package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/sys"
)

func CasbinRouter(Router *gin.RouterGroup) {
	{
		//RBAC
		{
			// role
			roleGroup := Router.Group("/role")
			roc := new(sys.RoleController)
			roleGroup.GET("getRole", roc.GetRole)
			roleGroup.GET("list", roc.GetRole) // 添加list路由，指向相同的处理函数
			roleGroup.POST("addRole", roc.AddRole)
			roleGroup.DELETE("deleteRole", roc.DeleteRoleById)

			// dept
			deptGroup := Router.Group("/dept")
			dc := new(sys.DeptController)
			deptGroup.GET("getDept", dc.GetDept)
			deptGroup.GET("list", dc.GetDept) // 添加list路由，指向相同的处理函数
			deptGroup.POST("addDept", dc.AddDept)
			deptGroup.DELETE("deleteDept", dc.DeleteDept)

		}
	}

}
