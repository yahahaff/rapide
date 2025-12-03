// Package sys
package sys

import (
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/http01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/go-acme/lego/v4/registration"
	"github.com/yahahaff/rapide/internal/models/sys"
	"github.com/yahahaff/rapide/pkg/database"
)

// SSLCertService SSL证书服务
type SSLCertService struct{}

// GetSSLCertList 获取SSL证书列表
func (ss *SSLCertService) GetSSLCertList(page int, size int) (data interface{}, total int64, err error) {
	// 参数验证和默认值处理
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20 // 默认每页20条，最大100条
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 构建查询
	db := database.DB.Model(&sys.SSLCert{})

	// 获取总记录数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 定义返回字段结构体
	type SSLCertListResponse struct {
		Domain       string    `json:"domain"`
		CommonName   string    `json:"commonName"`
		Organization string    `json:"organization"`
		Email        string    `json:"email"`
		Type         string    `json:"type"`
		ValidityEnd  time.Time `json:"validityEnd"`
		Provider     string    `json:"provider"`
		ApplyStatus  string    `json:"applyStatus"`
	}

	// 执行分页查询，只查询指定字段
	var certList []SSLCertListResponse
	if err := db.Select("domain, common_name, organization, email, type, validity_end, provider, apply_status").Order("id desc").Limit(size).Offset(offset).Find(&certList).Error; err != nil {
		return nil, 0, err
	}

	return certList, total, nil
}

// CreateSSLCert 创建SSL证书
func (ss *SSLCertService) CreateSSLCert(cert sys.SSLCert) (err error) {
	// 1. 创建初始证书记录，状态为 pending
	cert.ApplyStatus = "pending"
	if err := database.DB.Create(&cert).Error; err != nil {
		return err
	}

	// 2. 异步处理证书申请
	go func(certID uint64, cert sys.SSLCert) {
		// 更新状态为 applying
		updateData := map[string]interface{}{
			"apply_status": "applying",
		}
		if err := database.DB.Model(&sys.SSLCert{}).Where("id = ?", certID).Updates(updateData).Error; err != nil {
			return
		}

		// 3. 根据提供商选择不同的申请逻辑
		var certContent, privateKey, intermediateCert string
		var validityStart, validityEnd time.Time
		var fingerprint, serialNumber string

		if cert.Provider == "letsencrypt" {
			// Let's Encrypt 证书申请逻辑
			certContent, privateKey, intermediateCert, validityStart, validityEnd, fingerprint, serialNumber, err = ss.applyLetsEncryptCert(cert)
		} else if cert.Provider == "google" {
			// Google Trust Services 证书申请逻辑
			// 注意：Google Trust Services 不提供公开的 ACME API，需要使用其他方式申请
			certContent, privateKey, intermediateCert, validityStart, validityEnd, fingerprint, serialNumber, err = ss.applyGoogleTrustCert(cert)
		}

		// 4. 更新证书状态和信息
		resultData := map[string]interface{}{
			"certificate":       certContent,
			"private_key":       privateKey,
			"intermediate_cert": intermediateCert,
			"validity_start":    validityStart,
			"validity_end":      validityEnd,
			"fingerprint":       fingerprint,
			"serial_number":     serialNumber,
		}

		if err != nil {
			// 申请失败
			resultData["apply_status"] = "failed"
			resultData["error_msg"] = err.Error()
		} else {
			// 申请成功
			resultData["apply_status"] = "success"
		}

		if err := database.DB.Model(&sys.SSLCert{}).Where("id = ?", certID).Updates(resultData).Error; err != nil {
			return
		}
	}(cert.ID, cert)

	return nil
}

