package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/dto"
	"github.com/gin-gonic/gin"
)

func DistrictsText(c *gin.Context) {
	page := c.Query("page")
	offset := c.Query("offset")
	limit, err := districtService.QueryDistrictLimit(page, offset)
	if err != nil {
		response.FailWithMessage("查询数据出错", c)
		return
	}
	response.OkWithData(limit, c)
}

func DistrictsUpdate(c *gin.Context) {
	var dis system.District
	_ = c.ShouldBindJSON(&dis)
	err := districtService.UpdateDistrict(dis)
	if err != nil {
		response.FailWithMessage("修改数据出错", c)
		return
	}
	response.Ok(c)
}

func DistrictsDel(c *gin.Context) {
	var id dto.IdDto
	_ = c.ShouldBindJSON(&id)
	err := districtService.DeleteDistrict(id.Id)
	if err != nil {
		response.FailWithMessage("删除数据出错", c)
		return
	}
}
