// Package ssl
package ssl

import (
	"github.com/gin-gonic/gin"
	"rapide/internal/controllers/ssl"
)

// SSLCertRouter SSL证书路由
func SSLCertRouter(Router *gin.RouterGroup) {
	sslCertController := new(ssl.SSLCertController)
	// 获取SSL证书列表
	Router.GET("/list", sslCertController.GetSSLCertList)
	// 创建SSL证书
	Router.POST("/create", sslCertController.CreateSSLCert)
	// 下载SSL证书
	Router.GET("/download/:id", sslCertController.DownloadSSLCert)
	// 吊销SSL证书
	Router.POST("/revoke/:id", sslCertController.RevokeSSLCert)
	// 获取单个证书详情
	Router.GET("/detail/:id", sslCertController.GetSSLCertDetail)
}
