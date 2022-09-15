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

type FeverDto struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//名字
	FeverInfoName string `json:"fever_info_name" gorm:"comment:患者姓名;"`
	//性别
	Gender bool `json:"gender" gorm:"comment:性别;"`
	//年龄
	Age uint `json:"age" gorm:"comment:年龄;"`
	//身份证号
	IdentityCard string `json:"identity_card" gorm:"comment:身份证号;"`
	//手机号
	Phone string `json:"phone" gorm:"comment:手机号;"`
	//现住址
	Address string `json:"address" gorm:"comment:现住址;"`
	//临床表现
	ClinicalManifestation string `json:"clinical_manifestation" gorm:"comment:临床表现;"`
	//48小时核算证明
	NucleicAcid bool `json:"nucleic_acid" gorm:"comment:48小时核算证明;"`
	//是否红黄码
	Problematic bool `json:"problematic" gorm:"comment:是否红黄码;"`
	//检查校验项目及结果
	InspectionResults string `json:"inspection_results" gorm:"comment:检查校验项目及结果;"`
	//初步诊断
	PreliminaryDiagnosis string `json:"preliminary_diagnosis" gorm:"comment:初步诊断;"`
	//处置方式
	Disposal string `json:"disposal" gorm:"comment:处置方式;"`
	//医生name
	Name string `json:"name" `
}
