// Package utils
package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"
)

// CertificateFormats 支持的证书格式
var CertificateFormats = []string{
	"pem", // Apache/Nginx 格式
}

// GenerateApacheCertPackage 生成Apache/Nginx证书包
// 包含证书、私钥和中间证书，适合直接部署到Apache/Nginx服务器
func GenerateApacheCertPackage(certPEM, keyPEM, caPEM, domain string) (map[string][]byte, error) {
	packageFiles := make(map[string][]byte)

	// 主证书
	packageFiles[fmt.Sprintf("%s.crt", domain)] = []byte(certPEM)

	// 私钥
	packageFiles[fmt.Sprintf("%s.key", domain)] = []byte(keyPEM)

	// 中间证书
	if caPEM != "" {
		packageFiles[fmt.Sprintf("%s.ca-bundle", domain)] = []byte(caPEM)
	}

	// 合并证书（有些服务器需要将证书和中间证书合并）
	if caPEM != "" {
		fullChain := certPEM + "\n" + caPEM
		packageFiles[fmt.Sprintf("%s.fullchain.crt", domain)] = []byte(fullChain)
	}

	return packageFiles, nil
}

// GenerateCertPackage 生成指定格式的证书包
func GenerateCertPackage(format, certPEM, keyPEM, caPEM, domain, password string) (map[string][]byte, error) {
	// 验证格式
	format = strings.ToLower(format)
	valid := false
	for _, f := range CertificateFormats {
		if f == format {
			valid = true
			break
		}
	}

	if !valid {
		return nil, fmt.Errorf("不支持的证书格式: %s，支持的格式有: %s", format, strings.Join(CertificateFormats, ", "))
	}

	// 根据格式生成证书包
	switch format {
	case "pem", "apache", "nginx":
		return GenerateApacheCertPackage(certPEM, keyPEM, caPEM, domain)
	default:
		return nil, fmt.Errorf("未实现的证书格式: %s", format)
	}
}

// GenerateAllFormatsCertPackage 生成包含所有证书格式的压缩包
func GenerateAllFormatsCertPackage(certPEM, keyPEM, caPEM, domain, password string) ([]byte, error) {
	// 创建一个字节缓冲区来存储zip文件
	buf := new(bytes.Buffer)
	
	// 创建一个zip写入器
	w := zip.NewWriter(buf)
	defer w.Close()

	// 为每种支持的格式生成证书包
	for _, format := range CertificateFormats {
		// 生成该格式的证书文件
		packageFiles, err := GenerateCertPackage(format, certPEM, keyPEM, caPEM, domain, password)
		if err != nil {
			return nil, fmt.Errorf("生成%s格式证书失败: %w", format, err)
		}

		// 将证书文件添加到zip包中
		for filename, content := range packageFiles {
			// 创建zip文件头
			fileHeader := &zip.FileHeader{
				Name:     fmt.Sprintf("%s/%s", format, filename),
				Method:   zip.Deflate,
			}
			
			// 创建zip文件
			file, err := w.CreateHeader(fileHeader)
			if err != nil {
				return nil, fmt.Errorf("创建zip文件%s失败: %w", filename, err)
			}
			
			// 写入文件内容
			if _, err := file.Write(content); err != nil {
				return nil, fmt.Errorf("写入zip文件%s失败: %w", filename, err)
			}
		}
	}

	// 关闭zip写入器
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("关闭zip写入器失败: %w", err)
	}

	return buf.Bytes(), nil
}
