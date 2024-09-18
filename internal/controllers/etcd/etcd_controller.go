package etcd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	reqHttp "github.com/yahahaff/rapide/internal/requests/etcd"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/logger"
	"net/http"
)

type EtcdController struct {
	controllers.BaseAPIController
}

func (tc *EtcdController) GetKeyList(c *gin.Context) {
	// 查询所有keys
	data, err := service.Entrance.EtcdService.GetKeyList("AA==", "AA==")
	if err != nil {
		if err.Error() == "unexpected response format" {
			logger.ErrorString("etcd", "etcd", err.Error())
			response.Abort500(c)
		} else {
			// 解析错误消息和代码
			logger.ErrorString("etcd", "etcd", err.Error())
			response.Abort500(c)
		}
		return
	}
	response.OK(c, data)
}

func (tc *EtcdController) GetKeyDetail(c *gin.Context) {
	key := c.Query("key")
	data, err := service.Entrance.EtcdService.GetKeyDetail(key)
	if err != nil {
		if err.Error() == "unexpected response format" {
			logger.ErrorString("etcd", "etcd", err.Error())
			response.Abort500(c)
		} else {
			// 解析错误消息和代码
			logger.ErrorString("etcd", "etcd", err.Error())
			response.Abort500(c)
		}
		return
	}

	response.OK(c, data)

}

func (tc *EtcdController) PutKey(c *gin.Context) {

	// 1. 验证参数
	request := reqHttp.EtcdPutRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	data, err := service.Entrance.EtcdService.PutData(request.Key, request.Value)
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

func (tc *EtcdController) DeleteKey(c *gin.Context) {

	// 1. 验证参数
	request := reqHttp.EtcdRangeRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	data, err := service.Entrance.EtcdService.DeleteData(request.Key)
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
