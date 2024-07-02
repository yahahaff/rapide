// Package sys auth 授权相关逻辑
package httprequest

import (
	"encoding/base64"
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
	Header map[string]interface{} `json:"header"`
	Kvs    []interface{}          `json:"kvs"`
	Count  string                 `json:"count"`
}

type ResponseError struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Entry 处理后的数据格式
type Entry struct {
	CreateRevision string `json:"create_revision"`
	Key            string `json:"key"`
	ModRevision    string `json:"mod_revision"`
	Value          string `json:"value"`
	Version        string `json:"version"`
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
	return config.GetString("ETCD_HTTP_URL", "http://localhost:2379")
}

// DecodeBase64 方法用于对 Entry 结构体中的 Key 和 Value 进行 Base64 解码
func (e *Entry) DecodeBase64() error {
	decodedKey, err := base64.StdEncoding.DecodeString(e.Key)
	if err != nil {
		return err
	}
	e.Key = string(decodedKey)

	decodedValue, err := base64.StdEncoding.DecodeString(e.Value)
	if err != nil {
		return err
	}
	e.Value = string(decodedValue)

	return nil
}

// GetKey 查询key
func (es *EtcdService) GetKey(Key, Range string) (interface{}, error) {
	url := es.GetEtcdHttpUrl() + "/v3/kv/range"
	headers := RequestHeaders()

	// 判断参数是获取全部还是单个查询
	var data Data
	switch {
	case Key == "AA==" && Range == "AA==":
		data = Data{
			Key:      Key,
			RangeEnd: Range,
		}
	default:
		encodedKey := base64.StdEncoding.EncodeToString([]byte(Key))
		encodedRange := base64.StdEncoding.EncodeToString([]byte(Range))
		data = Data{
			Key:      encodedKey,
			RangeEnd: encodedRange,
		}
	}

	body, _ := httpclient.ConvertStructToReader(&data)
	responseBody, err := httpclient.Request("POST", url, headers, body)
	if err != nil {
		return nil, err
	}

	var responseOk ResponseOk
	err = json.Unmarshal(responseBody, &responseOk)
	if err == nil {
		// 解析 JSON 数据到结构体中
		responseOkKvs := responseOk.Kvs
		// 将 responseOk.Kvs 转换为 JSON 字符串
		jsonData, err := json.Marshal(responseOkKvs)
		if err != nil {
			panic(err)
		}

		// 将 JSON 字符串解析到 Entry 结构体中 返回数据处理
		var entries []Entry
		err = json.Unmarshal(jsonData, &entries)
		if err != nil {
			panic(err)
		}
		// 对每个条目进行 Base64 解码
		for i := range entries {
			err := entries[i].DecodeBase64()
			if err != nil {
				panic(err)
			}
		}
		return entries, nil
	}

	var responseError ResponseError
	err = json.Unmarshal(responseBody, &responseError)
	if err == nil && responseError.Code != 0 {
		return nil, fmt.Errorf("code: %d, message: %s", responseError.Code, responseError.Message)
	}

	return nil, fmt.Errorf("unexpected response format")
}

// PutData put etcd键值
func (es *EtcdService) PutData(key, value string) (interface{}, error) {

	url := es.GetEtcdHttpUrl() + "/v3/kv/put"
	headers := RequestHeaders()

	encodedKey := base64.StdEncoding.EncodeToString([]byte(key))
	encodedValue := base64.StdEncoding.EncodeToString([]byte(value))
	data := Data{
		Key:   encodedKey,
		Value: encodedValue,
	}

	body, _ := httpclient.ConvertStructToReader(&data)
	responseBody, err := httpclient.Request("POST", url, headers, body)
	if err != nil {
		return nil, err
	}

	var responseOk ResponseOk
	err = json.Unmarshal(responseBody, &responseOk)
	if err == nil {
		return responseOk, nil
	}

	var responseError ResponseError
	err = json.Unmarshal(responseBody, &responseError)
	if err == nil && responseError.Code != 0 {
		return nil, fmt.Errorf("code: %d, message: %s", responseError.Code, responseError.Message)
	}
	return nil, fmt.Errorf("unexpected response format")
}

// DeleteData delete etcd键值
func (es *EtcdService) DeleteData(key string) (interface{}, error) {

	url := es.GetEtcdHttpUrl() + "/v3/kv/deleterange"
	headers := RequestHeaders()

	encodedKey := base64.StdEncoding.EncodeToString([]byte(key))
	data := Data{
		Key: encodedKey,
	}

	body, _ := httpclient.ConvertStructToReader(&data)
	responseBody, err := httpclient.Request("POST", url, headers, body)
	if err != nil {
		return nil, err
	}

	var responseOk ResponseOk
	err = json.Unmarshal(responseBody, &responseOk)
	if err == nil {
		return responseOk, nil
	}

	var responseError ResponseError
	err = json.Unmarshal(responseBody, &responseError)
	if err == nil && responseError.Code != 0 {
		return nil, fmt.Errorf("code: %d, message: %s", responseError.Code, responseError.Message)
	}
	return nil, fmt.Errorf("unexpected response format")
}
