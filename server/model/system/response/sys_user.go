package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

/**
用户返回体
 */

//登录时获取用户信息的结构体
type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

//用户登录成功的返回信息
type LoginResponse struct {
	User system.SysUser `json:"user"`
	Token string		`json:"token"`
	ExpiresAt int64		`json:"expiresAt"`
}

//查询当前医院所有用户，返回用户信息vo
type GetDoctorByHos struct {
	ID 	uint `json:"id"`
	//身份证号
	IdentityCard string `json:"identity_card" gorm:"comment:身份证号;"` //想写一个，注册时默认登录密码为身份证号码后六位的接口
	//电话号码
	Phone string `json:"phone" gorm:"comment:电话号码;"`
	//电子邮箱
	Email string `json:"email" gorm:"comment:电子邮箱;"`
	// 真实姓名
	Name     string    `json:"name" gorm:"default:'系统用户';comment:真实姓名"`
	//所属医院
	HospitalName string `json:"hospital_name" gorm:"hospital_name"`
}
