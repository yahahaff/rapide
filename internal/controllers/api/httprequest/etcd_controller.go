package httprequest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/logger"
	"net/http"
)

type EtcdController struct {
	api.BaseAPIController
}

// GetList
// @Summary 获取etcd 所有键值
// @Security Bearer
// @Description
// @Tags 网关系统
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/etcdClient/getList [get]
func (tc *EtcdController) GetList(c *gin.Context) {
	// 用户请求参数中不包含 captchaID 和 captchaAnswer
	data, err := service.Entrance.HttpRequestService.GetListData()
	if err != nil {
		if err.Error() == "unexpected response format" {
			logger.ErrorString("http-request", "etcd", err.Error())
			response.Abort500(c)
		} else {
			// 解析错误消息和代码
			fmt.Println("Error:", err)
			response.Error(c, response.WithCode(http.StatusBadRequest), response.WithMessage(err.Error()))
		}
		return
	}

	response.OK(c, data)

}