// applyLetsEncryptCert 申请 Let's Encrypt 证书
func (ss *SSLCertService) applyLetsEncryptCert(cert sys.SSLCert) (certContent, privateKey, intermediateCert string, validityStart, validityEnd time.Time, fingerprint, serialNumber string, err error) {
	// 生成用户私钥
	privateKeyBytes, err := certcrypto.GeneratePrivateKey(certcrypto.RSA2048)
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("生成私钥失败: %v", err)
	}

	// 创建用户注册信息
	myUser := ss.newUser(cert.Email, privateKeyBytes)

	// 创建 lego 客户端
	config := lego.NewConfig(myUser)

	// 设置 CA 服务器 URL，使用 Let's Encrypt 生产环境或 staging 环境
	// 生产环境: "https://acme-v02.api.letsencrypt.org/directory"
	// Staging 环境: "https://acme-staging-v02.api.letsencrypt.org/directory"
	config.CADirURL = "https://acme-staging-v02.api.letsencrypt.org/directory"

	// 设置证书类型
	config.Certificate.KeyType = certcrypto.RSA2048

	// 创建 HTTP 客户端
	client, err := lego.NewClient(config)
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("创建 lego 客户端失败: %v", err)
	}

	// 根据验证类型配置对应的挑战
	if cert.ChallengeType == "dns-01" {
		// 配置 Cloudflare DNS-01 挑战
		// 注意：需要在环境变量或配置文件中设置 CLOUDFLARE_API_TOKEN
		// 示例：CLOUDFLARE_API_TOKEN=your-cloudflare-api-token
		cfProvider, err := cloudflare.NewDNSProvider()
		if err != nil {
			return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("配置 Cloudflare DNS 提供商失败: %v", err)
		}
		if err := client.Challenge.SetDNS01Provider(cfProvider); err != nil {
			return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("配置 DNS-01 挑战失败: %v", err)
		}
	} else {
		// 配置 HTTP-01 挑战
		if err := client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "80")); err != nil {
			return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("配置 HTTP-01 挑战失败: %v", err)
		}
	}

	// 注册用户
	reg, err := client.Registration.Register(registration.RegisterOptions{
		TermsOfServiceAgreed: true,
	})
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("注册用户失败: %v", err)
	}

	// 更新用户注册信息
	myUser.Registration = reg

	// 请求证书
	request := certificate.ObtainRequest{
		Domains: []string{cert.Domain},
		Bundle:  true,
	}

	certRes, err := client.Certificate.Obtain(request)
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("请求证书失败: %v", err)
	}

	// 提取证书内容
	certContent = string(certRes.Certificate)
	privateKey = string(certRes.PrivateKey)

	// 提取中间证书
	if len(certRes.IssuerCertificate) > 0 {
		intermediateCert = string(certRes.IssuerCertificate)
	}

	// 解析证书获取有效期和指纹
	certBlock, _ := pem.Decode([]byte(certContent))
	if certBlock == nil {
		return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("解析证书失败")
	}

	parsedCert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("解析证书失败: %v", err)
	}

	// 设置有效期
	validityStart = parsedCert.NotBefore
	validityEnd = parsedCert.NotAfter

	// 设置序列号
	serialNumber = parsedCert.SerialNumber.Text(16)

	// 计算指纹
	fingerprintBytes := sha256.Sum256(parsedCert.Raw)
	fingerprint = fmt.Sprintf("%x", fingerprintBytes)

	return certContent, privateKey, intermediateCert, validityStart, validityEnd, fingerprint, serialNumber, nil
}

// applyGoogleTrustCert 申请 Google Trust Services 证书
func (ss *SSLCertService) applyGoogleTrustCert(cert sys.SSLCert) (certContent, privateKey, intermediateCert string, validityStart, validityEnd time.Time, fingerprint, serialNumber string, err error) {
	// 注意：Google Trust Services 不提供公开的 ACME API
	// 这里返回模拟数据，实际项目中需要使用 Google Cloud Certificate Manager 或其他方式申请
	return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("Google Trust Services 证书申请需要使用 Google Cloud Certificate Manager API")
}

// User 注册用户结构体
type User struct {
	Email        string
	Registration *registration.Resource
	Key          crypto.PrivateKey
}

// GetEmail 获取用户邮箱
func (u *User) GetEmail() string {
	return u.Email
}

// GetRegistration 获取注册信息
func (u User) GetRegistration() *registration.Resource {
	return u.Registration
}

// GetPrivateKey 获取私钥
func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}

// newUser 创建新用户
func (ss *SSLCertService) newUser(email string, key crypto.PrivateKey) *User {
	return &User{
		Email: email,
		Key:   key,
	}
}
