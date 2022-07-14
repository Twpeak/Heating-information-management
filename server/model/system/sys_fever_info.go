package system

import "gorm.io/gorm"

type FeverInfo struct {
	gorm.Model
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
	//医生id
	DoctorId uint `json:"doctor_id" gorm:"comment:医生id;"`
}
