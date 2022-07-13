package system

import "gorm.io/gorm"

type Hospital struct {
	gorm.Model
	//企业名称
	Name string `json:"name"`
	//社会信用代码
	Code int `json:"code"`
	//注册地址
	Address string `json:"address"`
	//负责人/医生 id
	DoctorId uint `json:"doctor_id"`
	//区县id
	DistrictId uint `json:"district_id"`
}
