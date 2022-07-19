package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
	"strconv"
)

type DistrictService struct{}

func (DistrictService *DistrictService) InitDistrict() {
	ddb := global.G_DB.Model(&system.District{})
	DistrictDate := []system.District{
		{
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

func (d *DistrictService) QueryDistrictLimit(page, offset string) ([]system.District, error) {
	i, _ := strconv.Atoi(page)
	o, _ := strconv.Atoi(offset)
	var dis []system.District
	err := global.G_DB.Model(system.District{}).Offset((o - 1) * i).Limit(i).Find(&dis).Error

	return dis, err
}

func (d *DistrictService) UpdateDistrict(dis system.District) error {
	return global.G_DB.Model(system.District{}).Where("id = ?", dis.Id).Update("district_name", dis.DistrictName).Error
}

func (d *DistrictService) DeleteDistrict(id uint) error {
	return global.G_DB.Model(system.District{}).Where("id = ?", id).Delete(system.District{}, id).Error
}
