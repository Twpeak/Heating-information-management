package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api"
	"github.com/gin-gonic/gin"
)

type HospitalRouter struct {
	
}

func (s *BaseRouter)InitHospitalRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	hospitalRouter := Router.Group("hospital")
	hospitalApi := api.ApiGroupApp.SystemApiGroup.HospitalApi
	{
		hospitalRouter.GET("all",hospitalApi.GetHospitalAndBoss)
	}
	return hospitalRouter
}
