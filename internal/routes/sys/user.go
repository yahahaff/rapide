package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/sys"
	"github.com/yahahaff/rapide/internal/middlewares"
)

func UserRouter(Router *gin.RouterGroup) {

	usersGroup := Router.Group("/user")    //需要认证分组
	usersPubGroup := Router.Group("/user") //不需要JWT认证分组

	//用户中心
	{
		uc := new(sys.UsersController)
		usersGroup.Use(middlewares.AuthJWT(), middlewares.PermissionCheck())

		//刷新token,需要登录
		usersGroup.POST("/refresh-token", middlewares.AuthJWT(), uc.RefreshToken)

		// 获取当前用户
		usersGroup.GET("/getUserInfo", uc.CurrentUser)
		// 获取所有用户
		usersGroup.GET("/getUserList", uc.GetUserList)

		usersGroup.PUT("/updateProfile", uc.UpdateProfile)
		usersGroup.PUT("/updatePhone", uc.UpdatePhone)
		usersGroup.PUT("/updateEmail", uc.UpdateEmail)
		usersGroup.PUT("/updatePassword", uc.UpdatePassword)
		usersGroup.PUT("/updateAvatar", uc.UpdateAvatar)

		// 不需认证的接口  使用 usersPubGroup
		usersPubGroup.POST("/password-reset/using-email", uc.ResetByEmail) // 重置密码

	}

}
