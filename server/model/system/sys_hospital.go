package system

import "gorm.io/gorm"

type Hospital struct {
	gorm.Model
	//企业名称
<<<<<<< Updated upstream
	HospitalName string `json:"name" gorm:"comment:企业名称;"`
=======
	HospitalName string `json:"hospital_name" gorm:"comment:企业名称;"`
>>>>>>> Stashed changes
	//社会信用代码
	Code string `json:"code" gorm:"comment:社会信用代码;"`
	//注册地址
	Address string `json:"address" gorm:"comment:注册地址;"`
	//负责人/医生 id		仅逻辑关联，无外键关联。
	BoosId uint `json:"boos_id" gorm:"comment:负责人ID;"`
	//区县id
	DistrictId uint `json:"district_id" gorm:"comment:区县id;"`
}
