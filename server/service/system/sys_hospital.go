package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	reqCom "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
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
func (h *HospitalService) GetBossByBossId(HospitalId uint) (user system.SysUser, err error) {
	var userid uint
	global.G_DB.Model(system.Hospital{}).Select("boos_id").Where("id = ?",HospitalId).Scan(&userid)
	if err = global.G_DB.Model(&system.SysUser{}).Where("id = ?",userid ).Scan(&user).Error; err != nil {
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

//链表查询医院信息和负责人信息[分页]
func (h *HospitalService) GetHospitalsVo(pageInfo reqCom.PageInfo) (voDate []response.HospitalVo,total int64,  err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	if err = global.G_DB.Model(&system.Hospital{}).Limit(limit).Offset(offset).Debug().Select("hospitals.*,districts.district_name,sys_users.*").
		Joins("left join districts on districts.id = hospitals.district_id").
		Joins("left join sys_users on sys_users.id = hospitals.boos_id").
		Scan(&voDate).Count(&total).Error; err != nil {
		global.G_LOG.Error("链表查询医院信息VO数据失败", zap.Error(err))
		return
	}
	return
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

//修改医院信息同时修改负责人信息
func (h *HospitalService) UpdateHospitalAndBoss(HospitalAndBoss request.HospitalAndBoss) (err error) {
	//修改医院信息
	hospital :=  &system.Hospital{
		Model: 			gorm.Model{ID: HospitalAndBoss.ID},
		HospitalName: 	HospitalAndBoss.HospitalName,
		Code: 			HospitalAndBoss.Code,
		Address: 		HospitalAndBoss.Address,
		DistrictId: 	HospitalAndBoss.DistrictId,
	}
	boss := &system.SysUser{
		Username: 	HospitalAndBoss.Username,
		Name: 		HospitalAndBoss.Username,
		Email: 		HospitalAndBoss.Email,
		Phone: 		HospitalAndBoss.Phone,
		Password: 	HospitalAndBoss.IdentityCard[len(HospitalAndBoss.IdentityCard) - 6:],
		RoleId: 	3,
	}
	if err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&system.Hospital{}).Updates(hospital).Error; err != nil {
			global.G_LOG.Error("修改医院信息失败", zap.Error(err))
			return err
		}
		boss.HospitalId = hospital.ID
		boss.ID = hospital.BoosId
		if err = tx.Model(&system.SysUser{}).Updates(boss).Error; err != nil {
			global.G_LOG.Error("修改管理者信息失败", zap.Error(err))
			return err
		}
		return nil
	}); err != nil {
		global.G_LOG.Error("修改医院事务出现错误，数据回滚", zap.Error(err))
		return err
	}
	return err
}

//新增医院信息（同时添加管理员信息）
func (h *HospitalService) AddHospitalAndBoss(HospitalAndBoss request.HospitalAndBoss) (err error) {
	//添加医院信息
	hospital :=  &system.Hospital{
		HospitalName: HospitalAndBoss.HospitalName,
		Code: HospitalAndBoss.Code,
		Address: HospitalAndBoss.Address,
		DistrictId: HospitalAndBoss.DistrictId,
	}
	boss := &system.SysUser{
		Username: HospitalAndBoss.Username,
		Name: HospitalAndBoss.Username,
		Email: HospitalAndBoss.Email,
		Phone: HospitalAndBoss.Phone,
		Password: HospitalAndBoss.IdentityCard[len(HospitalAndBoss.IdentityCard) - 6:],
		RoleId: 3,
	}
	if err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&system.Hospital{}).Create(&hospital).Error; err != nil {
			global.G_LOG.Error("添加医院信息失败", zap.Error(err))
			return err
		}
		boss.HospitalId = hospital.ID
		if err = tx.Model(&system.SysUser{}).Create(&boss).Error; err != nil {
			global.G_LOG.Error("添加管理者信息失败", zap.Error(err))
			return err
		}
		return nil
	}); err != nil {
		global.G_LOG.Error("添加医院事务出现错误，数据回滚", zap.Error(err))
		return err
	}
	return err
}

//查询当前区域所有医院列表【分页】
func (h *HospitalService) GetHospitalByDistrictLimit(req request.HospitalReq) (hos []system.Hospital,total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	if err = global.G_DB.Limit(limit).Offset(offset).Model(&system.Hospital{}).
		Where("district_id = ?", req.DistrictId).Find(&hos).Count(&total).Error; err != nil {
		global.G_LOG.Error("查询当前区域所有医院列表信息失败", zap.Error(err))
		return
	}
	return
}

//通过医院名查询医院数据【分页】
func (h *HospitalService) GetHospitalByHospitalName(req request.KeyReq) (hos []system.Hospital,total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	key := req.Key
	if err = global.G_DB.Limit(limit).Offset(offset).Model(&system.Hospital{}).
		Where("hospital_name LIKE ?","%"+key+"%").Find(&hos).Count(&total).Error;err != nil{
		global.G_LOG.Error("通过医院名查找医院信息失败", zap.Error(err))
		return
	}
	return
}

//通过关键字查询医院视图数据【分页】
func (h *HospitalService) GetHospitalByKey(req request.KeyReq) (hos []response.HospitalVo,total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	key := req.Key
	vo,total, _ := h.GetHospitalsVo(req.PageInfo)
	for _,Item := range vo{
		switch  {
		case strings.Contains(strconv.Itoa(int(Item.ID)),key):
			hos = append(hos,Item)
		case strings.Contains(Item.HospitalName,key):
			hos = append(hos,Item)
		case strings.Contains(Item.Code,key):
			hos = append(hos,Item)
		case strings.Contains(Item.Address,key):
			hos = append(hos,Item)
		case strings.Contains(Item.Username,key):
			hos = append(hos,Item)
		case strings.Contains(Item.IdentityCard,key):
			hos = append(hos,Item)
		case strings.Contains(Item.Phone,key):
			hos = append(hos,Item)
		case strings.Contains(Item.DistrictName,key):
			hos = append(hos,Item)
		}
	}
	return hos[offset:offset+limit], total,err
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


//通过负责人查找医院？
func (h *HospitalService) QueryBoosId(id uint) (system.Hospital, error) {
	var hospital system.Hospital
	err := global.G_DB.Model(&system.Hospital{}).Where("boos_id = ?", id).Scan(&hospital).Error
	return hospital, err
}
