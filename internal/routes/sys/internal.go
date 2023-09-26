package sys

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yahahaff/rapide/docs"
	"github.com/yahahaff/rapide/pkg/app"
	"net/http"
)

func InternalRouter(Router *gin.RouterGroup) {
	// 健康监测
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// 本地环境开启swagger配置
	if app.IsLocal() {
		docs.SwaggerInfo.BasePath = ""
		Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
