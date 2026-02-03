package sys

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InternalRouter(Router *gin.RouterGroup) {
	// 健康监测
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
}
