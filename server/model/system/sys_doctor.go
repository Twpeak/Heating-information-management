package system

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model

	//用户名
	Username string `json:"username"`
	//姓名
	Name string `json:"name"`
	//身份证号
	IdentityCard string `json:"identity_card"`
	//电话号码
	Phone string `json:"phone"`
	//所属医院
	HospitalId uint `json:"hospital_id"`
}
