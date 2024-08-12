package service

import (
	"github.com/yahahaff/rapide/internal/service/http"
	"github.com/yahahaff/rapide/internal/service/sys"
)

var Entrance = ServiceGroup{}

type ServiceGroup struct {
	SysService  sys.SysGroup
	HttpService http.HttpGroup
}
