package oss

import "github.com/gin-gonic/gin"

func RouterGroup(router *gin.Engine) *gin.RouterGroup {
	oss := router.Group("")
	AliYunOssRouter(oss)
	return oss
}
