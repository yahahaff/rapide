package service

import (
	"github.com/yahahaff/rapide/internal/service/sys"
)

var Entrance = ServiceGroup{}

type ServiceGroup struct {
	SysService sys.SysGroup
	//AuthService         sys.AuthService
	//UserService         sys.UserService
	//CasbinService       sys.CasbinService
	//OperationLogService sys.OperationLogService
	//MenuService         sys.MenuService
	//SignupService       sys.SignupService
}
