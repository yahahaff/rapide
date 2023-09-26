package cloudflare

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api/cloudflare"
)

func DnsRouter(Router *gin.RouterGroup) {

	dnsGroup := Router.Group("/dns")
	dns := new(cloudflare.DnsController)
	// getDnsList
	dnsGroup.GET("getDnsList", dns.GetDnsList)
	dnsGroup.POST("createDnsRecord", dns.CreateDnsRecord)
	dnsGroup.DELETE("deleteDnsRecord", dns.DeleteDnsRecord)
	dnsGroup.GET("getDnsRecordDetail", dns.GetDnsRecordDetail)
	dnsGroup.PUT("UpdateDnsRecord", dns.UpdateDnsRecord)

}
