package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/dto"
	"go.uber.org/zap"
	"strconv"
)

type FeverService struct {
}

//删除
func (f *FeverService)DelFeverInfo(id string) (err error) {
	if err = global.G_DB.Model(&system.FeverInfo{}).Delete(&system.FeverInfo{},"id = ?",id).Error;
	err != nil{
		global.G_LOG.Error("删除发热记录数据失败", zap.Error(err))
		return err
	}
	return err
}

//修改
func (f *FeverService)UpdateFeverInfo(info system.FeverInfo) (err error) {
	if err = global.G_DB.Model(&system.FeverInfo{}).Save(&info).Error;err != nil{
		global.G_LOG.Error("修改发热记录数据失败", zap.Error(err))
		return err
	}
	return err
}

//新增
func (f *FeverService)AddFeverInfo(info system.FeverInfo) (err error) {
	if err = global.G_DB.Model(&system.FeverInfo{}).Create(&info).Error;err != nil{
		global.G_LOG.Error("新增发热记录数据失败", zap.Error(err))
		return err
	}
	return err
}
func (f *FeverService) QueryFeverLimit(page, offset, name, start, send string) ([]dto.FeverDto, int) {
	i, _ := strconv.Atoi(page)
	o, _ := strconv.Atoi(offset)
	var d []dto.FeverDto
	if name == "" {
		fe := global.G_DB.Model(&system.FeverInfo{}).Where("fever_infos.created_at BETWEEN ? AND ?", start, send).Offset((o - 1) * i).Limit(i).
			Select("fever_infos.id,fever_infos.created_at,fever_infos.updated_at,fever_infos.fever_info_name,fever_infos.gender,fever_infos.age,fever_infos.identity_card,fever_infos.phone,fever_infos.address,fever_infos.clinical_manifestation," +
				"fever_infos.nucleic_acid,fever_infos.problematic,fever_infos.inspection_results,fever_infos.preliminary_diagnosis,fever_infos.disposal,sys_users.name").Joins("left join sys_users on fever_infos.doctor_id=sys_users.id").
			Find(&d)
		return d, int(fe.RowsAffected)
	}
	fe := global.G_DB.Model(&system.FeverInfo{}).Where("(fever_infos.created_at BETWEEN ? AND ?) AND (fever_infos.fever_info_name LIKE ?)", start, send, "%"+name+"%").Offset((o - 1) * i).Limit(i).
		Select("fever_infos.id,fever_infos.created_at,fever_infos.updated_at,fever_infos.fever_info_name,fever_infos.gender,fever_infos.age,fever_infos.identity_card,fever_infos.phone,fever_infos.address,fever_infos.clinical_manifestation," +
			"fever_infos.nucleic_acid,fever_infos.problematic,fever_infos.inspection_results,fever_infos.preliminary_diagnosis,fever_infos.disposal,sys_users.name").Joins("left join sys_users on fever_infos.doctor_id=sys_users.id").
		Find(&d)
	return d, int(fe.RowsAffected)

}

func (f *FeverService) QueryFever() ([]dto.FeverDto, error) {
	var d []dto.FeverDto
	err := global.G_DB.Model(&system.FeverInfo{}).
		Select("fever_infos.id,fever_infos.created_at,fever_infos.updated_at,fever_infos.fever_info_name,fever_infos.gender,fever_infos.age,fever_infos.identity_card,fever_infos.phone,fever_infos.address,fever_infos.clinical_manifestation," +
			"fever_infos.nucleic_acid,fever_infos.problematic,fever_infos.inspection_results,fever_infos.preliminary_diagnosis,fever_infos.disposal,sys_users.name").Joins("left join sys_users on fever_infos.doctor_id=sys_users.id").
		Find(&d).Error
	return d, err
}
