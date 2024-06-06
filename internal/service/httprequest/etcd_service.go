// Package sys auth 授权相关逻辑
package httprequest

import (
	"encoding/json"
	"fmt"
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/httpclient"
)

type EtcdService struct{}

type Data struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	RangeEnd string `json:"range_end"`
}

type ResponseOk struct {
	Headers map[string]interface{} `json:"headers"`
	Kvs     []interface{}          `json:"kvs"`
	Count   string                 `json:"count"`
}

type ResponseError struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RequestHeaders() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Email": "rapide@example.com",
	}
	return headers
}

// GetEtcdHttpUrl 获取etcd地址
func (es *EtcdService) GetEtcdHttpUrl() string {
	return config.GetString("etcd_http_url")
}

// GetListData 获取所有etcd键值
func (es *EtcdService) GetListData() (interface{}, error) {
	url := es.GetEtcdHttpUrl() + "/v3/kv/range"
	headers := RequestHeaders()

	data := Data{
		Key:      "AA==",
		RangeEnd: "AA==",
	}

	body, _ := httpclient.ConvertStructToReader(&data)
	responseBody, err := httpclient.Request("POST", url, headers, body)
	if err != nil {
		return nil, err
	}

	var responseOk ResponseOk
	err = json.Unmarshal(responseBody, &responseOk)
	if err == nil && len(responseOk.Kvs) > 0 {
		return responseOk, nil
	}

	var responseError ResponseError
	err = json.Unmarshal(responseBody, &responseError)
	if err == nil && responseError.Code != 0 {
		return nil, fmt.Errorf("code: %d, message: %s", responseError.Code, responseError.Message)
	}

	return nil, fmt.Errorf("unexpected response format")
}

// PutData put etcd键值
func (es *EtcdService) PutData(k, v string) (interface{}, error) {

	url := es.GetEtcdHttpUrl() + "/v3/kv/put"
	headers := RequestHeaders()

	data := Data{
		Key:   k,
		Value: v,
	}

	body, _ := httpclient.ConvertStructToReader(&data)
	responseBody, err := httpclient.Request("POST", url, headers, body)
	if err != nil {
		return nil, err
	}

	var responseOk ResponseOk
	err = json.Unmarshal(responseBody, &responseOk)
	if err == nil && len(responseOk.Kvs) > 0 {
		return responseOk, nil
	}

	var responseError ResponseError
	err = json.Unmarshal(responseBody, &responseError)
	if err == nil && responseError.Code != 0 {
		return nil, fmt.Errorf("code: %d, message: %s", responseError.Code, responseError.Message)
	}

	return nil, fmt.Errorf("unexpected response format")
}
