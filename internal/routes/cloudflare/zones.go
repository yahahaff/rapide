package cloudflare

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/cloudflare"
)

func ZonesRouter(Router *gin.RouterGroup) {

	zonesGroup := Router.Group("/zones")
	zone := new(cloudflare.ZonesController)
	// 获取zonesList
	zonesGroup.GET("getZoneList", zone.GetZoneList)
	zonesGroup.POST("createZone", zone.CreateZone)
	zonesGroup.DELETE("deleteZone", zone.DeleteZone)
	zonesGroup.GET("getZoneDetails", zone.GetZoneDetails)
	zonesGroup.PATCH("editZone", zone.EditZone)
	zonesGroup.POST("purgeCache", zone.PurgeCacheZone)

}
