package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/sys"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func CasbinRouter(Router *gin.RouterGroup) {
	Router.Use(middlewares.AuthJWT(), middlewares.PermissionCheck())
	{
		//RBAC
		{
			// Casbin权限认证
			casbinGroup := Router.Group("/permissions")
			cas := new(sys.CabinController)
			casbinGroup.GET("", cas.GetPolicy)
			casbinGroup.POST("", cas.AddCasbin)

		}
		{
			// dept
			deptGroup := Router.Group("/dept")
			dep := new(sys.DeptController)
			deptGroup.GET("getDept", dep.GetDept)
			deptGroup.POST("addDept", dep.AddDept)
			deptGroup.DELETE("deleteDept", dep.DeleteDeptById)
		}

		{
			// role
			deptGroup := Router.Group("/role")
			roc := new(sys.RoleController)
			deptGroup.GET("getRole", roc.GetRole)
			deptGroup.POST("addRole", roc.AddRole)
			deptGroup.DELETE("deleteRole", roc.DeleteRoleById)
			deptGroup.POST("assignRoleMenu", roc.AssignRoleMenu)

		}
	}

}
