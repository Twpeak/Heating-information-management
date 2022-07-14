package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HospitalService struct{}

//查询所有医院信息
func (h *HospitalService) GetAllHospital() (list []system.Hospital, err error) {
	if err = global.G_DB.Model(&system.Hospital{}).Find(&list).Error; err != nil {
		global.G_LOG.Error("查询医院信息失败", zap.Error(err))
		return list, err
	}
	return list, err
}

//查询医院负责人信息
func (h *HospitalService) GetBossByBossId(id uint) (user system.SysUser, err error) {
	if err = global.G_DB.Model(&system.SysUser{}).Where("id = ?", id).Find(&user).Error; err != nil {
		global.G_LOG.Error("查询医院负责人信息失败", zap.Error(err))
		return user, err
	}
	return user, err
}

//查询该医院县区信息
func (h *HospitalService) GetDistrictByDistrictId(id uint) (dis system.District, err error) {
	if err = global.G_DB.Model(&system.District{}).Where("id = ?", id).Find(&dis).Error; err != nil {
		global.G_LOG.Error("查询医院区县信息失败", zap.Error(err))
		return dis, err
	}
	return dis, err
}

//链表查询医院信息和负责人信息
func (h *HospitalService) GetHospitalsVo() (voDate []response.HospitalVo, err error) {
	if err = global.G_DB.Model(&system.Hospital{}).Debug().Select("hospitals.*,districts.district_name,sys_users.*").
		Joins("left join districts on districts.id = hospitals.district_id").
		Joins("left join sys_users on sys_users.id = hospitals.boos_id").
		Scan(&voDate).Error; err != nil {
		global.G_LOG.Error("链表查询医院信息VO数据失败", zap.Error(err))
		return voDate, err
	}
	return voDate, err
}

//查询当前医院所有医生
func (h *HospitalService) GetUserByHospitalId(Hospital request.HospitalReq) (users []system.SysUser, err error) {
	HospitalId := Hospital.HospitalId
	if err = global.G_DB.Model(system.SysUser{}).Where("hospital_id = ?", HospitalId).Find(&users).Error; err != nil {
		global.G_LOG.Error("查询当前医院所有医生失败", zap.Error(err))
		return users, err
	}
	return users, err
}

//删除医院信息
func (h *HospitalService) DelHospital(req request.HospitalReq) (err error) {
	HospitalId := req.HospitalId
	if err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		//物理删除
		if err := tx.Model(&system.Hospital{}).Unscoped().Delete(&system.Hospital{}, "id = ?", HospitalId).Error; err != nil {
			global.G_LOG.Error("删除医院数据失败", zap.Error(err))
			return err
		}
		//物理删除
		if err := tx.Model(&system.SysUser{}).Unscoped().Delete(&system.SysUser{}, "hospital_id = ?", HospitalId).Error; err != nil {
			global.G_LOG.Error("删除医院中医生数据失败", zap.Error(err))
			return err
		}
		return nil
	}); err != nil {
		global.G_LOG.Error("删除医院事务出现错误，数据回滚", zap.Error(err))
		return err
	}
	return err
}

//修改医院信息
func (h *HospitalService) UpdateHospital(Hospital system.Hospital) (err error) {
	if err = global.G_DB.Model(&system.Hospital{}).Where("id = ?", Hospital.ID).
		Updates(Hospital).Error; err != nil {
		global.G_LOG.Error("修改医院信息失败", zap.Error(err))
		return err
	}
	return err
}

//修改负责人信息(id)
func (h *HospitalService) UpdateHospitalByUser(req request.HospitalReq) (err error) {
	if err = global.G_DB.Model(&system.Hospital{}).Where("id = ?", req.HospitalId).
		Update("boos_id", req.UserId).Error; err != nil {
		global.G_LOG.Error("修改医院负责人信息失败", zap.Error(err))
		return err
	}
	return err
}

//新增医院信息（基本信息）
func (h *HospitalService) AddHospital(Hospital system.Hospital) (err error) {
	if err = global.G_DB.Model(&system.Hospital{}).Create(&Hospital).Error; err != nil {
		global.G_LOG.Error("添加医院信息失败", zap.Error(err))
		return err
	}
	return err
}

//查询当前区域所有医院列表【分页】
func (h *HospitalService) GetHospitalByDistrictLimit(req request.HospitalReq) (hos []system.Hospital, err error) {
	if err = global.G_DB.Model(&system.Hospital{}).Where("district_id = ?", req.DistrictId).Find(&hos).Error; err != nil {
		global.G_LOG.Error("修改医院负责人信息失败", zap.Error(err))
		return hos, err
	}
	return hos, err
}

//初始化医院信息
func (HospitalService *HospitalService) InitHospital() {
	hdb := global.G_DB.Model(&system.Hospital{})
	sortDate := []system.Hospital{
		{
			HospitalName: "新乡市凤泉区人民医院",

			Code:       "11111111111",
			Address:    "凤泉区区府路西段",
			BoosId:     1,
			DistrictId: 1,
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
