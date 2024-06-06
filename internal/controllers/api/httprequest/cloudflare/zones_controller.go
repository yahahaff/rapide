package cloudflare

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	"github.com/yahahaff/rapide/internal/requests/cloudflare"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/pkg/httpclient"
	"github.com/yahahaff/rapide/pkg/logger"
)

type ZonesController struct {
	api.BaseAPIController
}

type CreateZoneRequest struct {
	Account AccountData `json:"account"`
	Name    string      `json:"name"`
	Type    string      `json:"type"`
}

type AccountData struct {
	ID string `json:"id"`
}

type ZonesCustomResponse struct {
	Zones interface{} `json:"zones"`
	Page  interface{} `json:"page"`
}

type ZonesResponse struct {
	Errors     []ZoneError `json:"errors"`
	Messages   []string    `json:"messages"`
	ResultInfo interface{} `json:"result_info"`
	Result     interface{} `json:"result"`
	Success    bool        `json:"success"`
}

type ZoneError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (d *ZonesController) GetZoneList(c *gin.Context) {

	url := "https://api.cloudflare.com/client/v4/zones"
	headers := CloudflareHeaders()
	body, err := httpclient.Request("GET", url, headers, nil)
	if err != nil {
		logger.ErrorString("cloudflare", "Error", fmt.Sprintf(err.Error()))
		response.Abort500(c)
		return
	}

	// 处理响应JSON数据
	var responseData ZonesResponse
	err = httpclient.ParseJSONResponse(body, &responseData)
	if err != nil {
		response.Abort500(c)
		return
	}

	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
		return
	}

	// 组装数据
	Data := ZonesCustomResponse{
		Zones: responseData.Result.(interface{}),
		Page:  responseData.ResultInfo.(interface{}),
	}
	response.OK(c, Data)
}

func (d *ZonesController) CreateZone(c *gin.Context) {
	// 1.参数验证
	var request cloudflare.CreateZoneRequest
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 请求第三方
	url := "https://api.cloudflare.com/client/v4/zones"
	headers := CloudflareHeaders()

	payload := CreateZoneRequest{
		Account: AccountData{
			ID: request.AccountID,
		},
		Name: request.Name,
		Type: request.Type,
	}
	body, _ := httpclient.ConvertStructToReader(&payload)
	responseBody, err := httpclient.Request("POST", url, headers, body)
	if err != nil {
		logger.ErrorString("cloudflare", "Error", fmt.Sprintf(err.Error()))
		response.Abort500(c)
		return
	}
	// 处理 JSON 数据
	var responseData ZonesResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		response.Abort500(c)
		return
	}

	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
		return

	}

	response.OK(c, responseData)

}

func (d *ZonesController) DeleteZone(c *gin.Context) {
	// 1.参数验证
	var request cloudflare.ZoneIDRequest
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 请求三方
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s", request.ZoneID)
	headers := CloudflareHeaders()
	body, err := httpclient.Request("DELETE", url, headers, nil)
	if err != nil {
		logger.ErrorString("cloudflare", "Error", fmt.Sprintf(err.Error()))
		response.Abort500(c)
		return
	}

	// 处理 JSON 数据
	var responseData ZonesResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
		return

	}
	response.OK(c, responseData)

}

func (d *ZonesController) GetZoneDetails(c *gin.Context) {
	// 1.参数验证
	var request cloudflare.ZoneIDRequest
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 请求三方
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s", request.ZoneID)
	headers := CloudflareHeaders()
	body, err := httpclient.Request("GET", url, headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 处理 JSON 数据
	var responseData ZonesResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
		return

	}
	response.OK(c, responseData.Result)

}

func (d *ZonesController) EditZone(c *gin.Context) {
	// 1.参数验证
	var request cloudflare.EditZoneRequest
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 请求第三方
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s", request.ZoneID)
	headers := CloudflareHeaders()
	payload := cloudflare.EditZoneRequest{
		Paused: request.Paused,
	}
	body, _ := httpclient.ConvertStructToReader(&payload)
	responseBody, err := httpclient.Request("PATCH", url, headers, body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 处理 JSON 数据
	var responseData ZonesResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
		return

	}
	response.OK(c, responseData)
}

func (d *ZonesController) PurgeCacheZone(c *gin.Context) {
	// 1.参数验证
	var request cloudflare.ZoneIDRequest
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	//请求第三方
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/purge_cache", request.ZoneID)
	headers := CloudflareHeaders()
	//payload := strings.NewReader("{\n  \"tags\": [\n    \"some-tag\",\n    \"another-tag\"\n  ]\n}")
	//payload := cloudflare.DnsUpdateRequest{
	//	ZoneId: request.ZoneID,
	//}
	// 将结构体转换为 JSON 字节流
	//body, _ := httpclient.ConvertStructToReader(&payload)
	responseBody, err := httpclient.Request("POST", url, headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 处理 JSON 数据
	var responseData ZonesResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
		return

	}
	response.OK(c, responseData)

}
