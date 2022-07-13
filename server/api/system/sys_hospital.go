package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HospitalApi struct {

}

//查询医院信息和负责人信息，返回vo
func (h *HospitalApi)GetHospitalAndBoss(c *gin.Context)  {
	//查询
	list,err := hospitalService.GetHospitalsVo()
	if err != nil {
		global.G_LOG.Error("查询失败",zap.Error(err))
		response.FailWithMessage("查询失败",c)
	}
	response.OkWithData(gin.H{"date":list},c)
}

//通过医院查询当前医院的所有医生


//删除医院将删除其下的所有医生账户

//修改负责人信息

//修改医院信息

//添加医院信息同时添加负责人信息
//同时通过负责人信息去自动添加用户信息