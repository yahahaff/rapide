package service

import (
	"github.com/yahahaff/rapide/backend/internal/service/http"
	"github.com/yahahaff/rapide/backend/internal/service/sys"
)

var Entrance = ServiceGroup{}

type ServiceGroup struct {
	SysService  sys.SysGroup
	HttpService http.HttpGroup
}
