package response

//返回医院信息 vo返回体   医院信息  + 负责人 + 县区
type HospitalVo struct {
	ID        					uint `gorm:"primarykey"`
	//企业名称
	HospitalName string 		`json:"hospital_name" gorm:"comment:企业名称;"`
	//社会信用代码
	Code string 				`json:"code" gorm:"comment:社会信用代码;"`
	//注册地址
	Address string 				`json:"address" gorm:"comment:注册地址;"`
	//负责人姓名
	Username string				`json:"user_name"`
	//身份证号
	IdentityCard string 		`json:"identity_card" gorm:"comment:身份证号;"`//想写一个，注册时默认登录密码为身份证号码后六位的接口
	//电话号码
	Phone 		 string 		`json:"phone" gorm:"comment:电话号码;"`
	//关联区县字段
	DistrictName string			`json:"district_name"` //县区名
}
