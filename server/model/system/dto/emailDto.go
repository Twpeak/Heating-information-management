package dto

type AdminPush struct {
	//医院id
	Id uint `json:"Id"`
	//医院名称
	HospitalName string `json:"hospital_name"`
	//医院管理员id
	BoosId uint `json:"boos_id"`
	//医院管理员姓名
	Name string `json:"name"`
	//医院管理员邮箱
	Email string `json:"email"`
	//医院登记总数
	Total      uint     `json:"total"`
	DockerPush []Docker `json:"docker_push"`
}

type Docker struct {
	//医生id
	Id uint `json:"id"`
	//医生姓名
	Name string `json:"name"`
	//登记数量
	Number uint `json:"number"`
}
