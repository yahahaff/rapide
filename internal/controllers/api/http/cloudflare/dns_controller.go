package cloudflare

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/api"
	"github.com/yahahaff/rapide/internal/requests/http/cloudflare"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/response"
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/httpclient"
	"github.com/yahahaff/rapide/pkg/logger"
)

type DnsController struct {
	api.BaseAPIController
}

type DnsResponse struct {
	Errors     []DnsError    `json:"errors"`
	Messages   []interface{} `json:"messages"`
	Result     interface{}   `json:"result"`
	ResultInfo interface{}   `json:"result_info"`
	Success    bool          `json:"success"`
}

type DnsError struct {
	Code       int          `json:"code"`
	Message    string       `json:"message"`
	ErrorChain []ErrorChain `json:"error_chain"`
}

type ErrorChain struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (d *DnsController) GetDnsList(c *gin.Context) {
	CloudflareApi := config.GetString("cloudflare.api")
	fmt.Println(CloudflareApi)
	zoneId := c.Query("zone_id")
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", zoneId)
	headers := CloudflareHeaders()
	body, err := httpclient.Request("GET", url, headers, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 处理 JSON 数据
	var responseData DnsResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal DnsError:", err)
		response.Abort500(c)
		return
	}

	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		if firstError.ErrorChain == nil {
			response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
			return
		}
		if len(firstError.ErrorChain) > 0 {
			firstErrorChain := firstError.ErrorChain[0]
			fmt.Println("Error chain code:", firstErrorChain.Code)
			fmt.Println("Error chain message:", firstErrorChain.Message)
			response.Error(c, response.WithCode(firstErrorChain.Code), response.WithMessage(firstErrorChain.Message))
			return
		}
	}

	// 组装数据
	Data := ZonesCustomResponse{
		Zones: responseData.Result.(interface{}),
		Page:  responseData.ResultInfo.(interface{}),
	}
	response.OK(c, Data)
}

func (d *DnsController) CreateDnsRecord(c *gin.Context) {
	// 1.参数验证
	var request cloudflare.DnsCreateRequest

	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2.请求三方
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", request.ZoneId)
	headers := CloudflareHeaders()

	payload := cloudflare.DnsCreateRequest{
		Content: request.Content,
		Name:    request.Name,
		Proxied: request.Proxied,
		Type:    request.Type,
		Comment: request.Comment,
		TTL:     request.TTL,
	}

	// 将结构体转换为 JSON 字节流
	body, _ := httpclient.ConvertStructToReader(&payload)
	responseBody, err := httpclient.Request("POST", url, headers, body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 处理 JSON 数据
	var responseData DnsResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal DnsError:", err)
		response.Abort500(c)
		return
	}
	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		if firstError.ErrorChain == nil {
			response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
			return
		}
		if len(firstError.ErrorChain) > 0 {
			firstErrorChain := firstError.ErrorChain[0]
			response.Error(c, response.WithCode(firstErrorChain.Code), response.WithMessage(firstErrorChain.Message))
			return
		}
	}

	response.OK(c, responseData.Result)

}

func (d *DnsController) DeleteDnsRecord(c *gin.Context) {
	// 1.参数验证
	request := cloudflare.DnsDeleteRequest{
		ZoneId: c.Query("zone_id"),
		DnsId:  c.Query("dns_id"),
	}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2.请求三方
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", request.ZoneId, request.DnsId)
	headers := CloudflareHeaders()

	responseBody, err := httpclient.Request("DELETE", url, headers, nil)
	if err != nil {
		logger.WarnString("cloudflare", "请求cloudflare", fmt.Sprintf("%s", err.Error()))
		response.Abort500(c)
		return
	}
	logger.DebugString("cloudflare", "cloudflare删除dns响应数据", fmt.Sprintf("%s", responseBody))

	// 处理 JSON 数据
	var responseData DnsResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal DnsError:", err)
		response.Abort500(c)
		return
	}
	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		if firstError.ErrorChain == nil {
			response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
			return
		}
	}

	response.OK(c, responseData.Result)

}

func (d *DnsController) GetDnsRecordDetail(c *gin.Context) {
	// 1.参数验证
	request := cloudflare.DnsDeleteRequest{
		ZoneId: c.Query("zone_id"),
		DnsId:  c.Query("dns_id"),
	}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2.请求三方
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", request.ZoneId, request.DnsId)
	headers := CloudflareHeaders()

	responseBody, err := httpclient.Request("GET", url, headers, nil)
	if err != nil {
		logger.WarnString("cloudflare", "请求cloudflare", fmt.Sprintf("%s", err.Error()))
		response.Abort500(c)
		return
	}
	logger.DebugString("cloudflare", "cloudflare-dns响应数据", fmt.Sprintf("%s", responseBody))

	// 处理 JSON 数据
	var responseData DnsResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal DnsError:", err)
		response.Abort500(c)
		return
	}
	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		if firstError.ErrorChain == nil {
			response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
			return
		}
	}

	response.OK(c, responseData.Result)

}

func (d *DnsController) UpdateDnsRecord(c *gin.Context) {
	// 1.参数验证
	var request cloudflare.DnsUpdateRequest

	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2.请求三方
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", request.ZoneId, request.DnsId)
	headers := CloudflareHeaders()
	payload := cloudflare.DnsUpdateRequest{
		Content: request.Content,
		Name:    request.Name,
		Proxied: request.Proxied,
		Type:    request.Type,
		Comment: request.Comment,
		TTL:     request.TTL,
	}

	// 将结构体转换为 JSON 字节流
	body, _ := httpclient.ConvertStructToReader(&payload)
	responseBody, err := httpclient.Request("PUT", url, headers, body)
	if err != nil {
		logger.WarnString("cloudflare", "请求cloudflare", fmt.Sprintf("%s", err.Error()))
		response.Abort500(c)
		return
	}
	logger.DebugString("cloudflare", "cloudflare-dns响应数据", fmt.Sprintf("%s", responseBody))

	// 处理 JSON 数据
	var responseData DnsResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("JSON Unmarshal DnsError:", err)
		response.Abort500(c)
		return
	}
	// 处理失败逻辑
	if !responseData.Success {
		firstError := responseData.Errors[0]
		if firstError.ErrorChain == nil {
			response.Error(c, response.WithCode(firstError.Code), response.WithMessage(firstError.Message))
			return
		}
	}

	response.OK(c, responseData.Result)

}

func CloudflareHeaders() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Email": "yahahaff@qq.com",
		"X-Auth-Key":   "ec748b98ea360d897a9a09429a1b2abdf6e4d",
	}
	return headers
}
