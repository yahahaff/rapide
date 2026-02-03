package ssl

// SSLCertCreateRequest SSL证书创建请求
type SSLCertCreateRequest struct {
	Domain           string `json:"domain" binding:"required,max=255"`
	CommonName       string `json:"commonName" binding:"omitempty,max=255"`
	Organization     string `json:"organization" binding:"omitempty,max=255"`
	OrganizationUnit string `json:"organizationUnit" binding:"omitempty,max=255"`
	Country          string `json:"country" binding:"omitempty,len=2"`
	State            string `json:"state" binding:"omitempty,max=255"`
	City             string `json:"city" binding:"omitempty,max=255"`
	Email            string `json:"email" binding:"required,email,max=255"`
	Provider         string `json:"provider" binding:"omitempty,oneof=letsencrypt google"`
	ChallengeType    string `json:"challengeType" binding:"omitempty,oneof=http-01 dns-01"`
	AutoRenew        bool   `json:"autoRenew" binding:"omitempty"`
	Algorithm        string `json:"algorithm" binding:"omitempty"`
	VerifyMethod     string `json:"verifyMethod" binding:"omitempty"`
	Type             string `json:"type" binding:"omitempty"`
}
