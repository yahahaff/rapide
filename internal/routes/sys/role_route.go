package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers/sys"
)

func RoleRouter(Router *gin.RouterGroup) {
	// role
	roleGroup := Router.Group("/role")
	roc := new(sys.RoleController)
	roleGroup.GET("getRole", roc.GetRole)
	roleGroup.GET("list", roc.GetRole) // 添加list路由，指向相同的处理函数
	roleGroup.POST("addRole", roc.AddRole)
	roleGroup.POST("create", roc.AddRole) // 添加create路由，指向相同的处理函数
	roleGroup.DELETE("deleteRole", roc.DeleteRoleById)
	// 添加根据ID删除角色的路由
	roleGroup.DELETE("delete/:id", roc.DeleteRoleById)
	// 添加更新角色路由
	roleGroup.PUT("update/:id", roc.UpdateRole)
	// 添加角色权限路由
	roleGroup.GET("permissions/:id", roc.GetRolePermissions)
	roleGroup.PUT("permissions/:id", roc.UpdateRolePermissions)
}
