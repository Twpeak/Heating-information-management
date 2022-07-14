package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type FeverApi struct {
}

//test
func (f *FeverApi)Test(c *gin.Context)  {
	response.Ok(c)
}




//删除  /del
func (f *FeverApi)DelFeverInfo(c *gin.Context) {
	//取参
	id := c.Query("id")
	if id == "" {
		response.FailWithMessage("Id为空", c)
		return
	}
	//删除
	if err := feverService.DelFeverInfo(id); err != nil {
		response.FailWithMessage("删除发热信息失败", c)
		return
	}
	response.OkWithMessage("删除发热息成功", c)
}

//编辑 /update
func (f *FeverApi)UpdateFeverInfo(c *gin.Context) {
	//取参
	var feverInfo system.FeverInfo
	if err := c.ShouldBindJSON(&feverInfo);err != nil{
		response.FailWithMessage(err.Error(),c)
		return
	}
	//校验
	if err := utils.Verify(feverInfo, utils.HospitalReqVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//获取当前用户id
	feverInfo.ID = utils.GetUserID(c)
	//编辑
	if err := feverService.UpdateFeverInfo(feverInfo); err != nil {
		response.FailWithDetailed(err.Error(),"修改发热信息失败", c)
		return
	}
	response.OkWithMessage("修改发热信息成功", c)
}


//新增
func (f *FeverApi)AddFeverInfo(c *gin.Context) {
	//取参
	var feverInfo system.FeverInfo
	if err := c.ShouldBindJSON(&feverInfo);err != nil{
		response.FailWithMessage(err.Error(),c)
		return
	}
	//校验
	if err := utils.Verify(feverInfo, utils.HospitalReqVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//获取当前用户id
	feverInfo.ID = utils.GetUserID(c)
	//编辑
	if err := feverService.AddFeverInfo(feverInfo); err != nil {
		response.FailWithDetailed(err.Error(),"新增发热信息失败", c)
		return
	}
	response.OkWithMessage("新增发热信息成功", c)
}