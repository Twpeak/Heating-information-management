package system

import "gorm.io/gorm"

type FeverInfo struct {
	gorm.Model
	//名字
	Name string `json:"name"`
	//性别
	Gender bool `json:"gender"`
	//年龄
	Age uint `json:"age"`
	//身份证号
	IdentityCard string `json:"identity_card"`
	//手机号
	Phone string `json:"phone"`
	//现住址
	Address string `json:"address"`
	//临床表现
	ClinicalManifestation string `json:"clinical_manifestation"`
	//48小时核算证明
	NucleicAcid bool `json:"nucleic_acid"`
	//是否红黄码
	Problematic bool `json:"problematic"`
	//检查校验项目及结果
	InspectionResults string `json:"inspection_results"`
	//初步诊断
	PreliminaryDiagnosis string `json:"preliminary_diagnosis"`
	//处置方式
	Disposal string `json:"disposal"`
	//医生id
	DoctorId uint `json:"doctor_id"`
}
