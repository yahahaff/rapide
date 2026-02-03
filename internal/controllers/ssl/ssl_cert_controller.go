// Package ssl
package ssl

import (
	"archive/zip"
	"bytes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"rapide/internal/controllers"
	sslModel "rapide/internal/models/ssl"
	requestsSSL "rapide/internal/requests/ssl"
	"rapide/internal/requests/validators"
	"rapide/internal/service"
	"rapide/pkg/response"
)

// SSLCertController SSL证书控制器
type SSLCertController struct {
	controllers.BaseAPIController
}

// GetSSLCertList 获取SSL证书列表
func (ctrl *SSLCertController) GetSSLCertList(c *gin.Context) {
	request := requestsSSL.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 10 // 设置默认值
	}

	// 处理页码参数，确保页码大于0
	page := request.Page
	if page <= 0 {
		page = 1
	}

	// 获取查询参数
	domain := request.Domain
	applyStatus := request.ApplyStatus

	data, total, err := service.Entrance.SSLService.SSLCertService.GetSSLCertList(page, pageSize, domain, applyStatus)
	if err != nil {
		response.Abort500(c, "获取SSL证书列表失败")
		return
	}

	result := gin.H{
		"list":     data,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}
	response.OK(c, result)
}

// CreateSSLCert 创建SSL证书
func (ctrl *SSLCertController) CreateSSLCert(c *gin.Context) {
	request := requestsSSL.SSLCertCreateRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 设置默认值
	provider := request.Provider
	if provider == "" {
		provider = "letsencrypt"
	}

	challengeType := request.ChallengeType
	if challengeType == "" {
		// 根据verifyMethod设置默认的验证方式
		if request.VerifyMethod == "auto-dns" {
			challengeType = "dns-01"
		} else {
			challengeType = "http-01"
		}
	}

	// 转换请求数据为证书模型
	cert := sslModel.SSLCert{
		Domain:           request.Domain,
		CommonName:       request.CommonName,
		Organization:     request.Organization,
		OrganizationUnit: request.OrganizationUnit,
		Country:          request.Country,
		State:            request.State,
		City:             request.City,
		Email:            request.Email,
		Type:             "DV", // Let's Encrypt 只提供 DV 证书
		Algorithm:        request.Algorithm,
		Provider:         provider,
		ChallengeType:    challengeType,
		ApplyStatus:      "pending",
		AutoRenew:        request.AutoRenew,
		RenewStatus:      "idle",
		Status:           1, // 默认为启用状态
	}

	// 调用服务层创建证书
	err := service.Entrance.SSLService.SSLCertService.CreateSSLCert(cert)
	if err != nil {
		response.Abort500(c, "创建SSL证书失败")
		return
	}

	response.OK(c, gin.H{"message": "SSL证书创建成功，正在申请中"})
}

// DownloadSSLCert 下载SSL证书
// @Summary 下载SSL证书
// @Description 下载SSL证书，返回包含key和pem文件的压缩包
// @Tags SSL证书
// @Accept json
// @Produce octet-stream
// @Param id path string true "证书ID"
// @Success 200 {file} binary "证书压缩包"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "证书不存在"
// @Failure 500 {object} response.Response "下载失败"
// @Router /api/ssl/download/{id} [get]
func (ctrl *SSLCertController) DownloadSSLCert(c *gin.Context) {
	// 1. 获取证书ID
	certID := c.Param("id")

	// 2. 获取证书信息
	cert, err := service.Entrance.SSLService.SSLCertService.GetSSLCertByID(certID)
	if err != nil {
		response.Abort404(c, "证书不存在")
		return
	}

	// 3. 验证证书状态
	if cert.ApplyStatus != "success" {
		response.Abort400(c, "证书尚未申请成功，无法下载")
		return
	}

	// 4. 处理通配符域名，将*替换为wildcard-或使用主域名
	cleanDomain := cert.Domain
	// 如果是通配符域名，移除*或替换为wildcard-
	if len(cleanDomain) > 2 && cleanDomain[:2] == "*." {
		// 选项1: 替换为wildcard-前缀
		cleanDomain = "wildcard-" + cleanDomain[2:]
		// 选项2: 直接使用主域名（如果需要可以取消注释）
		// cleanDomain = cleanDomain[2:]
	}

	// 5. 创建包含key和pem文件的压缩包
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	defer w.Close()

	// 6. 添加私钥文件
	keyFileName := fmt.Sprintf("%s.key", cleanDomain)
	keyFile, err := w.Create(keyFileName)
	if err != nil {
		response.Abort500(c, "生成证书压缩包失败: "+err.Error())
		return
	}
	if cert.PrivateKey == "" {
		response.Abort500(c, "证书私钥为空，无法下载")
		return
	}
	if _, err := keyFile.Write([]byte(cert.PrivateKey)); err != nil {
		response.Abort500(c, "生成证书压缩包失败: "+err.Error())
		return
	}

	// 7. 添加证书文件（PEM格式）
	certFileName := fmt.Sprintf("%s.pem", cleanDomain)
	certFile, err := w.Create(certFileName)
	if err != nil {
		response.Abort500(c, "生成证书压缩包失败: "+err.Error())
		return
	}
	// 证书文件包含主证书和中间证书（如果有）
	certContent := cert.Certificate
	if cert.IntermediateCert != "" {
		certContent += "\n" + cert.IntermediateCert
	}
	if certContent == "" {
		response.Abort500(c, "证书内容为空，无法下载")
		return
	}
	if _, err := certFile.Write([]byte(certContent)); err != nil {
		response.Abort500(c, "生成证书压缩包失败: "+err.Error())
		return
	}

	// 8. 关闭zip写入器
	if err := w.Close(); err != nil {
		response.Abort500(c, "生成证书压缩包失败: "+err.Error())
		return
	}

	zipContent := buf.Bytes()

	// 9. 返回压缩包
	c.Writer.Header().Set("Content-Description", "File Transfer")
	// 使用处理后的域名作为压缩包名称
	zipFileName := fmt.Sprintf("%s-ssl-cert.zip", cleanDomain)
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipFileName))
	c.Writer.Header().Set("Content-Type", "application/zip")
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(zipContent)))
	c.Writer.WriteHeader(200)
	c.Writer.Write(zipContent)
	c.Writer.Flush()
}

