package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

//医院展示信息中，有三表关联，所以参数最多需要三个id
type HospitalReq struct {
	request.PageInfo
	HospitalId	uint	`json:"hospital_id"`
	UserId   	uint	`json:"user_id"`
	DistrictId 	uint	`json:"district_id"`
}

//关键字
type KeyReq struct {
	request.PageInfo
	Key string	`json:"key"`
}

//同时添加医院和用户的接口
type HospitalAndBoss struct {
	ID uint			`json:"id"`	//医院id
	//企业名称
	HospitalName string `json:"hospital_name" gorm:"comment:企业名称;"`
	//社会信用代码
	Code string 		`json:"code" gorm:"comment:社会信用代码;"`
	//注册地址
	Address string 		`json:"address" gorm:"comment:注册地址;"`
	//区县id
	DistrictId uint 	`json:"district_id" gorm:"comment:区县id;"`
	//用户名
	Username string    `json:"username" gorm:"comment:用户登录名;unique;"`   // 用户登录名
	//身份证号
	IdentityCard string `json:"identity_card" gorm:"comment:身份证号;"` //想写一个，注册时默认登录密码为身份证号码后六位的接口
	//电话号码
	Phone string `json:"phone" gorm:"comment:电话号码;"`
	//电子邮箱
	Email string `json:"email" gorm:"comment:电子邮箱;"`
}