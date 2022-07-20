package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	reqCom "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HospitalApi struct {
}

//查询医院信息和负责人信息，返回vo[分页]
func (h *HospitalApi) GetHospitalAndBoss(c *gin.Context) {
	//取参
	var req reqCom.PageInfo
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//校验
	if err := utils.Verify(req, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//查询
	list, total,err := hospitalService.GetHospitalsVo(req)
	if err != nil {
		global.G_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:		list,
		Total: 		total,
		Page: 		req.Page,
		PageSize: 	req.PageSize,
	},"获取成功", c)
}

//获取所有医院信息
func (h *HospitalApi) GetAllHospital(c *gin.Context) {
	hospital, err := hospitalService.GetAllHospital()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(hospital,"删除医院信息成功", c)
}


//查询医院负责人信息
func (h *HospitalApi) GetBossByBossId(c *gin.Context) {
	var req request.HospitalReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := hospitalService.GetBossByBossId(req.HospitalId)
	if err != nil {
		global.G_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(user, c)
}

//获取所有医院负责人Id列表
func (h *HospitalApi) GetAllBossId(c *gin.Context) {
	var ids []uint
	if err := global.G_DB.Model(&system.Hospital{}).Select("boos_id").Scan(&ids).Error; err != nil {
		global.G_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(ids, c)
}

//通过医院查询当前医院的所有医生
func (h *HospitalApi) GetUserByHospitalId(c *gin.Context) {
	//取参
	var req request.HospitalReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//校验
	if err := utils.Verify(req, utils.HospitalReqVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//查询
	userlist, err := hospitalService.GetUserByHospitalId(req)
	if err != nil {
		global.G_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(userlist, c)
}

//删除医院将删除其下的所有医生账户
func (h *HospitalApi) DelHospital(c *gin.Context) {
	//取参
	var req request.HospitalReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//校验
	if err := utils.Verify(req, utils.HospitalReqVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//删除
	if err := hospitalService.DelHospital(req); err != nil {
		global.G_LOG.Error("删除医院信息失败", zap.Error(err))
		response.FailWithMessage("删除医院信息失败", c)
		return
	}
	response.OkWithMessage("删除医院信息成功", c)
}

//修改负责人信息
func (h *HospitalApi) UpdateBossByHospital(c *gin.Context) {
	//取参
	var req request.HospitalReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//校验
	if err := utils.Verify(req, utils.HospitalReqVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hospitalService.UpdateHospitalByUser(req); err != nil {
		global.G_LOG.Error("修改医院负责人信息失败", zap.Error(err))
		response.FailWithMessage("修改医院负责人信息失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

//修改医院信息（包括更换负责人）
func (h *HospitalApi) UpdateHospital(c *gin.Context) {
	var hospitalDate system.Hospital
	if err := c.ShouldBindJSON(&hospitalDate); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//校验
	if err := utils.Verify(hospitalDate, utils.HospitalVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//修改
	if err := hospitalService.UpdateHospital(hospitalDate); err != nil {
		global.G_LOG.Error("修改医院信息失败", zap.Error(err))
		response.FailWithMessage("修改医院信息失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)

}

//修改医院信息（同时修改负责人信息）
func (h *HospitalApi) UpdateHospitalAndBoss(c *gin.Context) {
	//取参
	var hospitalAndBoss request.HospitalAndBoss
	if err := c.ShouldBindJSON(&hospitalAndBoss); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//添加
	if err := hospitalService.UpdateHospitalAndBoss(hospitalAndBoss); err != nil {
		global.G_LOG.Error("添加医院信息失败", zap.Error(err))
		response.FailWithMessage("添加医院信息失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)


}

//添加医院信息
func (h *HospitalApi) AddHospital(c *gin.Context) {
	var hospitalDate system.Hospital
	if err := c.ShouldBindJSON(&hospitalDate); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//校验
	if err := utils.Verify(hospitalDate, utils.HospitalVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//添加
	if err := hospitalService.AddHospital(hospitalDate); err != nil {
		global.G_LOG.Error("添加医院信息失败", zap.Error(err))
		response.FailWithMessage("添加医院信息失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

//添加医院信息同时添加负责人信息
func (h *HospitalApi) AddHospitalAndBoss(c *gin.Context) {
	var hospitalAndBoss request.HospitalAndBoss
	if err := c.ShouldBindJSON(&hospitalAndBoss); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//添加
	if err := hospitalService.AddHospitalAndBoss(hospitalAndBoss); err != nil {
		global.G_LOG.Error("添加医院信息失败", zap.Error(err))
		response.FailWithMessage("添加医院信息失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

//当前区县内获取医院列表[分页]
func (h *HospitalApi) GetHospitalByDistrictLimit(c *gin.Context) {
	//取参
	var req request.HospitalReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//校验
	if &req.PageInfo != nil{
		if err := utils.Verify(req.PageInfo, utils.PageInfoVerify); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	//查询
	hospitallist,total, err := hospitalService.GetHospitalByDistrictLimit(req)
	if err != nil {
		global.G_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:		hospitallist,
		Total: 		total,
		Page: 		req.Page,
		PageSize: 	req.PageSize,
	},"获取成功", c)
}

//通过医院名查询医院数据[分页]
func (h *HospitalApi)GetHospitalByHospitalName(c *gin.Context)  {
	//取参
	var req request.KeyReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//校验
	if &req.PageInfo != nil{
		if err := utils.Verify(req.PageInfo, utils.PageInfoVerify); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	//查询
	hospitallist,total, err := hospitalService.GetHospitalByHospitalName(req)
	if err != nil {
		global.G_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:		hospitallist,
		Total: 		total,
		Page: 		req.Page,
		PageSize: 	req.PageSize,
	},"获取成功", c)
}

//通过关键字查询[分页]获取
//两种思路：1.利用mysql视图 2.使用redis缓存Zset排序
func (h *HospitalApi) GetHospital(c *gin.Context) {
	//取参
	var req request.KeyReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//查询
	hospitallist,total, err := hospitalService.GetHospitalByKey(req)
	if err != nil {
		global.G_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:		hospitallist,
		Total: 		total,
		Page: 		req.Page,
		PageSize: 	req.PageSize,
	},"获取成功", c)
}

//同时通过负责人信息去自动添加用户信息
