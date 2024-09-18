package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/etcd"
)

func EtcdRouter(Router *gin.RouterGroup) {

	etcdGroup := Router.Group("/v3/kv")
	etcd := new(etcd.EtcdController)
	// 获取etcd键值
	etcdGroup.GET("getKeyList", etcd.GetKeyList)
	etcdGroup.GET("getKeyDetail", etcd.GetKeyDetail)
	etcdGroup.POST("put", etcd.PutKey)
	etcdGroup.DELETE("deleteKey", etcd.DeleteKey)

}