// RevokeSSLCert 吊销SSL证书
// @Summary 吊销SSL证书
// @Description 吊销指定ID的SSL证书
// @Tags SSL证书
// @Accept json
// @Produce json
// @Param id path string true "证书ID"
// @Success 200 {object} response.Response "吊销成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "证书不存在"
// @Failure 500 {object} response.Response "吊销失败"
// @Router /api/ssl/revoke/{id} [post]
func (ctrl *SSLCertController) RevokeSSLCert(c *gin.Context) {
	// 1. 获取证书ID
	certID := c.Param("id")

	// 2. 调用服务层吊销证书
	err := service.Entrance.SSLService.SSLCertService.RevokeSSLCert(certID)
	if err != nil {
		response.Abort500(c, "吊销证书失败: "+err.Error())
		return
	}

	// 3. 返回成功响应
	response.OK(c, gin.H{"message": "证书吊销成功"})
}

// GetSSLCertDetail 获取单个SSL证书详情
// @Summary 获取单个SSL证书详情
// @Description 获取指定ID的SSL证书详情
// @Tags SSL证书
// @Accept json
// @Produce json
// @Param id path string true "证书ID"
// @Success 200 {object} response.Response "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "证书不存在"
// @Failure 500 {object} response.Response "获取失败"
// @Router /api/ssl/detail/{id} [get]
func (ctrl *SSLCertController) GetSSLCertDetail(c *gin.Context) {
	// 1. 获取证书ID
	certID := c.Param("id")

	// 2. 调用服务层获取证书详情
	cert, err := service.Entrance.SSLService.SSLCertService.GetSSLCertByID(certID)
	if err != nil {
		response.Abort404(c, "证书不存在")
		return
	}

	// 3. 计算证书剩余天数
	now := time.Now()
	expiresInDays := int(cert.ValidityEnd.Sub(now).Hours() / 24)
	// 确保剩余天数不为负数
	if expiresInDays < 0 {
		expiresInDays = 0
	}

	// 4. 创建返回结果，包含证书信息和剩余天数（移除敏感字段）
	result := gin.H{
		"id":               cert.ID,
		"domain":           cert.Domain,
		"commonName":       cert.CommonName,
		"organization":     cert.Organization,
		"organizationUnit": cert.OrganizationUnit,
		"country":          cert.Country,
		"state":            cert.State,
		"city":             cert.City,
		"email":            cert.Email,
		"type":             cert.Type,
		"algorithm":        cert.Algorithm,
		"validityStart":    cert.ValidityStart,
		"validityEnd":      cert.ValidityEnd,
		"status":           cert.Status,
		"provider":         cert.Provider,
		"challengeType":    cert.ChallengeType,
		"applyStatus":      cert.ApplyStatus,
		"errorMsg":         cert.ErrorMsg,
		"fingerprint":      cert.Fingerprint,
		"serialNumber":     cert.SerialNumber,
		"autoRenew":        cert.AutoRenew,
		"renewStatus":      cert.RenewStatus,
		"created_at":       cert.CreatedAt,
		"updated_at":       cert.UpdatedAt,
		"expiresInDays":    expiresInDays,
	}

	// 5. 返回证书详情
	response.OK(c, result)
}
