// Package sys
package sys

import (
	"time"

	"github.com/yahahaff/rapide/internal/models"
)

// SSLCert SSL证书模型
type SSLCert struct {
	models.BaseModel
	Domain           string `json:"domain" gorm:"type:varchar(255);uniqueIndex;not null;comment:'域名'"`
	CommonName       string `json:"commonName" gorm:"type:varchar(255);not null;comment:'通用名称'"`
	Organization     string `json:"organization" gorm:"type:varchar(255);comment:'组织'"`
	OrganizationUnit string `json:"organizationUnit" gorm:"type:varchar(255);comment:'组织单位'"`
	Country          string `json:"country" gorm:"type:varchar(2);comment:'国家'"`
	State            string `json:"state" gorm:"type:varchar(255);comment:'州/省'"`
	City             string `json:"city" gorm:"type:varchar(255);comment:'城市'"`
	Email            string `json:"email" gorm:"type:varchar(255);comment:'邮箱'"`
	Type             string    `json:"type" gorm:"type:varchar(10);comment:'证书类型: DV/OV/EV'"`
	ValidityStart    time.Time `json:"validityStart" gorm:"type:datetime;comment:'有效期开始时间'"`
	ValidityEnd      time.Time `json:"validityEnd" gorm:"type:datetime;comment:'有效期结束时间'"`
	Status           int       `json:"status" gorm:"default:1;comment:'状态 0:禁用 1:启用'"`
	models.CommonTimestampsField
}

// TableName Set the table name
func (*SSLCert) TableName() string {
	return "sys_ssl_cert"
}
