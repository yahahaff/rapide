package cloudflare

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/backend/internal/controllers/api/http/cloudflare"
)

func R2Router(Router *gin.RouterGroup) {

	r2Group := Router.Group("/cloudflare/r2")
	r2 := new(cloudflare.R2Controller)
	r2Group.GET("/file/getList", r2.GetR2List)
	r2Group.POST("/file/upload", r2.UploadFile)
	r2Group.DELETE("/file/delete", r2.DeleteFile)
	r2Group.GET("/file/metadata", r2.GetFileMetadata)

}
