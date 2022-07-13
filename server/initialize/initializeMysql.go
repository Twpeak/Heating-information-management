package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

var (
	casbinService 		= service.ServiceGroupApp.SystemServiceGroup.CasbinService
	userService   		= service.ServiceGroupApp.SystemServiceGroup.UserService
	hospitalService	  	= service.ServiceGroupApp.SystemServiceGroup.HospitalService
	districtService	  	= service.ServiceGroupApp.SystemServiceGroup.DistrictService
)

func InitBaseMysqlDate() {
	//Cabin鉴权数据初始化
	casbinService.InitCasbin()
	//角色和管理员数据初始化
	userService.InitUserRole()
	//地区数据初始化
	districtService.InitDistrict()
	//医院数据初始化
	hospitalService.InitHospital()
}