package dto

import (
	"time"
)

type UserInformationDto struct {
	ID        uint      `json:"id"`         // 主键ID
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	// 用户登录名
	Username string `json:"userName"`
	// 真实姓名
	Name string `json:"name"`
	//身份证号
	IdentityCard string `json:"identity_card" gorm:"comment:身份证号;"` //想写一个，注册时默认登录密码为身份证号码后六位的接口
	//电话号码
	Phone string `json:"phone" gorm:"comment:电话号码;"`
	//医院名称
	HospitalName string `json:"hospital_name"`
}

type UserCreateDto struct {
	//用户名
	Username string `json:"username"`
	//角色
	RoleId uint `json:"role_id"`
	//密码
	Password string `json:"password"`
	//姓名
	Name string `json:"name"`
	//身份证号
	IdentityCard string `json:"identity_card"`
	//电话号码
	Phone string `json:"phone" `
	//所属医院
	HospitalId uint `json:"hospital_id"`
}

type UserUpdateDto struct {
	Id uint `json:"id"`
	//角色
	RoleId uint `json:"role_id"`
	//密码
	Password string `json:"password"`
	//姓名
	Name string `json:"name"`
	//身份证号
	IdentityCard string `json:"identity_card"`
	//电话号码
	Phone string `json:"phone" `
}

type IdDto struct {
	Id uint `json:"id"`
}

type UserTextDto struct {
	Id uint `json:"id"`
	//真实姓名
	Name string `json:"name"`
	//身份证号
	IdentityCard string `json:"identity_card"`
	//电话号码
	Phone string `json:"phone"`
}

type MyPwdDto struct {
	Id     uint   `json:"id"`
	Pwd1   string `json:"pwd_1"`
	Pwd2   string `json:"pwd_2"`
	OldPwd string `json:"old_pwd"`
}
