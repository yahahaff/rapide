package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yahahaff/rapide/pkg/logger"
	"io"
	"net/http"
)

func Request(method, url string, headers map[string]string, body io.Reader) ([]byte, error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			// 处理关闭错误，例如打印日志
		}
	}()
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func ConvertStructToReader(data interface{}) (io.Reader, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.WarnString("httpclient", "", fmt.Sprintf("%s", err.Error()))
		return nil, err
	}
	return bytes.NewReader(jsonData), nil
}
