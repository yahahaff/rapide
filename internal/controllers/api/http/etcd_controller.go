package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	reqHttp "github.com/yahahaff/rapide/internal/requests/http"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/logger"
	"net/http"
)

type EtcdController struct {
	api.BaseAPIController
}

// GetKey
// @Summary  获取key
// @Schemes http.EtcdRangeRequest{}
// @Param data body http.EtcdRangeRequest{} true "body"
// @Security Bearer
// @Description
// @Tags ETCD
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v3/kv/range [post]
func (tc *EtcdController) GetKey(c *gin.Context) {
	// 1. 验证参数
	request := reqHttp.EtcdRangeRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}
	// 查询所有keys
	if request.Key == "all" {
		request.Key = "AA=="
		request.RangeEnd = "AA=="
	}

	data, err := service.Entrance.HttpRequestService.GetKey(request.Key, request.RangeEnd)
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

// PutKey
// @Schemes http.EtcdPutRequest{}
// @Param data body http.EtcdPutRequest{} true "body"
// @Security Bearer
// @Description
// @Tags ETCD
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v3/kv/put [post]
func (tc *EtcdController) PutKey(c *gin.Context) {

	// 1. 验证参数
	request := reqHttp.EtcdPutRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	data, err := service.Entrance.HttpRequestService.PutData(request.Key, request.Value)
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

// DeleteKey
// @Schemes http.EtcdRangeRequest{}
// @Param data body http.EtcdRangeRequest{} true "body"
// @Security Bearer
// @Description
// @Tags ETCD
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v3/kv/deleteRange [post]
func (tc *EtcdController) DeleteKey(c *gin.Context) {

	// 1. 验证参数
	request := reqHttp.EtcdRangeRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	data, err := service.Entrance.HttpRequestService.DeleteData(request.Key)
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
