package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/http"
)

func EtcdRouter(Router *gin.RouterGroup) {

	etcdGroup := Router.Group("/v3/kv")
	etcd := new(http.EtcdController)
	// 获取etcd键值
	etcdGroup.POST("range", etcd.GetKey)
	etcdGroup.POST("put", etcd.PutKey)
	etcdGroup.POST("deleteRange", etcd.DeleteKey)

}
