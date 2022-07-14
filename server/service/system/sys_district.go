package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

type DistrictService struct{}

func (DistrictService *DistrictService) InitDistrict() {
	ddb := global.G_DB.Model(&system.District{})
	DistrictDate := []system.District{
		{
<<<<<<< Updated upstream
			Id: 1,
			DistrictName:  "卫滨区",
		},
		{
			Id: 2,
			DistrictName: "红旗区",
		},
		{
			Id: 3,
			DistrictName: "牧野区",
		},
		{
			Id: 4,
			DistrictName: "凤泉区",
		},
		{
			Id: 5,
			DistrictName: "新乡县",
		},
		{
			Id: 6,
			DistrictName: "获嘉县",
		},
		{
			Id: 7,
			DistrictName: "卫辉市",
		},
		{
			Id: 8,
			DistrictName: "原阳县",
		},
		{
			Id: 9,
			DistrictName: "延津县",
		},
		{
			Id: 10	,
=======
			Id:           1,
			DistrictName: "卫滨区",
		},
		{
			Id:           2,
			DistrictName: "红旗区",
		},
		{
			Id:           3,
			DistrictName: "牧野区",
		},
		{
			Id:           4,
			DistrictName: "凤泉区",
		},
		{
			Id:           5,
			DistrictName: "新乡县",
		},
		{
			Id:           6,
			DistrictName: "获嘉县",
		},
		{
			Id:           7,
			DistrictName: "卫辉市",
		},
		{
			Id:           8,
			DistrictName: "原阳县",
		},
		{
			Id:           9,
			DistrictName: "延津县",
		},
		{
			Id:           10,
>>>>>>> Stashed changes
			DistrictName: "封丘县",
		},
	}

	for _, date := range DistrictDate {
		if err := ddb.FirstOrCreate(&system.District{}, &date).Error; err != nil {
			global.G_LOG.Error("分类数据初始化失败", zap.Error(err))
			return
		}
	}
	return
}
