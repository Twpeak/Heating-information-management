package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

var (
	casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
)

func InitBaseMysqlDate() {
	//Cabin鉴权数据初始化
	casbinService.InitCasbin()
	//角色和管理员数据初始化
	userService.InitUserRole()

}