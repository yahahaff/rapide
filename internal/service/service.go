package service

import (
	"rapide/internal/service/ssl"
	"rapide/internal/service/sys"
	"rapide/internal/service/traefik"
)

var Entrance = ServiceGroup{}

type ServiceGroup struct {
	SysService     sys.SysGroup
	SSLService     ssl.SSLGroup
	TraefikService traefik.TraefikGroup
}
