package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

type HospitalService struct {}

func (HospitalService *HospitalService)InitHospital()  {
	hdb := global.G_DB.Model(&system.Hospital{})
	sortDate := []system.Hospital{
		{
			Name: "新乡市凤泉区人民医院",
			Code: "11111111111",
			Address: "凤泉区区府路西段",
			BoosId: 1,
			DistrictId:1 ,
		},
	}

	for _,date := range sortDate{
		if err := hdb.FirstOrCreate(&system.Hospital{},&date).Error;err != nil{
			global.G_LOG.Error("分类数据初始化失败",zap.Error(err))
			return
		}
	}
	return
}