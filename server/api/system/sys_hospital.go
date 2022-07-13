package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
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
func (h *HospitalApi)GetUserByHospitalId(c *gin.Context)  {
	//取参
	HospitalId := c.Query("HospitalId")
	//校验
	if err := utils.Verify(HospitalId, utils.HospitalIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//查询
	userlist,err := hospitalService.GetUserByHospitalId(HospitalId)
	if  err != nil{
		global.G_LOG.Error("查询失败",zap.Error(err))
		response.FailWithMessage("查询失败",c)
	}
	response.OkWithData(gin.H{"date":userlist},c)
}


//删除医院将删除其下的所有医生账户
func (h *HospitalApi)DelHospital(c *gin.Context)  {
	//取参
	HospitalId := c.Query("HospitalId")
	//校验
	if err := utils.Verify(HospitalId, utils.HospitalIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//删除
	if  err := hospitalService.DelHospital(HospitalId);err !=nil{
		global.G_LOG.Error("删除医院信息失败",zap.Error(err))
		response.FailWithMessage("删除医院信息失败",c)
	}
	response.OkWithMessage("删除医院信息成功",c)
}


//修改负责人信息
func (h *HospitalApi)UpdateBossByHospital(c *gin.Context)  {
	//取参

}

//修改医院信息
func (h *HospitalApi)GetUserByHospitalId(c *gin.Context)  {


}

//添加医院信息同时添加负责人信息
func (h *HospitalApi)GetUserByHospitalId(c *gin.Context)  {


}

//同时通过负责人信息去自动添加用户信息
func (h *HospitalApi)GetUserByHospitalId(c *gin.Context)  {


}