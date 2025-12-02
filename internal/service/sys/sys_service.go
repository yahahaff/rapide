package sys

type SysGroup struct {
	AuthService
	SignupService
	UserService
	MenuService
	CasbinService
	OperationLogService
	SSLCertService
	// 其他方法
}
