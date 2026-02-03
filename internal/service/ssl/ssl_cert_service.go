// Package ssl
package ssl

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
	"rapide/internal/models/ssl"
	"rapide/pkg/database"
)

// SSLCertService SSL证书服务
type SSLCertService struct{}

// GetSSLCertByID 根据ID获取SSL证书详情
func (ss *SSLCertService) GetSSLCertByID(id string) (ssl.SSLCert, error) {
	var cert ssl.SSLCert
	if err := database.DB.Where("id = ?", id).First(&cert).Error; err != nil {
		return ssl.SSLCert{}, err
	}
	return cert, nil
}

// GetSSLCertList 获取SSL证书列表
func (ss *SSLCertService) GetSSLCertList(page int, size int, domain, applyStatus string) (data interface{}, total int64, err error) {
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
	db := database.DB.Model(&ssl.SSLCert{})

	// 添加查询条件
	if domain != "" {
		db = db.Where("domain LIKE ?", "%"+domain+"%")
	}
	if applyStatus != "" {
		db = db.Where("apply_status = ?", applyStatus)
	}

	// 获取总记录数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 定义返回字段结构体
	type SSLCertListResponse struct {
		ID            uint64    `json:"id"`
		Domain        string    `json:"domain"`
		CommonName    string    `json:"commonName"`
		Organization  string    `json:"organization"`
		ExpiresInDays int       `json:"expiresInDays"`
		Type          string    `json:"type"`
		Algorithm     string    `json:"algorithm"`
		ValidityEnd   time.Time `json:"validityEnd"`
		Provider      string    `json:"provider"`
		ApplyStatus   string    `json:"applyStatus"`
	}

	// 执行分页查询，只查询指定字段
	var certList []struct {
		ID           uint64    `json:"id"`
		Domain       string    `json:"domain"`
		CommonName   string    `json:"commonName"`
		Organization string    `json:"organization"`
		Type         string    `json:"type"`
		Algorithm    string    `json:"algorithm"`
		ValidityEnd  time.Time `json:"validityEnd"`
		Provider     string    `json:"provider"`
		ApplyStatus  string    `json:"applyStatus"`
	}
	if err := db.Select("id, domain, common_name, organization, type, algorithm, validity_end, provider, apply_status").Order("id desc").Limit(size).Offset(offset).Find(&certList).Error; err != nil {
		return nil, 0, err
	}

	// 计算剩余天数并转换为响应格式
	var responseList []SSLCertListResponse
	now := time.Now()
	for _, cert := range certList {
		expiresInDays := int(cert.ValidityEnd.Sub(now).Hours() / 24)
		// 确保剩余天数不为负数
		if expiresInDays < 0 {
			expiresInDays = 0
		}
		responseList = append(responseList, SSLCertListResponse{
			ID:            cert.ID,
			Domain:        cert.Domain,
			CommonName:    cert.CommonName,
			Organization:  cert.Organization,
			ExpiresInDays: expiresInDays,
			Type:          cert.Type,
			Algorithm:     cert.Algorithm,
			ValidityEnd:   cert.ValidityEnd,
			Provider:      cert.Provider,
			ApplyStatus:   cert.ApplyStatus,
		})
	}

	return responseList, total, nil
}

