package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONSlice 自定义JSON切片类型，用于处理数据库中的JSON字符串转换为[]string

type JSONSlice []string

// Value 实现 driver.Valuer 接口，将 JSONSlice 转换为数据库值
func (j JSONSlice) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan 实现 sql.Scanner 接口，将数据库值转换为 JSONSlice
func (j *JSONSlice) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		str, ok := value.(string)
		if !ok {
			return errors.New("invalid type for JSONSlice")
		}
		// 处理空字符串情况
		if str == "" {
			*j = nil
			return nil
		}
		bytes = []byte(str)
	}

	// 处理空字节切片情况
	if len(bytes) == 0 {
		*j = nil
		return nil
	}

	return json.Unmarshal(bytes, j)
}
