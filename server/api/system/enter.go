package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CaptchaApi
	BaseApi
	HospitalApi
}

var (
	userService 	= service.ServiceGroupApp.SystemServiceGroup.UserService
	casbinService 	= service.ServiceGroupApp.SystemServiceGroup.CasbinService
	jwtService   	= service.ServiceGroupApp.SystemServiceGroup.JwtService
	hospitalService = service.ServiceGroupApp.SystemServiceGroup.HospitalService
)
