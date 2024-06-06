package httprequest

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/httprequest"
)

func EtcdRouter(Router *gin.RouterGroup) {

	etcdGroup := Router.Group("/etcdClient")
	etcd := new(httprequest.EtcdController)
	// 获取etcd键值
	etcdGroup.GET("getList", etcd.GetList)

}
