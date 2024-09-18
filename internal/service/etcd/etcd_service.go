package etcd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/httpclient"
)

// EtcdService provides methods for interacting with etcd
type EtcdService struct{}

// Data represents the request payload for etcd operations
type Data struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	RangeEnd string `json:"range_end"`
	KeysOnly bool   `json:"keys_only"`
}

// Response represents the response from etcd
type Response struct {
	Header map[string]interface{} `json:"header"`
	Kvs    []Kvs                  `json:"kvs"`
	Count  string                 `json:"count"`
}

// ErrorResponse represents an error response from etcd
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Kvs represents the key-value pair returned from etcd
type Kvs struct {
	CreateRevision string `json:"create_revision"`
	Key            string `json:"key"`
	ModRevision    string `json:"mod_revision"`
	Value          string `json:"value"`
	Version        string `json:"version"`
}

// RequestHeaders returns the headers used in etcd requests
func RequestHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Email": "rapide@example.com",
	}
}

// GetEtcdHttpUrl returns the etcd URL from configuration
func (es *EtcdService) GetEtcdHttpUrl() string {
	return config.GetString("ETCD_HTTP_URL", "http://localhost:2379")
}

// DecodeBase64 decodes Base64 encoded key and value
func (k *Kvs) DecodeBase64() error {
	if decodedKey, err := base64.StdEncoding.DecodeString(k.Key); err != nil {
		return err
	} else {
		k.Key = string(decodedKey)
	}
	if decodedValue, err := base64.StdEncoding.DecodeString(k.Value); err != nil {
		return err
	} else {
		k.Value = string(decodedValue)
	}
	return nil
}

// request sends a request to etcd and processes the response
func (es *EtcdService) request(endpoint string, data *Data) (*Response, error) {
	url := es.GetEtcdHttpUrl() + endpoint
	headers := RequestHeaders()
	body, _ := httpclient.ConvertStructToReader(data)
	responseBody, err := httpclient.Request("POST", url, headers, body)
	if err != nil {
		return nil, err
	}

	var response Response
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// handleResponseError processes an error response from etcd
func handleResponseError(responseBody []byte) error {
	var responseError ErrorResponse
	if err := json.Unmarshal(responseBody, &responseError); err == nil && responseError.Code != 0 {
		return fmt.Errorf("code: %d, message: %s", responseError.Code, responseError.Message)
	}
	return fmt.Errorf("unexpected response format")
}

// GetKey retrieves keys from etcd
func (es *EtcdService) GetKeyList(Key, Range string) ([]Kvs, error) {
	data := Data{
		Key:      Key,
		RangeEnd: Range,
		KeysOnly: true,
	}
	response, err := es.request("/v3/kv/range", &data)
	if err != nil {
		return nil, err
	}

	if len(response.Kvs) == 0 {
		return nil, nil
	}

	for i := range response.Kvs {
		if err := response.Kvs[i].DecodeBase64(); err != nil {
			return nil, err
		}
	}

	return response.Kvs, nil
}

// GetKeyDetail retrieves detailed information for a specific key
func (es *EtcdService) GetKeyDetail(Key string) (*Kvs, error) {
	encodedKey := base64.StdEncoding.EncodeToString([]byte(Key))
	data := Data{Key: encodedKey}
	response, err := es.request("/v3/kv/range", &data)
	if err != nil {
		return nil, err
	}

	if len(response.Kvs) == 0 {
		return nil, nil
	}

	if err := response.Kvs[0].DecodeBase64(); err != nil {
		return nil, err
	}

	return &response.Kvs[0], nil
}

// PutData puts a key-value pair into etcd
func (es *EtcdService) PutData(key, value string) (*Response, error) {
	data := Data{
		Key:   base64.StdEncoding.EncodeToString([]byte(key)),
		Value: base64.StdEncoding.EncodeToString([]byte(value)),
	}
	response, err := es.request("/v3/kv/put", &data)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteData deletes a key from etcd
func (es *EtcdService) DeleteData(key string) (*Response, error) {
	data := Data{
		Key: base64.StdEncoding.EncodeToString([]byte(key)),
	}
	response, err := es.request("/v3/kv/deleterange", &data)
	if err != nil {
		return nil, err
	}

	return response, nil
}
