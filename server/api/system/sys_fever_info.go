package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func FeverTextLimit(c *gin.Context) {
	page := c.Query("page")
	i, _ := strconv.Atoi(page)
	offset := c.Query("offset")
	name := c.Query("name")
	startData := c.Query("startData")
	sendData := c.Query("sendData")
	limit, num := feverService.QueryFeverLimit(page, offset, name, startData, sendData)

	response.OkWithData(gin.H{
		"data":    limit,
		"number":  num,
		"pageAll": (num / i) + 1,
	}, c)
}
func Export(c *gin.Context) {
	fever, err := feverService.QueryFever()
	if err != nil {
		global.G_LOG.Error("接口:Export,导出excel失败,error:" + err.Error())
		response.FailWithMessage("导出excel失败", c)
		return
	}
	excel := excelize.NewFile()
	titleSlice := []interface{}{"创建时间", "更新时间", "名字", "性别", "年龄", "身份证号", "手机号", "现住址", "临床表现", "是否有48小时核算证明", "是否红黄码", "检查校验项目及结果", "初步诊断", "处置方式", "医生名字"}
	_ = excel.SetSheetRow("Sheet1", "A1", &titleSlice)
	var data []interface{}
	for i := 0; i < len(fever); i++ {
		var da []interface{}
		da = append(da, fever[i].CreatedAt)
		da = append(da, fever[i].UpdatedAt)
		da = append(da, fever[i].FeverInfoName)
		da = append(da, fever[i].Gender)
		da = append(da, fever[i].Age)
		da = append(da, fever[i].IdentityCard)
		da = append(da, fever[i].Phone)
		da = append(da, fever[i].Address)
		da = append(da, fever[i].ClinicalManifestation)
		da = append(da, fever[i].NucleicAcid)
		da = append(da, fever[i].Problematic)
		da = append(da, fever[i].InspectionResults)
		da = append(da, fever[i].PreliminaryDiagnosis)
		da = append(da, fever[i].Disposal)
		da = append(da, fever[i].Name)
		data = append(data, da)
	}
	for i, v := range data {
		axis := fmt.Sprintf("A%d", i+2)
		tmp, _ := v.([]interface{})
		_ = excel.SetSheetRow("Sheet1", axis, &tmp)
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"data.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	//回写到web 流媒体 形成下载
	_ = excel.Write(c.Writer)
	response.Ok(c)

}
