package service

import "server/service/system"

type ServicesGroup struct {
	SystemServiceGroup system.SysServiceGroup
}

var ServicesGroupApp = new(ServicesGroup)
