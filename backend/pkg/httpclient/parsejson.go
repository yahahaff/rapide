package httpclient

import (
	"encoding/json"
	"github.com/yahahaff/rapide/backend/pkg/logger"
)

func ParseJSONResponse(body []byte, responseData interface{}) error {
	err := json.Unmarshal(body, responseData)
	if err != nil {
		logger.ErrorString("cloudflare", "Error", err.Error())
		return err
	}
	return nil
}
