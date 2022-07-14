package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"go.uber.org/zap"
)

type HospitalService struct{}

<<<<<<< Updated upstream
//查询所有医院信息
func (h *HospitalService)GetAllHospital()(list []system.Hospital,err error)  {
	if err = global.G_DB.Model(&system.Hospital{}).Find(&list).Error; err != nil{
		global.G_LOG.Error("查询医院信息失败",zap.Error(err))
		return list,err
	}
	return list,err
}

//查询医院负责人信息
func (h *HospitalService)GetBossByBossId(id uint)(user system.SysUser,err error)  {
	if err = global.G_DB.Model(&system.SysUser{}).Where("id = ?",id).Find(&user).Error;err != nil{
		global.G_LOG.Error("查询医院负责人信息失败",zap.Error(err))
		return user,err
	}
	return user,err
}

//查询该医院县区信息
func (h *HospitalService)GetDistrictByDistrictId(id uint)(dis system.District,err error)  {
	if err = global.G_DB.Model(&system.District{}).Where("id = ?",id).Find(&dis).Error;err != nil{
		global.G_LOG.Error("查询医院区县信息失败",zap.Error(err))
		return dis,err
	}
	return dis,err
}


//链表查询医院信息和负责人信息
func (h *HospitalService)GetHospitalsVo()(voDate []response.HospitalVo,err error)  {
	if err = global.G_DB.Model(&system.Hospital{}).Debug().Select("hospitals.*,districts.district_name,sys_users.*").
		Joins("left join districts on districts.id = hospitals.district_id").
		Joins("left join sys_users on sys_users.id = hospitals.boos_id").
		Scan(&voDate).Error
	err != nil{
		global.G_LOG.Error("链表查询医院信息VO数据失败",zap.Error(err))
		return voDate,err
	}
	return voDate,err
}





//初始化医院信息
func (HospitalService *HospitalService)InitHospital()  {
=======
func (HospitalService *HospitalService) InitHospital() {
>>>>>>> Stashed changes
	hdb := global.G_DB.Model(&system.Hospital{})
	sortDate := []system.Hospital{
		{
			HospitalName: "新乡市凤泉区人民医院",
<<<<<<< Updated upstream
			Code: "11111111111",
			Address: "凤泉区区府路西段",
=======
			Code:         "11111111111",
			Address:      "凤泉区区府路西段",
			BoosId:       1,
			DistrictId:   1,
>>>>>>> Stashed changes
		},
	}

	for _, date := range sortDate {
		if err := hdb.FirstOrCreate(&system.Hospital{}, &date).Error; err != nil {
			global.G_LOG.Error("分类数据初始化失败", zap.Error(err))
			return
		}
	}
	return
}

func (h *HospitalService) QueryBoosId(id uint) (system.Hospital, error) {
	var hospital system.Hospital
	err := global.G_DB.Model(&system.Hospital{}).Where("boos_id = ?", id).Scan(&hospital).Error
	return hospital, err
}
