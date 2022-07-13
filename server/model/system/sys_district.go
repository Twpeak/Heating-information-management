package system

type District struct {
	//区县id
	Id uint 				`json:"id" gorm:"comment:区县ID;"`
	//区县名称
	Name string 			`json:"name" gorm:"comment:区县名称;"`
	//关联医院信息
	Hospitals []Hospital	`json:"hospitals" gorm:"FOREIGNKEY:DistrictId"`
}
