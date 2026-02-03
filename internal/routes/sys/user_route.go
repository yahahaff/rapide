package sys

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers/sys"
)

func UserRouter(Router *gin.RouterGroup) {

	usersGroup := Router.Group("/user")

	//用户中心
	{
		uc := new(sys.UsersController)
		//刷新token,需要登录
		usersGroup.POST("/refresh-token", uc.RefreshToken)

		// 获取当前用户
		usersGroup.GET("/info", uc.GetUserInfo)
		// 获取所有用户
		usersGroup.GET("/list", uc.GetUserList)
		// 删除用户
		usersGroup.DELETE("/delete/:id", uc.DeleteUser)
		// 获取单个用户详情
		usersGroup.GET("/:id", uc.GetUserByID)
		// 更新用户信息
		usersGroup.PUT("/update/:id", uc.UpdateUser)
	}

}
