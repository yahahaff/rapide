// Package ssl
package ssl

import (
	"time"

	"rapide/internal/models"
)

// SSLCert SSL证书模型
type SSLCert struct {
	models.BaseModel
	Domain           string    `json:"domain" gorm:"type:varchar(255);uniqueIndex;not null;comment:'域名'"`
	CommonName       string    `json:"commonName" gorm:"type:varchar(255);not null;comment:'通用名称'"`
	Organization     string    `json:"organization" gorm:"type:varchar(255);comment:'组织'"`
	OrganizationUnit string    `json:"organizationUnit" gorm:"type:varchar(255);comment:'组织单位'"`
	Country          string    `json:"country" gorm:"type:varchar(2);comment:'国家'"`
	State            string    `json:"state" gorm:"type:varchar(255);comment:'州/省'"`
	City             string    `json:"city" gorm:"type:varchar(255);comment:'城市'"`
	Email            string    `json:"email" gorm:"type:varchar(255);comment:'邮箱'"`
	Type             string    `json:"type" gorm:"type:varchar(10);comment:'证书类型: DV/OV/EV'"`
	Algorithm        string    `json:"algorithm" gorm:"type:varchar(20);default:'RSA-2048';comment:'加密算法'"`
	ValidityStart    time.Time `json:"validityStart" gorm:"type:datetime;comment:'有效期开始时间'"`
	ValidityEnd      time.Time `json:"validityEnd" gorm:"type:datetime;comment:'有效期结束时间'"`
	Status           int       `json:"status" gorm:"default:1;comment:'状态 0:禁用 1:启用'"`
	// 证书提供商相关字段
	Provider      string `json:"provider" gorm:"type:varchar(50);not null;default:'letsencrypt';comment:'证书提供商: letsencrypt/google'"`
	ChallengeType string `json:"challengeType" gorm:"type:varchar(20);not null;default:'http-01';comment:'验证方式: http-01/dns-01'"`
	ApplyStatus   string `json:"applyStatus" gorm:"type:varchar(20);default:'pending';comment:'申请状态: pending/applying/success/failed'"`
	ErrorMsg      string `json:"errorMsg" gorm:"type:text;comment:'错误信息'"`
	// 证书文件存储
	Certificate      string `json:"certificate" gorm:"type:text;comment:'证书内容'"`
	PrivateKey       string `json:"privateKey" gorm:"type:text;comment:'私钥内容'"`
	IntermediateCert string `json:"intermediateCert" gorm:"type:text;comment:'中间证书内容'"`
	// 证书验证相关
	Fingerprint  string `json:"fingerprint" gorm:"type:varchar(100);comment:'证书指纹'"`
	SerialNumber string `json:"serialNumber" gorm:"type:varchar(100);comment:'证书序列号'"`
	// 自动续期相关
	AutoRenew   bool   `json:"autoRenew" gorm:"default:true;comment:'是否自动续期'"`
	RenewStatus string `json:"renewStatus" gorm:"type:varchar(20);default:'idle';comment:'续期状态: idle/renewing/success/failed'"`
	models.CommonTimestampsField
}

// TableName Set the table name
func (*SSLCert) TableName() string {
	return "sys_ssl_cert"
}
