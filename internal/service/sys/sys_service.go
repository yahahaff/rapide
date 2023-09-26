package sys

type SysGroup struct {
	AuthService
	UserService
	CasbinService
	OperationLogService
	MenuService
	SignupService
	RoleService

	// 其他方法
}
