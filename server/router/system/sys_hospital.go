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
		//修改医院信息时，通过医院id获取医生列表，并获取所有区域列表，去修改。
		//新增医院时，先新增医院基本信息，再添加负责人（注册用户），通过新增用户Id去修改负责人信息
		hospitalRouter.GET("all",hospitalApi.GetAllHospital)				//查询所有医院信息
		hospitalRouter.GET("Boss",hospitalApi.GetBossByBossId)			//查询医院负责人信息
		hospitalRouter.GET("hosBoss",hospitalApi.GetHospitalAndBoss)		//查询医院信息和负责人信息，返回vo
		hospitalRouter.GET("allbossId",hospitalApi.GetAllBossId)			//获取所有医院负责人Id列表
		hospitalRouter.GET("getdoc",hospitalApi.GetUserByHospitalId)		//通过医院查询当前医院的所有医生
		hospitalRouter.GET("hosbydis",hospitalApi.GetHospitalByDistrictLimit)//当前区县内分页获取医院列表
		hospitalRouter.GET("keys",hospitalApi.GetHospital)				//通过关键字查询并分页获取
		hospitalRouter.GET("name",hospitalApi.GetHospitalByHospitalName)	//通过关键字查询并分页获取
		hospitalRouter.DELETE("del",hospitalApi.DelHospital) 			//删除医院将删除其下的所有医生账户
		hospitalRouter.POST("updateboss",hospitalApi.UpdateBossByHospital)//修改负责人信息
		hospitalRouter.POST("addhos",hospitalApi.AddHospital)			//添加医院基本信息
		hospitalRouter.POST("addHosAndBoss",hospitalApi.AddHospitalAndBoss)	//添加医院信息同时添加管理员信息
		hospitalRouter.POST("updatehos",hospitalApi.UpdateHospital)		//修改医院信息（更换负责人）
		hospitalRouter.POST("updatehosAndBoss",hospitalApi.UpdateHospitalAndBoss)//修改医院信息(同时修改负责人信息)
	}
	return hospitalRouter
}