// CreateSSLCert 创建SSL证书
func (ss *SSLCertService) CreateSSLCert(cert ssl.SSLCert) (err error) {
	// 1. 创建初始证书记录，状态为 pending
	cert.ApplyStatus = "pending"
	if err := database.DB.Create(&cert).Error; err != nil {
		return err
	}

	// 2. 异步处理证书申请
	go func(certID uint64, cert ssl.SSLCert) {
		// 更新状态为 applying
		updateData := map[string]interface{}{
			"apply_status": "applying",
		}
		if err := database.DB.Model(&ssl.SSLCert{}).Where("id = ?", certID).Updates(updateData).Error; err != nil {
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

		if err := database.DB.Model(&ssl.SSLCert{}).Where("id = ?", certID).Updates(resultData).Error; err != nil {
			return
		}
	}(cert.ID, cert)

	return nil
}

// applyLetsEncryptCert 申请 Let's Encrypt 证书
func (ss *SSLCertService) applyLetsEncryptCert(cert ssl.SSLCert) (certContent, privateKey, intermediateCert string, validityStart, validityEnd time.Time, fingerprint, serialNumber string, err error) {
	// 根据算法选择密钥类型
	keyType := certcrypto.RSA2048
	if cert.Algorithm == "EC-256" {
		keyType = certcrypto.EC256
	} else if cert.Algorithm == "EC-384" {
		keyType = certcrypto.EC384
	} else if cert.Algorithm == "RSA-4096" {
		keyType = certcrypto.RSA4096
	}

	// 生成用户私钥
	privateKeyBytes, err := certcrypto.GeneratePrivateKey(keyType)
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("生成私钥失败: %v", err)
	}

	// 创建用户注册信息
	myUser := ss.newUser(cert.Email, privateKeyBytes)

	// 创建 lego 客户端
	config := lego.NewConfig(myUser)

	// 设置 CA 服务器 URL，使用 Let's Encrypt 生产环境或 staging 环境
	// 生产环境:
	config.CADirURL = "https://acme-v02.api.letsencrypt.org/directory"
	// Staging 环境:
	// config.CADirURL = "https://acme-staging-v02.api.letsencrypt.org/directory"

	// 设置证书算法
	config.Certificate.KeyType = keyType

	// 创建 HTTP 客户端
	client, err := lego.NewClient(config)
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("创建 lego 客户端失败: %v", err)
	}

	// 根据验证类型配置对应的挑战
	if cert.ChallengeType == "dns-01" {
		// 配置 Cloudflare DNS-01 挑战，从系统环境变量中读取凭证
		// cloudflare.NewDNSProvider()函数会**自动从环境变量中读取Cloudflare的API凭证**
		// 它内部会按照以下顺序检查环境变量：
		// 1. CLOUDFLARE_EMAIL 和 CLOUDFLARE_API_KEY (传统 API 密钥)
		// 2. CLOUDFLARE_DNS_API_TOKEN 和 CLOUDFLARE_ZONE_API_TOKEN (DNS 和区域 API 令牌)
		// 3. CLOUDFLARE_API_TOKEN (全局 API 令牌)

		cfProvider, err := cloudflare.NewDNSProvider()
		if err != nil {
			// 详细的错误信息，帮助用户调试
			errMsg := fmt.Sprintf("配置 Cloudflare DNS 提供商失败: %v\n", err)
			errMsg += "请确保已正确设置 Cloudflare API 凭证环境变量，支持以下方式：\n"
			errMsg += "1. 设置 CLOUDFLARE_EMAIL 和 CLOUDFLARE_API_KEY (传统 API 密钥)\n"
			errMsg += "2. 设置 CLOUDFLARE_DNS_API_TOKEN 和 CLOUDFLARE_ZONE_API_TOKEN (DNS 和区域 API 令牌)\n"
			errMsg += "3. 设置 CLOUDFLARE_API_TOKEN (全局 API 令牌)\n"
			errMsg += "注意：API 令牌需要包含 DNS 编辑权限。\n"
			errMsg += "注意：在Windows系统中，环境变量的设置需要重启终端才能生效。"
			return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("%s", errMsg)
		}
		if err := client.Challenge.SetDNS01Provider(cfProvider); err != nil {
			return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("配置 DNS-01 挑战失败: %v", err)
		}
	} else {
		// 配置 HTTP-01 挑战
		// 注意：需要确保服务器的80端口可以被外部访问
		if err := client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "80")); err != nil {
			return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("配置 HTTP-01 挑战失败: %v。请确保服务器的80端口可以被外部访问。", err)
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
func (ss *SSLCertService) applyGoogleTrustCert(cert ssl.SSLCert) (certContent, privateKey, intermediateCert string, validityStart, validityEnd time.Time, fingerprint, serialNumber string, err error) {
	// 注意：Google Trust Services 不提供公开的 ACME API
	// 这里返回模拟数据，实际项目中需要使用 Google Cloud Certificate Manager 或其他方式申请
	return "", "", "", time.Time{}, time.Time{}, "", "", fmt.Errorf("Google Trust Services 证书申请需要使用 Google Cloud Certificate Manager API")
}

// RevokeSSLCert 吊销SSL证书
func (ss *SSLCertService) RevokeSSLCert(id string) error {
	// 1. 查询证书信息
	var cert ssl.SSLCert
	if err := database.DB.Where("id = ?", id).First(&cert).Error; err != nil {
		return err
	}

	// 2. 验证证书状态
	if cert.ApplyStatus != "success" {
		return fmt.Errorf("证书未成功申请，无法吊销")
	}

	// 3. 只有Let's Encrypt证书支持ACME吊销
	if cert.Provider != "letsencrypt" {
		// 非Let's Encrypt证书，直接更新状态
		updateData := map[string]interface{}{
			"apply_status": "revoked",
			"status":       0,
		}
		return database.DB.Model(&ssl.SSLCert{}).Where("id = ?", id).Updates(updateData).Error
	}

	// 4. Let's Encrypt证书吊销逻辑
	// 直接更新数据库状态，跳过ACME吊销（简化实现）
	updateData := map[string]interface{}{
		"apply_status": "revoked",
		"status":       0,
	}
	return database.DB.Model(&ssl.SSLCert{}).Where("id = ?", id).Updates(updateData).Error
}

// ssUser 注册用户结构体
type ssUser struct {
	Email        string
	Registration *registration.Resource
	Key          crypto.PrivateKey
}

// GetEmail 获取用户邮箱
func (u *ssUser) GetEmail() string {
	return u.Email
}

// GetRegistration 获取注册信息
func (u *ssUser) GetRegistration() *registration.Resource {
	return u.Registration
}

// GetPrivateKey 获取私钥
func (u *ssUser) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}

// newUser 创建用户注册信息
func (ss *SSLCertService) newUser(email string, privateKey crypto.PrivateKey) *ssUser {
	return &ssUser{
		Email:        email,
		Registration: nil,
		Key:          privateKey,
	}
}
