package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

type DistrictService struct {}

func (DistrictService *DistrictService)InitDistrict()  {
	ddb := global.G_DB.Model(&system.District{})
	DistrictDate := []system.District{
		{
			Id: 1,
			Name: "卫滨区",
		},
		{
			Id: 2,
			Name: "红旗区",
		},
		{
			Id: 3,
			Name: "牧野区",
		},
		{
			Id: 4,
			Name: "凤泉区",
		},
		{
			Id: 5,
			Name: "新乡县",
		},
		{
			Id: 6,
			Name: "获嘉县",
		},
		{
			Id: 7,
			Name: "卫辉市",
		},
		{
			Id: 8,
			Name: "原阳县",
		},
		{
			Id: 9,
			Name: "延津县",
		},
		{
			Id: 10	,
			Name: "封丘县",
		},
	}

	for _,date := range DistrictDate{
		if err := ddb.FirstOrCreate(&system.District{},&date).Error;err != nil{
			global.G_LOG.Error("分类数据初始化失败",zap.Error(err))
			return
		}
	}
	return
}