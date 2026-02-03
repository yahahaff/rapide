package sys

import (
	"rapide/internal/models"
)

type OperationLog struct {
	models.BaseModel
	ClientIP string `json:"client_ip" gorm:"not null; comment:'访问IP'"`
	Method   string `json:"method" gorm:"not null; comment:'请求方法'"`
	Path     string `json:"path" gorm:"not null; comment:'URL'"`
	Status   int    `json:"status" gorm:"not null; comment:'响应状态码'"`
	Latency  int64  `json:"latency" gorm:"not null; comment:'耗时 单位： 毫秒'"`
	Requests string `json:"requests" gorm:"type:text; comment:'请求体'"`
	Response string `json:"response" gorm:"type:text; comment:'响应体'"`
	Operator string `json:"operator" gorm:"comment:'操作人'"`
	models.CommonTimestampsField
}

// TableName Set the table name
func (OperationLog) TableName() string {
	return "sys_operation_log"
}
